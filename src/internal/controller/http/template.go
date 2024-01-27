package http

import (
	"io"
	"net/http"
	"text/template"

	echo "github.com/labstack/echo/v4"
)

type renderer struct {
	t *template.Template
}

func (r *renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.t.ExecuteTemplate(w, name, data)
}

func (r *orderRouter) mainPageHandler(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index.html", nil)
}
