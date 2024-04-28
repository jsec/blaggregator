package config

import (
	"github.com/jsec/blog-aggregator/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func NewConfig() ApiConfig {
	return ApiConfig{
		DB: database.New(database.Connect()),
	}
}
