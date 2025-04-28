package models

import "time"

type User struct {
	Id            int       `gorm:"primary_key;column:id"`
	Name          string    `gorm:"column:name"`
	Email         string    `gorm:"column:email"`
	EmailVerified time.Time `gorm:"column:email_verified"`
	Password      string    `gorm:"column:password"`
	RememberToken string    `gorm:"column:remember_token"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}
