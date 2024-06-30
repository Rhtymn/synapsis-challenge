package service

import (
	"context"
	"mime/multipart"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/util"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type transactionService struct {
	transactionRepository        domain.TransactionRepository
	transactionItemRepository    domain.TransactionItemRepository
	paymentMethodRepository      domain.PaymentMethodRepository
	shipmentMethodRepository     domain.ShipmentMethodRepository
	productRepository            domain.ProductRepository
	userAddressRepository        domain.UserAddressRepository
	userRepository               domain.UserRepository
	cartRepository               domain.CartRepositoryRedis
	shopShipmentMethodRepository domain.ShopShipmentMethodRepository
	shopPaymentMethodRepository  domain.ShopPaymentMethodRepository
	paymentRepository            domain.PaymentRepository
	transactor                   util.Transactor
	cloudinaryProvider           util.CloudinaryProvider
}

type TransactionServiceOpts struct {
	Transaction        domain.TransactionRepository
	TransactionItem    domain.TransactionItemRepository
	PaymentMethod      domain.PaymentMethodRepository
	ShipmentMethod     domain.ShipmentMethodRepository
	UserAddress        domain.UserAddressRepository
	User               domain.UserRepository
	Cart               domain.CartRepositoryRedis
	Product            domain.ProductRepository
	ShopShipmentMethod domain.ShopShipmentMethodRepository
	ShopPaymentMethod  domain.ShopPaymentMethodRepository
	Payment            domain.PaymentRepository
	Transactor         util.Transactor
	Cloudinary         util.CloudinaryProvider
}

func NewTransactionService(opts TransactionServiceOpts) *transactionService {
	return &transactionService{
		transactionRepository:        opts.Transaction,
		transactionItemRepository:    opts.TransactionItem,
		paymentMethodRepository:      opts.PaymentMethod,
		shipmentMethodRepository:     opts.ShipmentMethod,
		productRepository:            opts.Product,
		userAddressRepository:        opts.UserAddress,
		cartRepository:               opts.Cart,
		userRepository:               opts.User,
		shopShipmentMethodRepository: opts.ShopShipmentMethod,
		shopPaymentMethodRepository:  opts.ShopPaymentMethod,
		paymentRepository:            opts.Payment,
		transactor:                   opts.Transactor,
		cloudinaryProvider:           opts.Cloudinary,
	}
}

func (s *transactionService) CreateTransaction(ctx context.Context, t domain.CreateTransactionRequest) (domain.Transaction, error) {
	ts := domain.Transaction{}
	err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		accountID, err := util.GetAccountIDFromContext(ctx)
		if err != nil {
			return apperror.Wrap(err)
		}

		user, err := s.userRepository.GetByAccountID(ctx, accountID)

		product, err := s.productRepository.GetByIDAndLock(ctx, t.ProductID)
		if err != nil {
			return apperror.Wrap(err)
		}

		cart, err := s.cartRepository.GetByID(ctx, accountID, t.ProductID)
		if err != nil {
			return apperror.Wrap(err)
		}

		newStock := product.Stock - int64(cart.Amount)
		if newStock < 0 {
			return apperror.NewBadRequest(err, "insufficient product stock")
		}

		shipmentMethod, err := s.shipmentMethodRepository.GetByID(ctx, t.ShipmentMethodID)
		if err != nil {
			return apperror.Wrap(err)
		}

		ok, err := s.shopShipmentMethodRepository.IsSupportShipmentMethod(ctx, product.Shop.ID, shipmentMethod.ID)
		if err != nil {
			return apperror.Wrap(err)
		}
		if !ok {
			return apperror.NewBadRequest(nil, "unsupported choosed shipment method")
		}

		paymentMethod, err := s.paymentMethodRepository.GetByID(ctx, t.PaymentMethodID)
		if err != nil {
			return apperror.Wrap(err)
		}

		ok, err = s.shopPaymentMethodRepository.IsSupportPaymentMethod(ctx, product.Shop.ID, shipmentMethod.ID)
		if err != nil {
			return apperror.Wrap(err)
		}
		if !ok {
			return apperror.NewBadRequest(nil, "unsupported choosed payment method")
		}

		userAddress, err := s.userAddressRepository.GetByID(ctx, t.AddressID)
		if err != nil {
			return apperror.Wrap(err)
		}

		if userAddress.User.ID != user.ID {
			return apperror.NewNotFound(nil, "address not found")
		}

		transaction, err := s.transactionRepository.Add(ctx, domain.Transaction{
			Invoice:     util.CreateInvoice(),
			Status:      domain.WaitingForPayment,
			NumOfItems:  int64(cart.Amount),
			SubTotal:    int64(cart.TotalPrice),
			ShipmentFee: 10000,
			TotalFee:    int64(cart.TotalPrice) + 10000,
			Address:     userAddress,
			Shop: domain.Shop{
				ID: product.Shop.ID,
			},
			User:           user,
			ShipmentMethod: shipmentMethod,
			PaymentMethod:  paymentMethod,
		})
		if err != nil {
			return apperror.Wrap(err)
		}

		err = s.productRepository.UpdateStockByID(ctx, product.ID, newStock)
		if err != nil {
			return apperror.Wrap(err)
		}

		_, err = s.transactionItemRepository.Add(ctx, domain.TransactionItem{
			Amount:      int64(cart.Amount),
			TotalPrice:  int64(cart.TotalPrice),
			Transaction: transaction,
			Product:     product,
		})
		if err != nil {
			return apperror.Wrap(err)
		}

		err = s.cartRepository.Delete(ctx, accountID, t.ProductID)
		if err != nil {
			return apperror.Wrap(err)
		}

		ts = transaction
		return nil
	})
	if err != nil {
		return ts, apperror.Wrap(err)
	}
	return ts, nil
}

func (s *transactionService) PayTransaction(ctx context.Context, invoice string, file multipart.File) error {
	err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		t, err := s.transactionRepository.GetByInvoice(ctx, invoice)
		if err != nil {
			return apperror.Wrap(err)
		}

		if t.Status != domain.WaitingForPayment {
			return apperror.NewBadRequest(err, "invalid transaction status")
		}

		err = s.transactionRepository.UpdateStatus(ctx, t.ID, domain.WaitingForConfirmation)
		if err != nil {
			return apperror.Wrap(err)
		}

		var fileURL *string
		res, err := s.cloudinaryProvider.Upload(ctx, file, uploader.UploadParams{})
		if err != nil {
			return apperror.Wrap(err)
		}
		fileURL = &res.SecureURL

		_, err = s.paymentRepository.Add(ctx, domain.Payment{
			FileURL:     *fileURL,
			Transaction: t,
		})
		if err != nil {
			return apperror.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil

}
