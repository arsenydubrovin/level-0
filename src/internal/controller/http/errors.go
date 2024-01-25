package http

import (
	"log/slog"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func (r *orderRouter) errorResponse(ctx echo.Context, status int, mgs any) error {
	return ctx.JSON(status, envelope{"error": mgs})
}

func (r *orderRouter) serverErrorResponse(ctx echo.Context, err error) error {
	slog.Error("internal server error", slog.Any("error", err))

	msg := "the server has encountered an error and cannot process the request"
	return r.errorResponse(ctx, http.StatusInternalServerError, msg)
}

func (r *orderRouter) notFoundResponse(ctx echo.Context) error {
	msg := "the requested resource was not found"
	return r.errorResponse(ctx, http.StatusNotFound, msg)
}
