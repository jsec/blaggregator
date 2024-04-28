package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/jsec/blog-aggregator/api"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	s := api.NewServer()
	s.RegisterServices()
	s.Start()
}
