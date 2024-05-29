package security

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type StaffSecurity struct {
	db      *pg.DB
	roles   []string
	storeId uuid.UUID
	userId  uuid.UUID
	err     error
}

func StaffSecurityConn(db *pg.DB) *StaffSecurity {
	return &StaffSecurity{db: db}
}

func (sf *StaffSecurity) Credentials(storeId uuid.UUID, userId uuid.UUID) *StaffSecurity {
	if storeId == uuid.Nil {
		sf.err = errors.New("store id is missing")
		return sf
	}

	if userId == uuid.Nil {
		sf.err = errors.New("user id is missing")
		return sf
	}

	sf.userId = userId
	sf.storeId = storeId
	return sf
}

func (sf *StaffSecurity) Role(roles ...string) *StaffSecurity {
	sf.roles = roles
	return sf
}

func (sf *StaffSecurity) Exists(storeId uuid.UUID, userId uuid.UUID) (bool, error) {
	if len(sf.roles) == 0 {
		return false, errors.New("no roles were found")
	}

	for _, role := range sf.roles {
		if role == "owner" {
			exists, err := sf.db.Model(&models.Store{}).Where("id = ? and owner_id = ?", storeId, userId).Exists()
			if err != nil {
				return false, err
			}

			if exists {
				return true, nil
			}
		}
	}

	return false, nil
}

func (sf *StaffSecurity) Select() {

}

func (sf *StaffSecurity) CreateProduct() error {
	if sf.err != nil {
		return sf.err
	}

	exists, err := sf.db.Model(&models.Store{}).Where("id = ? and owner_id = ?", sf.storeId, sf.userId).Exists()
	if err != nil && err != pg.ErrNoRows {
		return err
	}

	if !exists {
		return errors.New("not enough permissions")
	}

	return nil
}

func (sf *StaffSecurity) DeleteProduct() error {
	if sf.err != nil {
		return sf.err
	}

	exists, err := sf.db.Model(&models.Store{}).Where("id = ? and owner_id = ?", sf.storeId, sf.userId).Exists()
	if err != nil && err != pg.ErrNoRows {
		return err
	}

	if !exists {
		return errors.New("not enough permissions")
	}

	return nil
}

func (sf *StaffSecurity) CreateProductCategory() error {
	if sf.err != nil {
		return sf.err
	}

	exists, err := sf.db.Model(&models.Store{}).Where("id = ? and owner_id = ?", sf.storeId, sf.userId).Exists()
	if err != nil && err != pg.ErrNoRows {
		return err
	}

	if !exists {
		return errors.New("not enough permissions")
	}

	return nil
}
