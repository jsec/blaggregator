package auth

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
)

func ParseAuthHeader(c echo.Context) (string, error) {
	header := c.Request().Header.Get("Authorization")

	if header == "" {
		return "", errors.New("No authorization header")
	}

	apiKey := strings.TrimPrefix(header, "ApiKey ")
	if apiKey == "" {
		return "", errors.New("ApiKey not found")
	}

	return apiKey, nil
}
