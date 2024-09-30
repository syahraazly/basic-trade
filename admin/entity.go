package admin

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        int   `gorm:"primaryKey" json:"id"`
	UUID      string `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Name      string `gorm:"type:varchar(100);not null" validate:"required" json:"name"`
	Email     string `gorm:"type:varchar(100);unique;not null" validate:"required,email" json:"email"`
	Password  string `gorm:"not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	a.UUID = uuid.New().String()
	return
}
