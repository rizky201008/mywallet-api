package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rizky201008/mywallet-backend/exception"
	"github.com/rizky201008/mywallet-backend/helper"
	"github.com/rizky201008/mywallet-backend/model/web"
	"github.com/rizky201008/mywallet-backend/repository"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserService interface {
	CreateUser(ctx *fiber.Ctx) web.ResponseUser
	UpdateUser(ctx *fiber.Ctx) web.ResponseUser
	FindUserById(ctx *fiber.Ctx) web.ResponseUser
	DeleteUser(ctx *fiber.Ctx)
	Login(ctx *fiber.Ctx) web.ResponseLogin
	GetBalance(ctx *fiber.Ctx) web.ResponseUserBalance
}

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *gorm.DB
	Viper    *viper.Viper
}

func NewUserService(db *gorm.DB, userRepo repository.UserRepository, vipers *viper.Viper) UserService {
	return UserServiceImpl{UserRepo: userRepo, DB: db, Viper: vipers}
}

func (service UserServiceImpl) GetBalance(ctx *fiber.Ctx) web.ResponseUserBalance {
	idString := ctx.GetRespHeader("Id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	var response web.ResponseUserBalance
	data, err := service.UserRepo.TotalBalance(service.DB, id)
	if err != nil {
		panic(err)
	}
	var balance float64
	if data != nil {
		balance = *data
	} else {
		balance = 0
	}
	response = web.ResponseUserBalance{
		CurrentBalance: balance,
	}
	return response
}

func (service UserServiceImpl) Login(ctx *fiber.Ctx) web.ResponseLogin {
	p := new(web.RequestUser)
	err := ctx.BodyParser(p)
	if err != nil {
		panic(err)
	}
	foundedUser, err := service.UserRepo.FindUserByUsername(service.DB, p.Username)
	if err != nil {
		panic(exception.NotFoundError{
			Err: err.Error(),
		})
	}

	notMatch := bcrypt.CompareHashAndPassword([]byte(foundedUser.Password), []byte(p.Password))
	if notMatch != nil {
		panic(exception.NotMatchError{
			Err: notMatch.Error(),
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": foundedUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	secret := service.Viper.GetString("secrets.JWT_SECRET")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return web.ResponseLogin{
		Token: tokenString,
	}
}

func (service UserServiceImpl) CreateUser(ctx *fiber.Ctx) web.ResponseUser {
	p := new(web.RequestUser)
	err := ctx.BodyParser(p)
	if err != nil {
		panic(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	p.Password = string(hashedPassword)
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
