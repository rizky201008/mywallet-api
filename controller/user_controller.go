package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/mywallet-backend/service"
)

type UserController interface {
	Register(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	UserService service.UserService
}

func (controller UserControllerImpl) Register(ctx *fiber.Ctx) error {
	response := controller.UserService.CreateUser(ctx)
	return ctx.JSON(fiber.Map{
		"status":  "0",
		"message": "success",
		"data":    response,
	})
}

func (controller UserControllerImpl) Update(ctx *fiber.Ctx) error {
	response := controller.UserService.UpdateUser(ctx)
	return ctx.JSON(fiber.Map{
		"status":  "0",
		"message": "success",
		"data":    response,
	})
}

func (controller UserControllerImpl) Login(ctx *fiber.Ctx) error {
	response := controller.UserService.Login(ctx)
	return ctx.JSON(fiber.Map{
		"status":  "0",
		"message": "success",
		"data": fiber.Map{
			"token": response,
		},
	})
}

func NewUserController(userService service.UserService) UserController {
	return UserControllerImpl{UserService: userService}
}
