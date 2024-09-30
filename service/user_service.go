package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/mywallet-backend/exception"
	"github.com/rizky201008/mywallet-backend/helper"
	"github.com/rizky201008/mywallet-backend/model/web"
	"github.com/rizky201008/mywallet-backend/repository"
	"gorm.io/gorm"
	"strconv"
)

type UserService interface {
	CreateUser(ctx *fiber.Ctx) web.ResponseUser
	UpdateUser(ctx *fiber.Ctx) web.ResponseUser
	FindUserById(ctx *fiber.Ctx) web.ResponseUser
	DeleteUser(ctx *fiber.Ctx)
}

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *gorm.DB
}

func NewUserService(userRepo repository.UserRepository, db *gorm.DB) UserService {
	return UserServiceImpl{UserRepo: userRepo, DB: db}
}

func (service UserServiceImpl) CreateUser(ctx *fiber.Ctx) web.ResponseUser {
	p := new(web.RequestUser)
	err := ctx.BodyParser(p)
	if err != nil {
		panic(err)
	}
	var response web.ResponseUser
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		user, err := service.UserRepo.CreateUser(service.DB, helper.RequestUserToUser(*p))
		if err != nil {
			return err
		}
		response = helper.UserToResponseUser(user)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return response
}

func (service UserServiceImpl) UpdateUser(ctx *fiber.Ctx) web.ResponseUser {
	p := new(web.RequestUser)
	err := ctx.BodyParser(p)
	if err != nil {
		panic(err)
	}
	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	var response web.ResponseUser
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		foundedUser, err := service.UserRepo.FindUserById(tx, id)
		if err != nil {
			return exception.NotFoundError{
				Err: err.Error(),
			}
		}

		foundedUser.Username = p.Username
		updatedUser, err := service.UserRepo.UpdateUser(tx, foundedUser)
		if err != nil {
			return err
		}
		response = helper.UserToResponseUser(updatedUser)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return response
}

func (service UserServiceImpl) FindUserById(ctx *fiber.Ctx) web.ResponseUser {
	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	foundedUser, err := service.UserRepo.FindUserById(service.DB, id)
	if err != nil {
		panic(exception.NotFoundError{
			Err: err.Error(),
		})
	}
	return helper.UserToResponseUser(foundedUser)
}

func (service UserServiceImpl) DeleteUser(ctx *fiber.Ctx) {
	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	foundedUser, err := service.UserRepo.FindUserById(service.DB, id)
	if err != nil {
		panic(exception.NotFoundError{
			Err: err.Error(),
		})
	}

	err = service.UserRepo.DeleteUser(service.DB, foundedUser)
	if err != nil {
		panic(err)
	}
}
