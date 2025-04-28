package controllers

import (
	"cudo_task_service/exceptions"
	"cudo_task_service/helpers"
	"cudo_task_service/models"
	"cudo_task_service/repositories"
	"cudo_task_service/web/requests"
	"cudo_task_service/web/responses"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sync"
)

type TransactionControllerImpl struct {
	DB                    *gorm.DB
	TransactionRepository repositories.TransactionRepository
}

func NewTransactionController(
	DB *gorm.DB,
	transactionRepository repositories.TransactionRepository,
) *TransactionControllerImpl {
	return &TransactionControllerImpl{
		DB:                    DB,
		TransactionRepository: transactionRepository,
	}
}

func (t TransactionControllerImpl) FraudDetection(ctx *fiber.Ctx) error {
	data := requests.TransactionFilter{
		Page: ctx.Query("page", "1"),
		Size: ctx.Query("size", "10"),
	}
	fmt.Println(data)

	transactions, err := t.TransactionRepository.GetTransactions(t.DB, data)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Gagal get data transcations")
	}

	var wg sync.WaitGroup
	transactionLists := make([]responses.TransactionResponse, len(transactions))

	for i, transaction := range transactions {
		wg.Add(1)

		go func(i int, transaction models.Transaction) {
			defer wg.Done()

			chanFreqScore := make(chan string, 1)
			chanFreqNarration := make(chan string, 1)
			chanPatternScore := make(chan float64, 1)
			chanPatternPercentage := make(chan float64, 1)

			var innerWg sync.WaitGroup
			innerWg.Add(2)

			// Frequency check
			go func() {
				defer innerWg.Done()
				userRepeatTransactionCount := t.TransactionRepository.UserRepeatOrderTransactionCount(
					t.DB,
					transaction.UserId,
					transaction.TransactionDate,
					transaction.Id,
				)
				score, narration := helpers.FrequencyCheck(userRepeatTransactionCount)
				chanFreqScore <- score
				chanFreqNarration <- narration
			}()

			// Pattern check
			go func() {
				defer innerWg.Done()
				newAmount := transaction.Amount
				baseline, _ := t.TransactionRepository.GetAverageUserTransaction(t.DB, transaction.UserId, transaction.Id)
				percentage, score := helpers.PatternCheck(baseline, newAmount)
				chanPatternScore <- score
				chanPatternPercentage <- percentage
			}()

			innerWg.Wait()

			close(chanFreqScore)
			close(chanFreqNarration)
			close(chanPatternScore)
			close(chanPatternPercentage)

			freqScore := <-chanFreqScore
			freqNarration := <-chanFreqNarration
			patternScore := <-chanPatternScore
			patternPercentage := <-chanPatternPercentage

			transactionLists[i].TransactionId = transaction.Id
			transactionLists[i].DetectionResult.FrequencyCheck.Score = freqScore
			transactionLists[i].DetectionResult.FrequencyCheck.Narration = freqNarration
			transactionLists[i].DetectionResult.PatternCheck.Score = patternScore
			transactionLists[i].DetectionResult.PatternCheck.Percentage = patternPercentage
		}(i, transaction)
	}
	wg.Wait()
	return ctx.Status(fiber.StatusOK).JSON(transactionLists)
}
