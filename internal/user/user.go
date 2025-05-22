package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"` // UUID as primary key
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string
	CreatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
