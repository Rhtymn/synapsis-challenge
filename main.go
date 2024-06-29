package main

import (
	"fmt"
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/config"
	"github.com/Rhtymn/synapsis-challenge/database"
	"github.com/Rhtymn/synapsis-challenge/handler"
	"github.com/Rhtymn/synapsis-challenge/middleware"
	repository "github.com/Rhtymn/synapsis-challenge/repository/postgres"
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

	userAccessProvider := util.NewJWTProvider(
		conf.JWTIssuer,
		conf.UserAccessSecret,
		conf.AccessTokenLifespan,
	)
	sellerAccessProvider := util.NewJWTProvider(
		conf.JWTIssuer,
		conf.SellerAccessSecret,
		conf.AccessTokenLifespan,
	)
	adminAccessProvider := util.NewJWTProvider(
		conf.JWTIssuer,
		conf.AdminAccessSecret,
		conf.AccessTokenLifespan,
	)

	accountRepository := repository.NewAccountRepository(db)
	userRepository := repository.NewUserRepository(db)
	emailVerifyTokenRepository := repository.NewEmailVerifyTokenRepository(db)
	passwordHasher := util.NewPasswordHasherBcrypt(10)
	transactor := util.NewTransactor(db)

	accountSrv := service.NewAccountService(service.AccountServiceOpts{
		Account:              accountRepository,
		User:                 userRepository,
		EmailVerifyToken:     emailVerifyTokenRepository,
		PasswordHasher:       passwordHasher,
		Transactor:           transactor,
		UserAccessProvider:   userAccessProvider,
		SellerAccessProvider: sellerAccessProvider,
		AdminAccessProvider:  adminAccessProvider,
	})

	accountHandler := handler.NewAccountHandler(handler.AccountHandlerOpts{
		Account: accountSrv,
		Domain:  "account",
	})
	corsHandler := middleware.CorsHandler(conf.CorsDomain)
	errorHandler := middleware.ErrorHandler()

	router := server.SetupServer(server.ServerOpts{
		CorsHandler:    corsHandler,
		ErrorHandler:   errorHandler,
		AccountHandler: accountHandler,
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
