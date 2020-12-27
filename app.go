package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-transaction-microservice/application/repository"
	"github.com/ydhnwb/elib-transaction-microservice/application/usecase"
	"github.com/ydhnwb/elib-transaction-microservice/infrastructure/persistence"
	"gorm.io/gorm"
)

var (
	db                       *gorm.DB                         = persistence.SetupDatabaseConnection()
	transactionRepository    repository.TransactionRepository = repository.NewTransactionRepository(db)
	transactionCreateUseCase usecase.TransactionCreateUseCase = usecase.NewTransactionCreateUseCase(transactionRepository)
	transactionDeleteUseCase usecase.TransactionDeleteUseCase = usecase.NewTransactionDeleteUseCase(transactionRepository)
	transactionListUseCase   usecase.TransactionListUseCase   = usecase.NewTransactionListUseCase(transactionRepository)
)

func main() {
	defer persistence.CloseDatabaseConnection(db)
	r := gin.Default()

	transactionRoutes := r.Group("api/transactions")
	{
		transactionRoutes.GET("/", transactionListUseCase.AllTransaction)
		transactionRoutes.POST("/", transactionCreateUseCase.CreateTransaction)
		transactionRoutes.DELETE("/:id", transactionDeleteUseCase.DeleteTransaction)
	}

	r.Run(":8083")
}
