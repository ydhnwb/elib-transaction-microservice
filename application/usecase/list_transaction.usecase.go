package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-transaction-microservice/application/repository"
	"github.com/ydhnwb/elib-transaction-microservice/infrastructure/helper"
)

//TransactionListUseCase is a contract
type TransactionListUseCase interface {
	AllTransaction(ctx *gin.Context)
}

type transactionListUseCase struct {
	transactionRepository repository.TransactionRepository
}

//NewTransactionListUseCase creates a new instance of TransactionListUseCase
func NewTransactionListUseCase(repo repository.TransactionRepository) TransactionListUseCase {
	return &transactionListUseCase{
		transactionRepository: repo,
	}
}

func (ctl *transactionListUseCase) AllTransaction(ctx *gin.Context) {
	transactions := ctl.transactionRepository.AllTransaction()
	helper.BuildResponse(http.StatusOK, transactions, ctx)
}
