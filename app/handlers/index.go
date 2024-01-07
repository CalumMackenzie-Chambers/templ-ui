package handlers

import (
	"github.com/CalumMackenzie-Chambers/templ-ui/docs/views/layouts"
	"github.com/CalumMackenzie-Chambers/templ-ui/docs/views/pages"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct {
}

func (h IndexHandler) GetHome(c echo.Context) error {

	opts := renderOptions{
		Ctx:           c,
		Component:     pages.Home(),
		BaseComponent: layouts.Base(indexMeta),
		PageTitle:     indexMeta.MetaData.Title,
	}
	return render(opts)
}
