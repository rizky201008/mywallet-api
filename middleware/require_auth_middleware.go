package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rizky201008/mywallet-backend/exception"
	"github.com/rizky201008/mywallet-backend/model/domain"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func (middleware MiddlewareImpl) RequireAuth(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		panic(exception.NotMatchError{
			Err: "Authorization error, invalid token",
		})
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(middleware.Viper.GetString("secrets.JWT_SECRET")), nil
	})
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			panic(exception.NotMatchError{
				Err: "Authorization error, token expired",
			})
		}

		var user domain.User
		middleware.DB.Find(&user, claims["sub"])
		if user.ID == 0 {
			panic(exception.NotMatchError{
				Err: "Authorization error, invalid token",
			})
		}

		ctx.Set("username", user.Username)
		ctx.Set("id", strconv.Itoa(int(user.ID)))

		err := ctx.Next()
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
	return nil
}

func NewMiddleware(viper *viper.Viper, db *gorm.DB) Middleware {
	return MiddlewareImpl{Viper: viper, DB: db}
}
