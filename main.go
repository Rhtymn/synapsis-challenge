package main

import (
	"fmt"
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/config"
	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/database"
	"github.com/Rhtymn/synapsis-challenge/handler"
	"github.com/Rhtymn/synapsis-challenge/middleware"
	repository "github.com/Rhtymn/synapsis-challenge/repository/postgres"
	"github.com/Rhtymn/synapsis-challenge/repository/redis"
	"github.com/Rhtymn/synapsis-challenge/server"
	"github.com/Rhtymn/synapsis-challenge/service"
	"github.com/Rhtymn/synapsis-challenge/util"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Printf("Init config error: %s\n", err)
		return
	}

	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Load .env error: %s\n", err)
		return
	}

	db, err := database.ConnectPostgresDB(conf.DatabaseURL)
	if err != nil {
		fmt.Printf("DB Connection Error: %s\n", err)
	}
	defer db.Close()

	rdb := database.NewRedisClient(conf.RedisAddr, conf.RedisPassword, 0)

	userAccessProvider := util.NewJWTProvider(
		constants.USER_PERMISSION,
		conf.JWTIssuer,
		conf.UserAccessSecret,
		conf.AccessTokenLifespan,
	)
	sellerAccessProvider := util.NewJWTProvider(
		constants.SELLER_PERMISSION,
		conf.JWTIssuer,
		conf.SellerAccessSecret,
		conf.AccessTokenLifespan,
	)
	adminAccessProvider := util.NewJWTProvider(
		constants.ADMIN_PERMISSION,
		conf.JWTIssuer,
		conf.AdminAccessSecret,
		conf.AccessTokenLifespan,
	)
	anyAccessProvider := util.NewJWTProviderAny([]util.JWTProvider{userAccessProvider, sellerAccessProvider, adminAccessProvider})

	authenticator := middleware.Authenticator(anyAccessProvider)

	accountRepository := repository.NewAccountRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	transactionItemRepository := repository.NewTransactionItemRepository(db)
	userRepository := repository.NewUserRepository(db)
	userAddressRepository := repository.NewUserAddressRepository(db)
	productRepository := repository.NewProductRepository(db)
	emailVerifyTokenRepository := repository.NewEmailVerifyTokenRepository(db)
	paymentMethodRepository := repository.NewPaymentMethodRepository(db)
	shipmentMethodRepository := repository.NewShipmentMethodRepository(db)
	shopShipmentMethodRepository := repository.NewShopShipmentMethodRepository(db)
	shopPaymentMethodRepository := repository.NewShopPaymentMethodRepository(db)
	PaymentRepository := repository.NewPaymentRepository(db)

	cartRepositoryRedis := redis.NewCartRepositoryRedis(rdb)

	passwordHasher := util.NewPasswordHasherBcrypt(10)
	transactor := util.NewTransactor(db)
	cloudinaryProvider, err := util.NewCloudinaryProvider(util.CloudinaryProviderOpts{
		CloudinaryName:      conf.CloudinaryName,
		CloudinaryAPIKey:    conf.CloudinaryAPIKey,
		CloudinaryAPISecret: conf.CloudinartAPISecret,
	})
	if err != nil {
		fmt.Printf("Cloudinary error: %s\n", err)
	}
	emailProvider := util.NewEmailProvider(util.EmailProviderOpts{
		Username: conf.AuthEmailUsername,
		Password: conf.AuthEmailPassword,
	})
	appEmail, err := util.NewAppEmail(util.AppEmailOpts{
		FEVerivicationURL: conf.FEVerificationURL,
	})
	if err != nil {
		fmt.Printf("Email template error: %s\n", err)
	}

	accountSrv := service.NewAccountService(service.AccountServiceOpts{
		Account:              accountRepository,
		User:                 userRepository,
		EmailVerifyToken:     emailVerifyTokenRepository,
		PasswordHasher:       passwordHasher,
		Transactor:           transactor,
		UserAccessProvider:   userAccessProvider,
		SellerAccessProvider: sellerAccessProvider,
		AdminAccessProvider:  adminAccessProvider,
		RandomTokenProvider:  util.NewRandomTokenProvider(32),
		EmailProvider:        emailProvider,
		AppEmail:             appEmail,
	})
	userSrv := service.NewUserService(service.UserServiceOpts{
		User:               userRepository,
		UserAddress:        userAddressRepository,
		Account:            accountRepository,
		Transactor:         transactor,
		CloudinaryProvider: cloudinaryProvider,
	})
	productSrv := service.NewProductService(service.ProductServiceOpts{
		Product: productRepository,
	})
	cartSrv := service.NewCartService(service.CartServiceOpts{
		Cart:    cartRepositoryRedis,
		Product: productRepository,
		Account: accountRepository,
	})
	transactionSrv := service.NewTransactionService(service.TransactionServiceOpts{
		Transaction:        transactionRepository,
		TransactionItem:    transactionItemRepository,
		PaymentMethod:      paymentMethodRepository,
		ShipmentMethod:     shipmentMethodRepository,
		UserAddress:        userAddressRepository,
		User:               userRepository,
		Cart:               cartRepositoryRedis,
		Product:            productRepository,
		ShopShipmentMethod: shopShipmentMethodRepository,
		ShopPaymentMethod:  shopPaymentMethodRepository,
		Payment:            PaymentRepository,
		Transactor:         transactor,
		Cloudinary:         cloudinaryProvider,
	})

	accountHandler := handler.NewAccountHandler(handler.AccountHandlerOpts{
		Account: accountSrv,
		Domain:  "account",
	})
	userHandler := handler.NewUserHandler(handler.UserHandlerOpts{
		User:   userSrv,
		Domain: "user",
	})
	productHandler := handler.NewProductHandler(handler.ProductHandlerOpts{
		Product: productSrv,
		Domain:  "product",
	})
	cartHandler := handler.NewCartHandler(handler.CartHandlerOpts{
		Cart:   cartSrv,
		Domain: "cart",
	})
	transactionHandler := handler.NewTransactionHandler(handler.TransactionHandlerOpts{
		Transaction: transactionSrv,
		Domain:      "transaction",
	})
	corsHandler := middleware.CorsHandler(conf.CorsDomain)
	errorHandler := middleware.ErrorHandler()

	router := server.SetupServer(server.ServerOpts{
		CorsHandler:        corsHandler,
		ErrorHandler:       errorHandler,
		AccountHandler:     accountHandler,
		CartHandler:        cartHandler,
		UserHandler:        userHandler,
		TransactionHandler: transactionHandler,
		ProductHandler:     productHandler,
		Authenticator:      authenticator,
	})
	srv := &http.Server{
		Addr:    conf.ServerAddr,
		Handler: router,
	}

	fmt.Printf("Starting Server...\n")
	fmt.Printf("Server running on port %s\n", conf.ServerAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error running server: %v\n", err)
	}
}
