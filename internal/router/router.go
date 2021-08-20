package router

import (
	"bitly-clone/internal/api/link"
	"bitly-clone/internal/api/middleware"
	"bitly-clone/internal/api/user"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	r := e.Group("/api")
	r.GET("/:link", link.Redirect)
	r.POST("/register", user.Register)
	r.POST("/show/token", user.GetToken)

	r.Use()

	g := r.Group("")
	{
		g.Use(middleware.Auth())

		g.POST("/link/store", link.Store)
		g.DELETE("/link/delete", link.Delete)
		g.GET("/link/list", link.List)
	}

}
