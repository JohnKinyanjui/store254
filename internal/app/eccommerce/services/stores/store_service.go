package store_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"

	"github.com/go-pg/pg/v10"
)

func RegisterStore(cfg *config.Config, data StorePostData) error {
	return registerStore(cfg, data)
}

func GetStoreSearch(db *pg.DB, skip int, params StoreSearchParam) (*[]models.Store, error) {
	var stores = make([]models.Store, 0)

	p := db.Model(&stores)
	if len(params.Name) != 0 {
		p.Where("(name ILIKE '%' || ? || '%')", params.Name, params.Name)
	}

	if len(params.Collection) != 0 {
		p.Where("collection && ?", pg.Array(params.Collection))
	}

	if len(params.StoreId) != 0 {
		p.Where("store_id = ?", params.StoreId)
	}

	err := p.Offset(skip).Limit(15).Select()
	if err != nil {
		return nil, err
	}

	return &stores, nil

}
