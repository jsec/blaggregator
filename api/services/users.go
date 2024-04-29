package services

import (
	"net/http"

	"github.com/jsec/blog-aggregator/internal/auth"
	"github.com/jsec/blog-aggregator/internal/database"
	"github.com/jsec/blog-aggregator/internal/types/dto"
	"github.com/jsec/blog-aggregator/internal/users"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	repo users.UserRepository
}

func NewUserService(db *database.Queries) UserService {
	return UserService{
		repo: users.NewUserRepository(db),
	}
}

func (s *UserService) RegisterRoutes(g *echo.Group) {
	g.GET("/", s.getUser)
	g.POST("/", s.createUser)
}

func (s *UserService) createUser(c echo.Context) error {
	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	user, err := s.repo.CreateUser(req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (s *UserService) getUser(c echo.Context) error {
	apiKey, err := auth.ParseAuthHeader(c)
	if err != nil {
		return err
	}

	user, err := s.repo.GetUserByApiKey(apiKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, user)
}
