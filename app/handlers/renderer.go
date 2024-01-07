package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type renderOptions struct {
	Ctx           echo.Context
	Component     templ.Component
	BaseComponent templ.Component
	PageTitle     string
}

func render(options renderOptions) error {
	ctx := options.Ctx.Request().Context()
	layout := options.BaseComponent
	return layout.Render(templ.WithChildren(ctx, options.Component), options.Ctx.Response())
}
