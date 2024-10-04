package variant

import (
	"basic_trade/common"

	"gorm.io/gorm"
)

type Repository interface {
	Save(variant *common.Variant) (*common.Variant, error)
	Update(variant *common.Variant) (*common.Variant, error)
	Delete(uuid string) (*common.Variant, error)
	GetAll(page, limit int, search string) ([]common.Variant, error)
	GetByUUID(uuid string) (*common.Variant, error)
	GetProductByID(productID int) (*common.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(variant *common.Variant) (*common.Variant, error) {
	err := r.db.Create(&variant).Error
	if err != nil {
		return nil, err
	}
	return variant, nil
}

func (r *repository) Update(variant *common.Variant) (*common.Variant, error) {
	err := r.db.Save(&variant).Error
	if err != nil {
		return nil, err
	}
	return variant, nil
}

func (r *repository) Delete(uuid string) (*common.Variant, error) {
	var variant common.Variant
	err := r.db.Where("uuid = ?", uuid).Delete(&variant).Error
	if err != nil {
		return nil, err
	}
	return &variant, nil
}

func (r *repository) GetAll(page, limit int, search string) ([]common.Variant, error) {
	var variants []common.Variant

	offset := (page - 1) * limit

	query := r.db.Limit(limit).Offset(offset)
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Find(&variants).Error
	if err != nil {
		return nil, err
	}
	return variants, nil
}

func (r *repository) GetByUUID(uuid string) (*common.Variant, error) {
	var variant common.Variant
	err := r.db.Where("uuid = ?", uuid).First(&variant).Error
	if err != nil {
		return nil, err
	}
	return &variant, nil
}

func (r *repository) GetProductByID(productID int) (*common.Product, error) {
	var product common.Product
	err := r.db.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
