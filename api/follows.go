package api

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jsec/blog-aggregator/internal/database"
	"github.com/jsec/blog-aggregator/internal/feeds"
	"github.com/jsec/blog-aggregator/internal/types/dto"
	"github.com/labstack/echo/v4"
)

type FollowService struct {
	repo feeds.FollowRepository
}

func NewFollowService(db *database.Queries) FollowService {
	return FollowService{
		repo: feeds.NewFollowRepository(db),
	}
}

func (s *FollowService) CreateFollow(c echo.Context) error {
	var req dto.CreateFollowRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	userIdContext := c.Get("UserID")
	userID, ok := userIdContext.(uuid.UUID)
	if !ok {
		return errors.New("Unable to get UserID from context")
	}

	feedId, err := uuid.Parse(req.FeedID)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	follow, err := s.repo.CreateFollow(feedId, userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, follow)
}

func (s *FollowService) DeleteFollow(c echo.Context) error {
	userIdContext := c.Get("UserID")
	userId, ok := userIdContext.(uuid.UUID)
	if !ok {
		return errors.New("Unable to get UserID from context")
	}

	followId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if err = s.repo.DeleteFollow(followId, userId); err != nil {
		return err
	}

	return c.String(http.StatusNoContent, "OK")
}

func (s *FollowService) GetFollowsByUserID(c echo.Context) error {
	userIdContext := c.Get("UserID")
	userId, ok := userIdContext.(uuid.UUID)
	if !ok {
		return errors.New("Unable to get UserID from context")
	}

	follows, err := s.repo.GetFollowsByUserID(userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, follows)
}
