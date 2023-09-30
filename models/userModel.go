package models

import (
	"time"
)

type User struct {
	Id_u      int     `gorm:"unique"`
	Username  *string `gorm:"unique"`
	Email     *string `valid:"email"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time `valid:"-"`
}
