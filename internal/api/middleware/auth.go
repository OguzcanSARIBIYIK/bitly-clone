package middleware

import (
	"bitly-clone/internal/repository"
	"bitly-clone/models"
	"net/http"

	"github.com/labstack/echo"
)

func Auth() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("token")

			var user models.User
			repository.Get().User().FindByToken(token, &user)
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
