package api

import (
	"context"
	"net/http"

	"github.com/jsec/blog-aggregator/internal/auth"
	"github.com/labstack/echo/v4"
)

var Unauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")

func (s *Server) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		apiKey, err := auth.ParseAuthHeader(ctx)
		if err != nil {
			return Unauthorized
		}

		user, err := s.config.DB.GetUserByApiKey(context.Background(), apiKey)
		if err != nil {
			return Unauthorized
		}

		ctx.Set("UserID", user.ID)

		return next(ctx)
	}
}
