package restHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"scaleX/internal/dto"
	"scaleX/internal/usecase"
)

type (
	AuthHandler interface {
		Login(c echo.Context) error
	}

	authHandler struct {
		usecase.AuthService
	}
)

func (h *authHandler) Login(c echo.Context) error {
	var req dto.LoginReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	r, err := h.AuthService.Login(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, r)
}
