package handlers

import (
	"github.com/CalumMackenzie-Chambers/templ-ui/app/views/layouts"
	"github.com/CalumMackenzie-Chambers/templ-ui/app/views/pages/docs"
	"github.com/labstack/echo/v4"
)

type DocsHandler struct {
}

func (h DocsHandler) GetDocs(c echo.Context) error {

	opts := renderOptions{
		Ctx:           c,
		Component:     docs.Docs(),
		BaseComponent: layouts.Base(indexMeta),
		PageTitle:     indexMeta.MetaData.Title,
	}
	return render(opts)
}
