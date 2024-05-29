package public_cart_service

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func GetCartInformation(db *bun.DB, customerId uuid.UUID) (*CartInfoResponse, error) {
	if customerId == uuid.Nil {
		return nil, errors.New("customer id not found")
	}

	return getCustomerCartInfo(db, customerId)
}

func UpdateCartItem(db *bun.DB, data CartPostData, storeId uuid.UUID, customerId uuid.UUID) (*CartInfoResponse, error) {
	if customerId == uuid.Nil {
		return nil, errors.New("customer id not found")
	}

	cart, err := createNewCart(db, customerId, storeId)
	if err != nil {
		return nil, err
	}

	log.Println("cart id is ", cart.ID)
	err = addNewCartItem(db, cart.ID, data)
	if err != nil {
		return nil, err
	}

	return getCustomerCartInfo(db, customerId)
}
