package product_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/security"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type ProductCategoryService interface {
	CreateProductCategory() (*models.ProductCategory, error)
	UpdateCategory() (*models.ProductCategory, error)
	GetProductCategories(search string, parent string) ([]models.ProductCategory, error)
	DeleteProductCategory() error
}

type ProductCategoryData struct {
	Config   *config.Config          `json:"-"`
	Category *models.ProductCategory `json:"category"`
	userId   uuid.UUID               `json:"-"`
	storeId  uuid.UUID               `json:"-"`
}

func (ct *ProductCategoryData) StaffID(param string) error {
	userId, err := uuid.Parse(param)
	if err != nil {
		return errors.New("your user information is invalid")
	}

	ct.userId = userId

	return nil
}

func (ct *ProductCategoryData) StoreID(param string) error {
	storeId, err := uuid.Parse(param)
	if err != nil {
		return errors.New("your user information is invalid")
	}

	ct.storeId = storeId

	return nil
}

func (ct *ProductCategoryData) CreateProductCategory() (*models.ProductCategory, error) {
	if ct.Category.StoreId == uuid.Nil {
		return nil, errors.New("store id is missing")
	}

	err := security.StaffSecurityConn(ct.Config.DB).Credentials(ct.Category.StoreId, ct.userId).CreateProductCategory()
	if err != nil {
		return nil, err
	}

	if len(ct.Category.CategoryName) == 0 {
		return nil, errors.New("category name not found")
	}

	exists, err := ct.Config.DB.Model(ct.Category).Where("store_id = ? and category_name = ?", ct.Category.StoreId, ct.Category.CategoryName).Exists()
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("this category already exists")
	}

	if len(ct.Category.Slug) == 0 {
		ct.Category.Slug = strings.ReplaceAll(strings.ToLower(ct.Category.CategoryName), " ", "_")
	} else {
		ct.Category.Slug = strings.ReplaceAll(strings.ToLower(ct.Category.Slug), " ", "_")
	}

	_, err = ct.Config.DB.Model(ct.Category).Insert()
	if err != nil {
		return nil, err
	}

	return ct.Category, nil
}

func (ct *ProductCategoryData) UpdateCategory() (*models.ProductCategory, error) {
	exists, err := security.StaffSecurityConn(ct.Config.DB).Role("owner").Exists(ct.Category.StoreId, ct.userId)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("you dont have enough permission")
	}

	if ct.Category.Id == 0 {
		return nil, errors.New("id is missing")
	}

	exists, err = ct.Config.DB.Model(&models.ProductCategory{}).Where("id = ?", ct.Category.Id).Exists()
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("this category does exists")
	}

	_, err = ct.Config.DB.Model(ct.Category).Where("id = ?", ct.Category.Id).Column("category_name", "slug", "parent_category_id").Update()
	if err != nil {
		return nil, err
	}

	return ct.Category, nil
}

func (ct *ProductCategoryData) GetProductCategories(search string, parent string) ([]models.ProductCategory, error) {
	var category = make([]models.ProductCategory, 0)
	q := ct.Config.DB.Model(&category).Where("store_id = ?", ct.storeId)

	if len(search) != 0 {
		q.Where(fmt.Sprintf("category_name ILIKE '%s'", "%"+search+"%"))
	}

	if len(parent) != 0 {
		i, err := strconv.Atoi(parent)
		if err != nil {
			return category, errors.New("id needs to be an integer")
		}

		q.Where("parent_category_id = ?", i)
	}

	err := q.Select()
	if err != nil {
		return category, err
	}

	return category, nil
}

func (ct *ProductCategoryData) DeleteProductCategory() error {

	if ct.Category.StoreId == uuid.Nil {
		return errors.New("store id is missing")
	}

	err := security.StaffSecurityConn(ct.Config.DB).Credentials(ct.Category.StoreId, ct.userId).CreateProductCategory()
	if err != nil {
		return err
	}

	exists, err := ct.Config.DB.Model(ct.Category).Where("id = ? and store_id = ?", ct.Category.Id, ct.Category.StoreId).Exists()
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("this category does not exists")
	}

	_, err = ct.Config.DB.Model(ct.Category).Where("id = ? and store_id = ?", ct.Category.Id, ct.Category.StoreId).Delete()
	if err != nil {
		return err
	}

	return nil
}
