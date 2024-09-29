package admin

import "gorm.io/gorm"

type Repository interface {
	Save(admin *Admin) (*Admin, error)
	FindByEmail(email string) (*Admin, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(admin *Admin) (*Admin, error) {
	err := r.db.Create(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *repository) FindByEmail(email string) (*Admin, error) {
	var admin Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
