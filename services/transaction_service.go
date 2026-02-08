package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (service *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return service.repo.CreateTransaction(items)
}
