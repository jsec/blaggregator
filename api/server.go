package api

import (
	"os"

	"github.com/jsec/blog-aggregator/api/services"
	config "github.com/jsec/blog-aggregator/internal"
	"github.com/labstack/echo/v4"
)

type Server struct {
	api    *echo.Echo
	config config.ApiConfig
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func NewServer() Server {
	return Server{
		api:    echo.New(),
		config: config.NewConfig(),
	}
}

func (s *Server) RegisterServices() {
	adminService := services.NewUserService(s.config.DB)
	adminService.RegisterRoutes(s.api.Group("/admin"))

	v1 := s.api.Group("/v1")

	userService := services.NewUserService(s.config.DB)
	userService.RegisterRoutes(v1.Group("/users"))
}

func (s *Server) Start() {
	port := os.Getenv("")
	api := s.api
	api.Logger.Fatal(api.Start(port))
}
