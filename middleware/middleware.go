package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Middleware interface {
	RequireAuth(c *fiber.Ctx) error
}

type MiddlewareImpl struct {
	Viper *viper.Viper
	DB    *gorm.DB
}
