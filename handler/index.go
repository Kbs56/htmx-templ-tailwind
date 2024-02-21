package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/Kbs56/htmx-templ-tailwind/view/index"
)

type IndexHandler struct{}

func (h IndexHandler) HandleIndexShow(c echo.Context) error {
	return render(c, index.Show())
}
