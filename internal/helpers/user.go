package helpers

import (
	"bitly-clone/models"
	"github.com/labstack/echo"
)

func GetUser(c echo.Context) models.User {
	return c.Get("user").(models.User)
}
