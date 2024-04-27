package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func readiness(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthCheckResponse{
		Status: "ok",
	})
}

func errorHandle(_ echo.Context) error {
	return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	port := os.Getenv("PORT")
	e := echo.New()
	e.GET("/v1/readiness", readiness)
	e.GET("/v1/err", errorHandle)

	e.Logger.Fatal(e.Start(port))
}
