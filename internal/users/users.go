package users

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jsec/blog-aggregator/internal/database"
	"github.com/jsec/blog-aggregator/internal/types/dto"
)

type UserRepository struct {
	db *database.Queries
}

func NewUserRepository(db *database.Queries) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(req dto.CreateUserRequest) (database.User, error) {
	ctx := context.Background()
	user, err := r.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return database.User{}, err
	}

	return user, nil
}
