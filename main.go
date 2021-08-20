package main

import (
	"bitly-clone/internal/repository"
	"bitly-clone/internal/router"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	router.Init(e)

	repository.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
