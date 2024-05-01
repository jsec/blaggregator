package feeds

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jsec/blog-aggregator/internal/database"
)

type FollowRepository struct {
	db *database.Queries
}

func NewFollowRepository(db *database.Queries) FollowRepository {
	return FollowRepository{
		db: db,
	}
}

func (r *FollowRepository) CreateFollow(feedId uuid.UUID, userId uuid.UUID) (database.FeedFollow, error) {
	follow, err := r.db.CreateFollow(context.Background(), database.CreateFollowParams{
		ID:        uuid.New(),
		UserID:    userId,
		FeedID:    feedId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return database.FeedFollow{}, err
	}

	return follow, nil
}
