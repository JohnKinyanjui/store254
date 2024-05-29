package public_cart_service

import (
	"context"
	"database/sql"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func createNewCart(db *bun.DB, customerId uuid.UUID, storeId uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	count, err := db.NewSelect().Model(&cart).
		Where("customer_id = ? and store_id = ?", customerId, storeId).
		ScanAndCount(context.Background())

	if err != nil && err != sql.ErrNoRows {
		return nil, middlewares.FLog(middlewares.CART_SERVICE_LOG, err, "unable to check cart").Log()
	}

	if count == 0 {
		var cart = models.Cart{
			StoreId:    storeId,
			CustomerID: customerId,
			CreatedAt:  time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
		}

		_, err = db.NewInsert().Model(&cart).
			Column("id", "store_id", "customer_id", "created_at", "updated_at").
			Exec(context.Background())

		if err != nil {
			return nil, middlewares.FLog(middlewares.CART_SERVICE_LOG, err, "unable to create new cart").Log()
		}
		return &cart, nil
	}

	return &cart, nil
}

func addNewCartItem(db *bun.DB, cartId uuid.UUID, data CartPostData) error {
	var cartItem models.CartItem

	count, err := db.NewSelect().Model(&cartItem).
		Where("cart_id = ? and product_id = ?", cartId, data.ProductId).
		ScanAndCount(context.Background())

	if err != nil && err != sql.ErrNoRows {
		return middlewares.FLog(middlewares.CART_SERVICE_LOG, err, "unable to check cart").Log()
	}

	if count == 0 {
		if data.Add {
			_, err = db.NewInsert().Model(&models.CartItem{
				CartID:    cartId,
				ProductID: data.ProductId,
				VariantID: data.VariantId,
				Quantity:  1,
				CreatedAt: time.Now().UTC(),
			}).
				Column("id", "cart_id", "product_id", "variant_id", "quantity", "created_at").
				Exec(context.Background())
			if err != nil {
				return middlewares.FLog(middlewares.CART_SERVICE_LOG, err, "unable to update cart item").Log()
			}
		}
	} else {
		if data.Add {
			cartItem.Quantity = cartItem.Quantity + 1
			_, err = db.NewUpdate().Model(&cartItem).
				Column("quantity").
				Where("cart_id = ? and product_id = ?", cartId, data.ProductId).
				Exec(context.Background())
			if err != nil {
				return middlewares.FLog(middlewares.CART_SERVICE_LOG, err, "unable to update cart item").Log()
			}
		} else {
			cartItem.Quantity = cartItem.Quantity - 1
			if cartItem.Quantity > 0 {
				_, err = db.NewUpdate().Model(&cartItem).
					Column("quantity").
					Where("cart_id = ? and product_id = ?", cartId, data.ProductId).
					Exec(context.Background())
				if err != nil {
					return middlewares.FLog(middlewares.CART_SERVICE_LOG, err, "unable to update cart item").Log()
				}
			}

			if data.Remove {
				_, err = db.NewDelete().Model(&cartItem).
					Where("cart_id = ? and product_id = ?", cartId, data.ProductId).
					Exec(context.Background())
				if err != nil {
					return middlewares.FLog(middlewares.CART_SERVICE_LOG, err, "unable to update cart item").Log()
				}
			}
		}
	}

	return nil
}

func getCustomerCartInfo(db *bun.DB, customerID uuid.UUID) (*CartInfoResponse, error) {
	ctx := context.Background()

	// Find the customer's cart
	var cart models.Cart
	err := db.NewSelect().Model(&cart).Where("customer_id = ?", customerID).Limit(1).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return &CartInfoResponse{
				TotalCost: 0,
				Count:     0,
				Products:  make([]CartProductInfo, 0),
			}, nil
		}
		return nil, err
	}

	// Query for cart items and corresponding product details
	var cartItems []models.CartItem
	err = db.NewSelect().Model(&cartItems).Where("cart_id = ?", cart.ID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	response := &CartInfoResponse{}
	for _, item := range cartItems {

		var product models.Product
		err = db.NewSelect().Model(&product).Where("id = ?", item.ProductID).Limit(1).Scan(ctx)
		if err != nil {
			return nil, err
		}
		price := product.RegularPrice

		if product.SalesPrice != 0 {
			price = product.SalesPrice
		}

		var variantName string
		if item.VariantID != uuid.Nil {
			var variant models.ProductVariant
			err = db.NewSelect().Model(&variant).Where("id = ?", item.VariantID).Limit(1).Scan(ctx)
			if err != nil {
				return nil, err
			}

			variantName = variant.Name
			price = variant.Price
		}

		firstImage := ""
		if len(product.Images) > 0 {
			firstImage = product.Images[0]
		}

		response.Products = append(response.Products, CartProductInfo{
			ID:       product.Id,
			Image:    firstImage,
			Name:     product.Name,
			Variant:  variantName,
			Price:    price,
			Quantity: item.Quantity,
		})

		response.TotalCost += price * float64(item.Quantity)
	}

	if len(response.Products) == 0 {
		response.Products = make([]CartProductInfo, 0)
	}

	response.Count = len(response.Products)

	return response, nil
}
