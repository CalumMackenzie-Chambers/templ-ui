package handlers

import (
	"github.com/CalumMackenzie-Chambers/templ-ui/docs/views/layouts"
)

var indexMeta = layouts.PageData{
	MetaData: layouts.MetaData{
		Title:       "templ-ui - UI componennts for golang templ and tailwind",
		Description: "Build websites faster using set of standerd, minimally styled components for templ",
	},
	HeadComponent: layouts.EmptyHead,
}
