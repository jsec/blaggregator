package services

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jsec/blog-aggregator/internal/database"
	dto "github.com/jsec/blog-aggregator/internal/types"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	db *database.Queries
}

func NewUserService(db *database.Queries) *UserService {
	return &UserService{
		db: db,
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

	ctx := context.Background()
	user, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
