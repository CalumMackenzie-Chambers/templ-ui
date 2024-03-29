package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/CalumMackenzie-Chambers/templ-ui/app/handlers"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.Static("/", "./app/static")

	indexHandler := handlers.IndexHandler{}
	e.GET("/", indexHandler.GetHome)

	componentHandler := handlers.ComponentHandler{}
	e.GET("/components", componentHandler.GetIndex)
	e.GET("/components/button", componentHandler.GetButton)
	e.GET("/components/accordion", componentHandler.GetAccordion)
	e.GET("/components/alert", componentHandler.GetAlert)
	e.GET("/components/avatar", componentHandler.GetAvatar)
	e.GET("/components/separator", componentHandler.GetSeparator)
	e.GET("/components/badge", componentHandler.GetBadge)

	docHandler := handlers.DocsHandler{}
	e.GET("/docs", docHandler.GetDocs)

	examplesHandler := handlers.ExamplesHandler{}
	e.GET("/examples", examplesHandler.GetExamples)

	e.Logger.Fatal(e.Start(":1323"))
}
