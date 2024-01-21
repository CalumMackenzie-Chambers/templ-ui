package handlers

import (
	"github.com/CalumMackenzie-Chambers/templ-ui/app/views/layouts"
	"github.com/CalumMackenzie-Chambers/templ-ui/app/views/pages/components"
	"github.com/labstack/echo/v4"
)

type ComponentHandler struct {
}

func (h ComponentHandler) GetIndex(c echo.Context) error {

	opts := renderOptions{
		Ctx:           c,
		Component:     components.Index(),
		BaseComponent: layouts.Shell(indexMeta),
		PageTitle:     indexMeta.MetaData.Title,
	}
	return render(opts)
}

func (h ComponentHandler) GetButton(c echo.Context) error {

	opts := renderOptions{
		Ctx:           c,
		Component:     components.Button(),
		BaseComponent: layouts.Shell(buttonMeta),
		PageTitle:     buttonMeta.MetaData.Title,
	}
	return render(opts)
}

func (h ComponentHandler) GetAccordion(c echo.Context) error {
	opts := renderOptions{
		Ctx:           c,
		Component:     components.Accordion(),
		BaseComponent: layouts.Shell(accordionMeta),
		PageTitle:     accordionMeta.MetaData.Title,
	}
	return render(opts)
}

func (h ComponentHandler) GetAlert(c echo.Context) error {
	opts := renderOptions{
		Ctx:           c,
		Component:     components.Alert(),
		BaseComponent: layouts.Shell(alertMeta),
		PageTitle:     alertMeta.MetaData.Title,
	}
	return render(opts)
}

func (h ComponentHandler) GetAvatar(c echo.Context) error {
	opts := renderOptions{
		Ctx:           c,
		Component:     components.Avatar(),
		BaseComponent: layouts.Shell(avatarMeta),
		PageTitle:     avatarMeta.MetaData.Title,
	}
	return render(opts)
}

func (h ComponentHandler) GetSeparator(c echo.Context) error {
	opts := renderOptions{
		Ctx:           c,
		Component:     components.Separator(),
		BaseComponent: layouts.Shell(separatorMeta),
		PageTitle:     separatorMeta.MetaData.Title,
	}
	return render(opts)
}

func (h ComponentHandler) GetBadge(c echo.Context) error {
	opts := renderOptions{
		Ctx: c,
		Component: components.Badge(),
		BaseComponent: layouts.Shell(badgeMeta),
		PageTitle: badgeMeta.MetaData.Title,
	}
	return render(opts)
}
