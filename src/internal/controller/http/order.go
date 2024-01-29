package http

import (
	"errors"
	"net/http"

	"github.com/arsenydubrovin/level-0/src/internal/model"
	echo "github.com/labstack/echo/v4"
)

func (r *orderRouter) fetchOrderHandler(ctx echo.Context) error {
	uid := ctx.QueryParam("uid")

	if len(uid) != model.OrderUIDLength {
		return r.notFoundResponse(ctx)
	}

	order, err := r.service.Fetch(ctx.Request().Context(), uid)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrOrderNotFound):
			return r.notFoundResponse(ctx)
		default:
			return r.serverErrorResponse(ctx, err)
		}
	}

	return ctx.JSONPretty(http.StatusOK, envelope{"order": order}, "  ")
}

func (r *orderRouter) listUIDsHandler(ctx echo.Context) error {
	uids, err := r.service.ListUIDs(ctx.Request().Context())
	if err != nil {
		return r.serverErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, envelope{"uids": uids})
}
