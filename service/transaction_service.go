package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/mywallet-backend/helper"
	"github.com/rizky201008/mywallet-backend/model/domain"
	"github.com/rizky201008/mywallet-backend/model/web"
	"github.com/rizky201008/mywallet-backend/repository"
	"gorm.io/gorm"
	"strconv"
)

type TransactionService interface {
	GetAllTransaction() []web.ResponseTransaction
	GetTransaction(ctx *fiber.Ctx) web.ResponseTransaction
	CreateTransaction(ctx *fiber.Ctx) web.ResponseTransaction
	UpdateTransaction(ctx *fiber.Ctx) web.ResponseTransaction
	DeleteTransaction(ctx *fiber.Ctx)
}

type TransactionServiceImpl struct {
	DB              *gorm.DB
	TransactionRepo repository.TransactionRepository
}

func NewTransactionService(db *gorm.DB, transactionRepo repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		DB:              db,
		TransactionRepo: transactionRepo,
	}
}

func (service TransactionServiceImpl) GetAllTransaction() []web.ResponseTransaction {
	var response []domain.Transaction
	data, err := service.TransactionRepo.FindAll(service.DB)
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		response = append(response, value)
	}

	return helper.TransactionsToResponseTransactions(response)
}

func (service TransactionServiceImpl) GetTransaction(ctx *fiber.Ctx) web.ResponseTransaction {
	var response domain.Transaction
	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	data, err := service.TransactionRepo.Find(service.DB, id)

	if err != nil {
		panic(err)
	}
	response = data

	return helper.TransactionToResponseTransaction(response)
}

func (service TransactionServiceImpl) CreateTransaction(ctx *fiber.Ctx) web.ResponseTransaction {
	p := new(web.RequestTransaction)
	err := ctx.BodyParser(p)
	if err != nil {
		panic(err)
	}
	var response domain.Transaction
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		transaction := helper.RequestTransactionToTransaction(*p)
		created, err := service.TransactionRepo.Create(tx, transaction)
		if err != nil {
			return err
		}
		response = created
		return nil
	})
	if err != nil {
		panic(err)
	}

	return helper.TransactionToResponseTransaction(response)
}

func (service TransactionServiceImpl) UpdateTransaction(ctx *fiber.Ctx) web.ResponseTransaction {
	p := new(web.RequestTransaction)
	err := ctx.BodyParser(p)
	if err != nil {
		panic(err)
	}
	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	var response domain.Transaction
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		transaction := helper.RequestTransactionToTransaction(*p)
		updated, err := service.TransactionRepo.Update(tx, transaction, id)
		if err != nil {
			return err
		}

		response = updated
		return nil
	})
	if err != nil {
		panic(err)
	}

	return helper.TransactionToResponseTransaction(response)
}

func (service TransactionServiceImpl) DeleteTransaction(ctx *fiber.Ctx) {
	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	err = service.TransactionRepo.Delete(service.DB, id)
	if err != nil {
		panic(err)
	}
}
