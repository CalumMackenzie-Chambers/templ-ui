package handlers

import (
	"github.com/CalumMackenzie-Chambers/templ-ui/app/views/layouts"
	"github.com/CalumMackenzie-Chambers/templ-ui/app/views/pages/examples"
	"github.com/labstack/echo/v4"
)

type ExamplesHandler struct {
}

func (h ExamplesHandler) GetExamples(c echo.Context) error {

	opts := renderOptions{
		Ctx:           c,
		Component:     examples.Examples(),
		BaseComponent: layouts.Base(indexMeta),
		PageTitle:     indexMeta.MetaData.Title,
	}
	return render(opts)
}
