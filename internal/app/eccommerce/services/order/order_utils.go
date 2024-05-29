package order_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func createOrder(db *bun.DB, data PostOrderData) (*models.Order, error) {
	products, totalPrice, err := createOrderProducts(db, data, data.CustomerID)
	if err != nil {
		return nil, err
	}

	order, err := validateOrderDetails(db, data, totalPrice)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	query := tx.
		NewInsert().Model(order).
		Column("id", "store_id", "total_cost", "note", "paid", "billing", "shipping", "payment_date", "created_at")

	if order.CustomerID != uuid.Nil {
		query = query.Column("customer_id")
	}

	if order.DeliveryMethodID != uuid.Nil {
		query = query.Column("delivery_method_id")
	}

	if order.PaymentMethodID != uuid.Nil {
		query = query.Column("payment_method_id")
	}

	_, err = query.Exec(context.Background())
	if err != nil {
		return nil, err
	}

	for _, prod := range products {
		prod.OrderID = order.ID
		query1 := tx.NewInsert().Model(&prod).Column("order_id", "store_id", "product_id", "quantity", "price", "total_price", "created_at")
		if prod.CustomerID != uuid.Nil {
			query1 = query1.Column("customer_id")
		}

		_, err := query1.Exec(context.Background())
		if err != nil {
			return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to save shipment details").Log()
		}
	}

	if order.PaymentMethodID != uuid.Nil {
		data.PaymentDetails["amount"] = strconv.FormatFloat(order.TotalCost, 'f', 2, 64)
		err = initiatePayment(db, data)
		if err != nil {
			return nil, err
		}
	}

	if order.DeliveryMethodID != uuid.Nil {
		shipment, err := getShipmemnt(db, order, order.DeliveryMethodID)
		if err != nil {
			return nil, err
		}

		shipment.OrderID = order.ID
		_, err = tx.NewInsert().
			Model(shipment).
			Column("id", "order_id", "delivery_method_id", "tracking_number", "estimated_delivery_date").
			Column("status", "created_at", "updated_at").
			Exec(context.Background())
		if err != nil {
			return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to save shipment details").Log()
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func createOrderProducts(db *bun.DB, data PostOrderData, customerId string) ([]models.OrderProduct, float64, error) {
	var totalCost float64
	var orderProducts = make([]models.OrderProduct, 0)
	customrId := uuid.Nil

	if len(customerId) != 0 {
		customerId, err := uuid.Parse(data.CustomerID)
		if err != nil {
			return orderProducts, totalCost, errors.New("delivery id is a uuid not a string")
		}

		customrId = customerId
	}

	if len(data.OrderProducts) == 0 {
		return orderProducts, totalCost, errors.New("no products were added")
	}

	for _, pd := range data.OrderProducts {
		var product Product
		var price float64

		err := db.NewSelect().Model(&product).Column("id", "regular_price", "sales_price").Where("id = ?", pd.ProductID).Scan(context.Background())
		if err != nil {
			return orderProducts, totalCost, err
		}

		if product.SalesPrice > 0 {
			price = product.SalesPrice
		} else {
			price = product.RegularPrice
		}

		orderProducts = append(orderProducts, models.OrderProduct{
			CustomerID: customrId,
			StoreID:    data.StoreID,
			ProductID:  pd.ProductID,
			Quantity:   pd.Quantity,
			Price:      price,
			TotalPrice: price * float64(pd.Quantity),
			CreatedAt:  time.Now().UTC(),
		})

		totalCost = totalCost + (price * float64(pd.Quantity))
	}

	return orderProducts, totalCost, nil
}
