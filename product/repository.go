package product

import "gorm.io/gorm"

type Repository interface {
	Save(product *Product) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(uuid string) (*Product, error)
	GetAll() ([]Product, error)
	GetByUUID(uuid string) (*Product, error)
	GetByAdminID(adminID int) ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product *Product) (*Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) Update(product *Product) (*Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) Delete(uuid string) (*Product, error) {
	var product Product
	err := r.db.Where("uuid = ?", uuid).Delete(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) GetByUUID(uuid string) (*Product, error) {
	var product Product
	err := r.db.Where("uuid = ?", uuid).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) GetByAdminID(adminID int) ([]Product, error) {
	var products []Product
	err := r.db.Where("admin_id = ?", adminID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
