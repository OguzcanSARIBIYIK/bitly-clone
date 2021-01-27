package middleware

import (
	"bitly-clone/models"
	"fmt"
	"github.com/labstack/echo"
)

func Auth() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("token")
			c.Set("user", models.User{})
			fmt.Println("param : ", token)
			if token != "oguzcan" {
				panic("gg")
			}

			return handlerFunc(c)
		}
	}
}
