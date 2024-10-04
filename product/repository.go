package product

import (
	"basic_trade/common"

	"gorm.io/gorm"
)

type Repository interface {
	Save(product *common.Product) (*common.Product, error)
	Update(product *common.Product) (*common.Product, error)
	Delete(uuid string) (*common.Product, error)
	GetAll(offset, limit int, search string) ([]common.Product, error)
	GetByUUID(uuid string) (*common.Product, error)
	GetByAdminID(adminID int, offset, limit int, search string) ([]common.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product *common.Product) (*common.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) Update(product *common.Product) (*common.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) Delete(uuid string) (*common.Product, error) {
	var product common.Product
	err := r.db.Where("uuid = ?", uuid).Delete(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) GetAll(offset, limit int, search string) ([]common.Product, error) {
	var products []common.Product

	query := r.db.Preload("Variants").Limit(limit).Offset(offset)
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) GetByUUID(uuid string) (*common.Product, error) {
	var product common.Product
	err := r.db.Where("uuid = ?", uuid).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) GetByAdminID(adminID int, offset, limit int, search string) ([]common.Product, error) {
	var products []common.Product

	query := r.db.Where("admin_id = ?", adminID).Limit(limit).Offset(offset)
	if search != "" {
		query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
