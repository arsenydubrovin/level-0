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
	// uids, err := r.s.ListUIDs(ctx.Request().Context())
	// if err != nil {
	// 	return r.serverErrorResponse(ctx, err)
	// }

	// return ctx.Render(http.StatusOK, "index.html", map[string]interface{}{"uids": uids})
	return ctx.Render(http.StatusOK, "index.html", nil)
}
