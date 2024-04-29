package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func ParseAuthHeader(c echo.Context) (string, error) {
	header := c.Request().Header.Get("Authorization")

	if header == "" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	apiKey := strings.TrimPrefix(header, "ApiKey ")
	if apiKey == "" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	return apiKey, nil
}
