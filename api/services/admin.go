package services

import (
	"net/http"

	"github.com/jsec/blog-aggregator/internal/types/dto"
	"github.com/labstack/echo/v4"
)

type AdminService struct{}

func NewAdminService() AdminService {
	return AdminService{}
}

func (a *AdminService) RegisterRoutes(g *echo.Group) {
	g.GET("/health", a.healthCheck)
}

func (a *AdminService) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.HealthCheckResponse{
		Status: "ok",
	})
}
