package usecase

import (
	"context"
	"strings"
	"time"

	"stone/transaction-challenge/domain"
)

type transactionUsecase struct {
	transactionRepo domain.TransactionRepository
	contextTimeout  time.Duration
}

func NewTransactionUsecase(r domain.TransactionRepository, timeout time.Duration) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: r,
		contextTimeout:  timeout,
	}
}

func (u *transactionUsecase) Insert(authType string, t *domain.Transaction) (result map[string]interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.contextTimeout)
	defer cancel()

	result, err = u.transactionRepo.Insert(ctx, strings.ToLower(authType), t)
	return
}
