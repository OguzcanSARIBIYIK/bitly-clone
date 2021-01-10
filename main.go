package main

import (
	test "bitly-clone/internal/api"
	"bitly-clone/internal/api/user"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/",test.Main)

	e.POST("/register", user.Register)
	e.POST("/show/token", user.GetToken)

	e.Logger.Fatal(e.Start(":8000"))
}
