package api

import (
	"net/http"

	"github.com/jsec/blog-aggregator/internal/types/dto"
	"github.com/labstack/echo/v4"
)

type AdminService struct{}

func NewAdminService() AdminService {
	return AdminService{}
}

func (a *AdminService) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.HealthCheckResponse{
		Status: "ok",
	})
}
