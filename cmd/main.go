package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Kbs56/htmx-templ-tailwind/handler"
)

func main() {
	e := echo.New()
	e.Static("/static", "css")

	userHandler := handler.UserHandler{}
	indexHandler := handler.IndexHandler{}

	e.GET("/", indexHandler.HandleIndexShow)
	e.GET("/user", userHandler.HandleUserShow)

	e.POST("/clicked", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
