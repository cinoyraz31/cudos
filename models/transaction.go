package models

import "time"

type Transaction struct {
	Id              int       `gorm:"primary_key;column:id"`
	UserId          int       `gorm:"column:user_id"`
	OrderId         string    `gorm:"column:order_id"`
	Amount          float64   `gorm:"column:amount"`
	PaymentMethod   string    `gorm:"column:payment_method"`
	Status          string    `gorm:"column:status"`
	TransactionDate time.Time `gorm:"column:transaction_date"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
