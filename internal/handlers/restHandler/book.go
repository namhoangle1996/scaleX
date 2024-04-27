package restHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"scaleX/internal/dto"
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

	var req dto.AddBookRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.BookService.AddBook(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, req)

}

func (h *bookHandler) DeleteBook(c echo.Context) error {

	var req dto.DeleteBookRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.BookService.DeleteBook(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, req)
}

func (h *bookHandler) FetchBook(c echo.Context) error {
	userId := c.Get("userId").(string)
	r, err := h.BookService.FetchBook(c.Request().Context(), userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, r)
}
