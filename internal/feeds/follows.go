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
	return r.db.CreateFollow(context.Background(), database.CreateFollowParams{
		ID:        uuid.New(),
		UserID:    userId,
		FeedID:    feedId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func (r *FollowRepository) DeleteFollow(followId, userId uuid.UUID) error {
	return r.db.DeleteFollow(context.Background(), database.DeleteFollowParams{
		ID:     followId,
		UserID: userId,
	})
}
