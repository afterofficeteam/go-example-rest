package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/afteroffice/go-example-rest/internal/usecase"
	"github.com/labstack/echo/v4"
)

type handler struct {
	uc usecase.BooksUCInterface
}

func NewHandler(uc usecase.BooksUCInterface) handler {
	return handler{
		uc: uc,
	}
}

func (h *handler) GetBooks(c echo.Context) error {
	// parse request
	limit, err := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	if err != nil {
		slog.Error("GetBooks", "err", err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	// call usecase
	books, err := h.uc.GetBooks(&usecase.GetBookRequest{
		Limit: int(limit),
	})
	if err != nil {
		slog.Error("GetBooks", "err", err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	// return response
	return c.JSON(http.StatusOK, books)
}
