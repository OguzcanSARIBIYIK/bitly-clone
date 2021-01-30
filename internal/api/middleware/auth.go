package middleware

import (
	"bitly-clone/configs/db"
	"bitly-clone/models"
	"github.com/labstack/echo"
	"net/http"
)

func Auth() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("token")

			var user models.User
			db.MyDB.Model(&models.User{}).Where("token = ?", token).First(&user)
			if user.ID == 0 {
				return c.JSON(http.StatusBadRequest, models.Response{
					Code: 400,
					Data: "Geçersiz token!",
				})
			} else {
				c.Set("user", user)
				return handlerFunc(c)
			}

		}
	}
}
