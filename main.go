package main

import (
	test "bitly-clone/internal/api"
	"bitly-clone/internal/api/link"
	"bitly-clone/internal/api/middleware"
	"bitly-clone/internal/api/user"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", test.Main)

	e.POST("/register", user.Register)
	e.POST("/show/token", user.GetToken)

	e.Use()

	g := e.Group("")
	{
		g.Use(middleware.Auth())

		g.POST("/link/store", link.Store)
		g.GET("/link/list", link.List)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
