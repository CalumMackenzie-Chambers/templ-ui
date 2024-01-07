package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/CalumMackenzie-Chambers/templ-ui/docs/handlers"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.Static("/", "./docs/static")

	indexHandler := handlers.IndexHandler{}
	e.GET("/", indexHandler.GetHome)

	e.Logger.Fatal(e.Start(":1323"))
}