package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        int    `gorm:"primaryKey"`
	UUID      string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string `gorm:"type:varchar(100);not null" validate:"required"`
	ImageURL  string `gorm:"type:varchar(255);not null" validate:"required,url"`
	AdminID   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Variants  []Variant `gorm:"foreignKey:ProductID"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New().String()
	return
}

type Variant struct {
	ID          int    `gorm:"primaryKey"`
	UUID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	VariantName string `gorm:"type:varchar(100);not null" validate:"required"`
	Quantity    int    `gorm:"not null" validate:"required"`
	ProductID   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	v.UUID = uuid.New().String()
	return
}
