package exception

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	notFoundErr, notFound := notFoundError(ctx, err)
	if notFound {
		return notFoundErr
	}

	notMatchErr, notMatch := notMatchError(ctx, err)
	if notMatch {
		return notMatchErr
	}

	ctx.Status(fiber.StatusInternalServerError)
	return ctx.JSON(fiber.Map{
		"status":  "1",
		"message": err.Error(),
		"data":    nil,
	})
}

func notFoundError(ctx *fiber.Ctx, err error) (error, bool) {
	var exception NotFoundError
	ok := errors.As(err, &exception)
	if ok {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"status":  "2",
			"message": exception.Error(),
			"data":    nil,
		}), true

	} else {
		return nil, false
	}
}

func notMatchError(ctx *fiber.Ctx, err error) (error, bool) {
	var exception NotMatchError
	ok := errors.As(err, &exception)
	if ok {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"status":  "3",
			"message": exception.Error(),
			"data":    nil,
		}), true

	} else {
		return nil, false
	}
}
