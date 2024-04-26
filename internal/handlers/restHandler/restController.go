package restHandler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"scaleX/internal/usecase"
	"strings"
)

type RestHandler struct {
	AuthHandler
	BookHandler
	//config
}

func Echo(h *RestHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	g := e.Group("/api/v1")

	// Author
	g.POST("/login", h.AuthHandler.Login)

	g.GET("/home", h.BookHandler.FetchBook, validJwtMiddleware)
	g.POST("/addBook", h.BookHandler.AddBook, validJwtMiddleware)
	g.DELETE("/deleteBook", h.BookHandler.DeleteBook, validJwtMiddleware)

	return e

}

func NewRestHandler(authService usecase.AuthService, bookService usecase.BookService) *RestHandler {
	return &RestHandler{
		AuthHandler: &authHandler{authService},
		BookHandler: &bookHandler{bookService},
	}
}