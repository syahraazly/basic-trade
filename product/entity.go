package product

import (
	"basic_trade/variant"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        int   `gorm:"primaryKey"`
	UUID      string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string `gorm:"type:varchar(100);not null" validate:"required"`
	ImageURL  string `gorm:"type:varchar(255);not null" validate:"required,url"`
	AdminID   int   
	CreatedAt time.Time
	UpdatedAt time.Time
	Variants  []variant.Variant `gorm:"foreignKey:ProductID"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New().String()
	return
}
