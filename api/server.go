package api

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	api    *echo.Echo
	config ApiConfig
}

func NewServer() Server {
	return Server{
		api:    echo.New(),
		config: NewConfig(),
	}
}

func (s *Server) RegisterMiddleware() {
	s.api.Pre(middleware.RemoveTrailingSlash())
}

func (s *Server) RegisterServices() {
	adminService := NewAdminService()
	s.api.GET("/admin/health", adminService.HealthCheck)

	v1 := s.api.Group("/v1")

	userService := NewUserService(s.config.DB)
	v1.GET("/users", userService.GetUser, s.AuthMiddleware)
	v1.POST("/users", userService.CreateUser)

	feedService := NewFeedService(s.config.DB)
	v1.POST("/feeds", feedService.CreateFeed, s.AuthMiddleware)
}

func (s *Server) Start() {
	port := os.Getenv("PORT")
	api := s.api
	api.Logger.Fatal(api.Start(port))
}
