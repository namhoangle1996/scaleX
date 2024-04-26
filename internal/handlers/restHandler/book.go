package restHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"scaleX/internal/usecase"
)

type (
	BookHandler interface {
		AddBook(c echo.Context) error
		DeleteBook(c echo.Context) error
		FetchBook(c echo.Context) error
	}

	bookHandler struct {
		usecase.BookService
	}
)

func (h *bookHandler) AddBook(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *bookHandler) DeleteBook(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *bookHandler) FetchBook(c echo.Context) error {
	userId := c.Get("userId").(string)
	r, err := h.BookService.FetchBook(c.Request().Context(), userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, r)
}
