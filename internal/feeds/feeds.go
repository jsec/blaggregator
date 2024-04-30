package feeds

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jsec/blog-aggregator/internal/database"
	"github.com/jsec/blog-aggregator/internal/types/dto"
)

type FeedRepository struct {
	db *database.Queries
}

func NewFeedRepository(db *database.Queries) FeedRepository {
	return FeedRepository{
		db: db,
	}
}

func (r *FeedRepository) CreateFeed(req dto.CreateFeedRequest, userId uuid.UUID) (database.Feed, error) {
	feed, err := r.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      req.Name,
		Url:       req.Url,
		UserID:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		fmt.Println(err.Error())
		return database.Feed{}, err
	}

	return feed, nil
}
