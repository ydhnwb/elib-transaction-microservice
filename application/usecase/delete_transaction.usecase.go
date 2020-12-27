package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-transaction-microservice/application/repository"
	"github.com/ydhnwb/elib-transaction-microservice/infrastructure/helper"
)

//TransactionDeleteUseCase is a contract
type TransactionDeleteUseCase interface {
	DeleteTransaction(ctx *gin.Context)
}

type transactionDeleteUseCase struct {
	transactionRepository repository.TransactionRepository
}

//NewTransactionDeleteUseCase creates a new instance of DeleteUseCAse
func NewTransactionDeleteUseCase(repo repository.TransactionRepository) TransactionDeleteUseCase {
	return &transactionDeleteUseCase{
		transactionRepository: repo,
	}
}

func (ctl *transactionDeleteUseCase) DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	transaction, err := ctl.transactionRepository.FindByID(id)
	if err != nil {
		helper.BuildErrorResponse(http.StatusNotFound, err.Error(), helper.EmptyObj{}, ctx)
		return
	}

	ctl.transactionRepository.DeleteTransaction(transaction)
	helper.BuildResponse(http.StatusOK, helper.EmptyObj{}, ctx)

}
