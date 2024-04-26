package restHandler

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func validJwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenHeader := c.Request().Header.Get("Authorization")
		if len(tokenHeader) < 7 || !strings.HasPrefix(tokenHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization Header")
		}

		bearerToken := strings.TrimPrefix(tokenHeader, "Bearer ")

		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusInternalServerError, "Unexpected signing method")
			}
			secret := []byte("JwtSecretKey")
			return secret, nil
		})
		if err != nil {
			return err
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId, ok := claims["user_id"]
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "InvalidJWTClaims")
			}
			c.Set("userId", userId)
		}

		return next(c)
	}
}
