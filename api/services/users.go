package services

import (
	"net/http"

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
