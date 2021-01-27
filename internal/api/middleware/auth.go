package middleware

import (
	"bitly-clone/models"
	"fmt"
	"github.com/labstack/echo"
	"time"
)

func Auth() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("token")
			c.Set("user", models.User{
				ID:        0,
				Username:  "asda",
				Password:  "pass",
				Token:     "qweqwe",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			})
			fmt.Println("param : ", token)
			if token != "oguzcan" {
				panic("gg")
			}

			return handlerFunc(c)
		}
	}
}
