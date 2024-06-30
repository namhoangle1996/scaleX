package restHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"scaleX/internal/usecase"
)

type (
	OpenApiHandler interface {
		SummarizeChapters(c echo.Context) error
	}

	openApiHandler struct {
		usecase.OpenApiService
	}
)

func (h *openApiHandler) SummarizeChapters(c echo.Context) error {
	err := h.OpenApiService.SummarizeChapters(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)

}
