package repositories

import (
	"cudo_task_service/models"
	"cudo_task_service/web/requests"
	"gorm.io/gorm"
	"time"
)

type UserWithTransaction struct {
	Name             string
	TotalTransaction int
}

type TransactionRepository interface {
	GetTransactions(db *gorm.DB, params requests.TransactionFilter) ([]models.Transaction, error)
	UserRepeatOrderTransactionCount(db *gorm.DB, UserId int, TransactionTime time.Time, NewTransactionId int) int64
	GetAverageUserTransaction(db *gorm.DB, UserId int, NewTransactionId int) (float64, error)
}
