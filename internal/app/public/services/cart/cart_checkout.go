package public_cart_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	order_service "eccomerce_apis/internal/app/eccommerce/services/order"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"log"

	"github.com/google/uuid"
)

func CheckOut(cfg *config.Config, data order_service.PostOrderData) (*models.Order, error) {
	var cart models.Cart
	var cartProducts = make([]models.CartItem, 0)
	var orderProducts = make([]order_service.PostOrderProductData, 0)

	if data.StoreID == uuid.Nil {
		return nil, errors.New("store id is missing")
	}

	err := cfg.BDB.NewSelect().Model(&cart).
		Where("customer_id = ? and store_id = ?", data.CustomerID, data.StoreID).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	err = cfg.BDB.NewSelect().Model(&cartProducts).
		Where("cart_id = ?", cart.ID).
		Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.CART_SERVICE_LOG, err, "unable to get cart products").Log()
	}

	for _, cProduct := range cartProducts {
		orderProducts = append(orderProducts, order_service.PostOrderProductData{
			ProductID: cProduct.ProductID,
			VariantID: cProduct.VariantID,
			Quantity:  cProduct.Quantity,
		})
	}

	data.OrderProducts = orderProducts
	log.Println(data)
	order, err := order_service.CreateOrder(cfg, data)
	if err != nil {
		return nil, err
	}

	go func() {
		_, err := cfg.BDB.NewDelete().Model(&[]models.CartItem{}).Where("cart_id = ?", cart.ID).Exec(context.Background())
		if err != nil {
			middlewares.FLog(middlewares.CART_SERVICE_LOG, err, "unable to get cart products").Log()
		}
	}()

	return order, nil
}
