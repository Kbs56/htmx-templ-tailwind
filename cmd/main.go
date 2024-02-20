package main

import (
	"github.com/labstack/echo/v4"

	"github.com/Kbs56/htmx-templ-tailwind/handler"
)

func main() {
	e := echo.New()

	userHandler := handler.UserHandler{}

	e.GET("/", userHandler.HandleUserShow)

	e.Logger.Fatal(e.Start(":3000"))
}
