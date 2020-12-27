package repository

import (
	"github.com/ydhnwb/elib-transaction-microservice/domain/entity"
	"gorm.io/gorm"
)

//TransactionRepository is a contract
type TransactionRepository interface {
	CreateTransaction(t entity.Transaction) (entity.Transaction, error)
	DeleteTransaction(t entity.Transaction)
	FindByID(transactionID string) (entity.Transaction, error)
	AllTransaction() []entity.Transaction
}

type transactionRepository struct {
	db *gorm.DB
}

//NewTransactionRepository creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) CreateTransaction(t entity.Transaction) (entity.Transaction, error) {
	err := r.db.Save(&t).Error
	if err != nil {
		return t, err
	}
	r.db.Preload("Student").Preload("Book").Find(&t, t.ID).Take(&t)
	return t, nil
}

func (r *transactionRepository) DeleteTransaction(t entity.Transaction) {
	r.db.Delete(t)
}

func (r *transactionRepository) AllTransaction() []entity.Transaction {
	transactions := []entity.Transaction{}
	r.db.Preload("Student").Preload("Book").Find(&transactions)
	return transactions
}

func (r *transactionRepository) FindByID(transactionID string) (entity.Transaction, error) {
	transaction := entity.Transaction{}
	err := r.db.Preload("Student").Preload("Book").Find(&transaction, transactionID).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
