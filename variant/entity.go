package variant

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	ID          int   `gorm:"primaryKey"`
	UUID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	VariantName string `gorm:"type:varchar(100);not null" validate:"required"`
	Quantity    int    `gorm:"not null" validate:"required"`
	ProductID   int   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	v.UUID = uuid.New().String()
	return
}
