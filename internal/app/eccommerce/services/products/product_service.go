package product_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"mime/multipart"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type ProductService interface {
	UpdateProduct(userId string) (*models.Product, error)
	DeleteProduct() error
}

type ProductServiceData struct {
	Config  *config.Config `json:"-"`
	Product models.Product `json:"product"`
	Files   struct {
		Images            []multipart.File
		ImageHeaders      []*multipart.FileHeader
		Attachments       []multipart.File
		AttachmentHeaders []*multipart.FileHeader
	} `json:"-"`

	userId uuid.UUID
	err    error
}

func (ps *ProductServiceData) UserId(id string) {
	userId, err := uuid.Parse(id)
	if err != nil {
		ps.err = errors.New("unable to parse user")
		return
	}

	ps.userId = userId
}

func CreateProduct(cfg config.Config, data ProductPostData) (*models.Product, error) {
	product, err := validateProductData(data)
	if err != nil {
		return nil, err
	}

	tx, err := cfg.BDB.Begin()
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to create product").Log()
	}

	_, err = tx.NewInsert().Model(&product).Exec(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to create product").Log()
	}

	data.Inventory.ProductID = product.Id
	data.Inventory.CreatedAt = time.Now().UTC()
	data.Inventory.UpdatedAt = time.Now().UTC()

	// add products
	_, err = tx.NewInsert().Model(&data.Inventory).Exec(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to create product").Log()
	}

	// variants
	for _, variant := range data.Variants {
		variant.ProductId = product.Id
		_, err = tx.NewInsert().Model(&variant).Exec(context.Background())
		if err != nil {
			return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to create product").Log()
		}
	}

	// add products categories
	for _, ct := range data.Categories {
		var category models.ProductRelCategory
		category.CategoryId = ct
		category.ProductId = product.Id

		_, err = tx.NewInsert().Model(&category).Exec(context.Background())
		if err != nil {
			return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to create product").Log()
		}

	}

	// commit ptoducts
	err = tx.Commit()
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to create product").Log()
	}

	return &product, nil

}

func UpdateProduct(cfg config.Config, data ProductPostData) (*models.Product, error) {
	product, err := validateProductData(data)
	if err != nil {
		return nil, err
	}

	tx, err := cfg.BDB.Begin()
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to update product").Log()
	}

	_, err = tx.NewUpdate().Model(&product).
		Set("images = ?, name = ?, description = ?, regular_price = ?, sales_price = ?, active = ?, downloadable = ?",
			pgdialect.Array(product.Images), product.Name, product.Description, product.RegularPrice, product.SalesPrice, product.Active, product.Downloadable).
		Where("id = ?", data.Info.Id).
		Returning("*").
		Exec(context.Background())

	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to update product").Log()
	}

	data.Inventory.ProductID = product.Id
	data.Inventory.CreatedAt = time.Now().UTC()
	data.Inventory.UpdatedAt = time.Now().UTC()

	// add products
	data.Inventory.UpdatedAt = time.Now()
	_, err = tx.NewUpdate().Model(&data.Inventory).
		Column("quantity", "minimum_quantity", "status", "unlimited_stock", "updated_at").
		Where("id = ?", data.Inventory.ID).
		Exec(context.Background())
	if err != nil {

		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to update product").Log()
	}

	// delete variants
	for _, variantId := range data.RemoveVariants {
		_, err = tx.NewDelete().Model(&models.ProductVariant{}).
			Where("id = ?", variantId).
			Exec(context.Background())
		if err != nil {
			return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to update product").Log()
		}

	}

	// variants
	for _, variant := range data.Variants {
		if variant.Id == uuid.Nil {
			variant.ProductId = product.Id
			_, err = tx.NewInsert().Model(&variant).Exec(context.Background())
			if err != nil {
				return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to update product").Log()
			}
		} else {
			_, err = tx.NewUpdate().Model(&variant).
				Column("name", "price", "stock_level", "attributes").
				Where("id = ?", variant.Id).
				Exec(context.Background())
			if err != nil {
				return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to update product").Log()
			}
		}
	}

	// add products categories

	err = UpdateProductCategories(cfg.DB, data.Info.Id, data.Categories)
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to update product").Log()
	}

	// commit ptoducts
	err = tx.Commit()
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to update product").Log()
	}

	return &product, nil

}
func (ps *ProductServiceData) DeleteProduct(cfg config.Config, data ProductPostData) error {
	if ps.Product.Id == uuid.Nil {
		return errors.New("id was not found")
	}

	if ps.Product.StoreId == uuid.Nil {
		return errors.New("store id was not found")
	}

	if ps.Product.Id == uuid.Nil {
		return errors.New("product id was not found")
	}

	err := ps.Config.DB.Model(&models.Product{}).Where("id = ?", ps.Product.Id).WhereOr("id = ? and deleted = true", ps.Product.Id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return errors.New("seems the product is already deleted")
		}

		return err
	}
	exists, err := ps.Config.DB.Model(&[]models.OrderProduct{}).Where("product_id = ? and store_id = ?", ps.Product.Id, ps.Product.StoreId).Exists()
	if err != nil {
		return err
	}

	if exists {
		_, err = ps.Config.DB.Model(&models.Product{Deleted: true, DeletedAt: time.Now().UTC()}).Column("deleted", "deleted_at").Where("id = ? and store_id = ?", ps.Product.Id, ps.Product.StoreId).Update()
		if err != nil {
			return err
		}
	} else {
		_, err = ps.Config.DB.Model(&models.Product{}).Where("id = ? and store_id = ?", ps.Product.Id, ps.Product.StoreId).Delete()
		if err != nil {
			return err
		}
	}

	return nil
}
