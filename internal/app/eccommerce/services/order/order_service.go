package order_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"

	"github.com/google/uuid"
)

func CreateOrder(cfg *config.Config, data PostOrderData) (*models.Order, error) {
	order, err := createOrder(cfg.BDB, data)
	if err != nil {
		return order, err
	}

	return order, nil
}

func GetOrders(cfg *config.Config, data OrderSearchQuery) (*map[string]interface{}, error) {
	var orders = make([]OrderData, 0)
	query := cfg.BDB.NewSelect().
		Model(&orders).
		Column("id", "store_id", "customer_id", "delivery_cost", "total_cost", "order_status", "paid", "shipping", "created_at").
		Column("stre.currency").
		Join("LEFT JOIN stores AS stre ON stre.id = ord.store_id").
		Where("store_id = ?", data.StoreId)

	if len(data.OrderId) != 0 {
		query.Where("ord.id = ?", data.OrderId)
	}

	count, err := query.Limit(25).ScanAndCount(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to get orders").Log()
	}

	return &map[string]interface{}{
		"count":  count,
		"orders": orders,
	}, nil
}

func GetOrderInfo(cfg *config.Config, id int) (interface{}, error) {
	var orderInfo models.Order
	var products []OrderProductData
	var store models.Store

	var paymentMethod = &models.StorePaymentMethod{}
	var shipment = &models.OrderShipment{}
	var deliveryMethod = &models.StoreDeliveryMethod{}

	err := cfg.BDB.NewSelect().Model(&orderInfo).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to get order info").Log()
	}

	err = cfg.BDB.NewSelect().Model(&store).Column("currency").Where("id = ?", orderInfo.StoreID).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to get order info").Log()
	}

	err = cfg.BDB.NewSelect().Model(&products).
		ModelTableExpr("order_products as order_product").
		ColumnExpr("order_product.*").
		ColumnExpr("product.images").
		ColumnExpr("product.name").
		Join("LEFT JOIN products AS product ON product.id = order_product.product_id").
		Where("order_id = ?", id).
		Scan(context.Background())

	if err != nil {
		return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to get products").Log()
	}

	if orderInfo.PaymentMethodID != uuid.Nil {
		err = cfg.BDB.NewSelect().Model(paymentMethod).Where("id = ?", orderInfo.PaymentMethodID).
			Scan(context.Background())
		if err != nil {
			return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to get payment method").Log()
		}
	}

	if orderInfo.DeliveryMethodID != uuid.Nil {
		err = cfg.BDB.NewSelect().Model(deliveryMethod).Where("id = ?", orderInfo.DeliveryMethodID).
			Scan(context.Background())
		if err != nil {
			return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to get delivery method").Log()
		}

		err = cfg.BDB.NewSelect().Model(shipment).Where("order_id = ?", orderInfo.ID).
			Scan(context.Background())
		if err != nil {
			return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to get shipment details").Log()
		}
	}

	return map[string]interface{}{
		"currency":        store.Currency,
		"order":           orderInfo,
		"products":        products,
		"payment_method":  paymentMethod,
		"delivery_method": deliveryMethod,
		"shipment":        shipment,
	}, nil
}

func UpdateOrderAddress(cfg *config.Config, data PostAddressData) error {
	if data.Id == 0 {
		return errors.New("order id is not found")
	}
	err := data.Address.Validation()
	if err != nil {
		return err
	}

	if data.AddressType == BILLING {
		_, err := cfg.BDB.NewUpdate().Model(&models.Order{
			Billing: data.Address,
		}).Column("billing").Where("id = ?", data.Id).
			Exec(context.Background())

		if err != nil {
			return err
		}

		return nil
	} else if data.AddressType == SHIPPING {
		_, err := cfg.BDB.NewUpdate().Model(&models.Order{
			Shipping: data.Address,
		}).Column("shipping").Where("id = ?", data.Id).Exec(context.Background())

		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("address type should be either billing or shipping")
}

func UpdateOrderStatus(cfg *config.Config, data PostStatusData) error {
	if data.Id == 0 {
		return errors.New("order id is not found")
	}

	order := models.Order{
		Paid:        data.Paid,
		OrderStatus: data.OrderStatus,
		PaymentDate: data.PaymentDate,
		ProcessedAt: data.ProcessedAt,
		CancelledAt: data.CancelledAt,
		CompletedAt: data.CompletedAt,
	}

	_, err := cfg.BDB.NewUpdate().Model(&order).
		Column("order_status", "paid", "payment_date", "processed_at", "cancelled_at", "completed_at").
		Where("id = ?", data.Id).Exec(context.Background())
	if err != nil {
		return middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to update order").Log()
	}

	return nil
}

func UpdateOrderShipment(cfg *config.Config, data PostShipmentData) error {
	if data.Id == uuid.Nil {
		return errors.New("shipment id is not found")
	}

	shipment := models.OrderShipment{
		Price:                 data.Price,
		Status:                data.Status,
		ShippedAt:             data.ShippedAt,
		EstimatedDeliveryDate: data.DateOfArrival,
	}

	_, err := cfg.BDB.NewUpdate().Model(&shipment).
		Column("price", "status", "shipped_at", "estimated_delivery_date").
		Where("id = ?", data.Id).Returning("order_id").Exec(context.Background())

	if err != nil {
		return middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to update shipment").Log()
	}

	_, err = cfg.BDB.NewUpdate().Model(&models.Order{
		DeliveryCost: shipment.Price,
	}).
		Column("delivery_cost").
		Where("id = ?", shipment.OrderID).Exec(context.Background())
	if err != nil {
		return middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to update shipment").Log()
	}

	return nil
}
