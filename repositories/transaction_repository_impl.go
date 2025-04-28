package repositories

import (
	"cudo_task_service/models"
	"cudo_task_service/web/requests"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type TransactionRepositoryImpl struct{}

func NewTransactionRepository() *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{}
}

func (t TransactionRepositoryImpl) GetAverageUserTransaction(db *gorm.DB, UserId int, NewTransactionId int) (float64, error) {
	var amountBaseline float64
	err := db.
		Table("transactions").
		Select("AVG(amount) as avg").
		Where("user_id = ? AND id != ?", UserId, NewTransactionId).
		Debug().
		Scan(&amountBaseline).Error

	if err != nil {
		return 0, err
	}
	return amountBaseline, nil
}

func (t TransactionRepositoryImpl) UserRepeatOrderTransactionCount(db *gorm.DB, UserId int, TransactionTime time.Time) int64 {
	var count int64
	db.Table("transactions").Where("user_id = ? and transaction_date < ?", UserId, TransactionTime).
		Debug().
		Count(&count)
	return count
}

func (t TransactionRepositoryImpl) GetTransactions(db *gorm.DB, params requests.TransactionFilter) ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := db.Table("transactions").Where("status = ?", "completed")

	if params.Page != "" && params.Size != "" {
		intPage, err := strconv.Atoi(params.Page)
		if err != nil {
			intPage = 1
		}

		intPerPage, err := strconv.Atoi(params.Size)
		if err != nil {
			intPerPage = 10
		}

		offset := (intPage - 1) * intPerPage

		query = query.Offset(offset).Limit(intPerPage)
	}

	result := query.Debug().Scan(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}
