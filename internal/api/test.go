package test

import (
	"github.com/labstack/echo"
	"net/http"
)

func Main(c echo.Context) error {


	return c.JSON(http.StatusOK, "Eklendi haci")

}
