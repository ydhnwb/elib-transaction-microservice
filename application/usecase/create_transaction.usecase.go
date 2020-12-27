package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/ydhnwb/elib-transaction-microservice/application/repository"
	"github.com/ydhnwb/elib-transaction-microservice/domain/entity"
	"github.com/ydhnwb/elib-transaction-microservice/infrastructure/dto"
	"github.com/ydhnwb/elib-transaction-microservice/infrastructure/helper"
)

//TransactionCreateUseCase is a contract
type TransactionCreateUseCase interface {
	CreateTransaction(ctx *gin.Context)
}

type transactionCreateUseCase struct {
	transactionRepository repository.TransactionRepository
}

//NewTransactionCreateUseCase creates a new instance of TransactionCreateUseCase
func NewTransactionCreateUseCase(repo repository.TransactionRepository) TransactionCreateUseCase {
	return &transactionCreateUseCase{
		transactionRepository: repo,
	}
}

func (ctl *transactionCreateUseCase) CreateTransaction(ctx *gin.Context) {
	transactionCreateDTO := dto.TransactionCreateDTO{}
	err := ctx.ShouldBind(&transactionCreateDTO)
	if err != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, err.Error(), helper.EmptyObj{}, ctx)
		return
	}

	transaction := entity.Transaction{}
	err = smapping.FillStruct(&transaction, smapping.MapFields(&transactionCreateDTO))
	if err != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, err.Error(), helper.EmptyObj{}, ctx)
		return
	}

	res, e := ctl.transactionRepository.CreateTransaction(transaction)
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, ctx)
		return
	}

	helper.BuildResponse(http.StatusCreated, res, ctx)
}
