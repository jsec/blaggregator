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

type FeedService struct {
	repo feeds.FeedRepository
}

func NewFeedService(db *database.Queries) FeedService {
	return FeedService{
		repo: feeds.NewFeedRepository(db),
	}
}

func (s *FeedService) CreateFeed(c echo.Context) error {
	var req dto.CreateFeedRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	userIdContext := c.Get("UserID")
	userID, ok := userIdContext.(uuid.UUID)
	if !ok {
		return errors.New("Unable to get UserID from context")
	}

	feed, err := s.repo.CreateFeed(req, userID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, feed)
}

func (s *FeedService) GetFeeds(c echo.Context) error {
	feeds, err := s.repo.GetFeeds()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, feeds)
}
