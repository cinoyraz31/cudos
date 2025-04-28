package controllers

import "github.com/gofiber/fiber/v2"

type TransactionController interface {
	FraudDetection(ctx *fiber.Ctx) error
}
