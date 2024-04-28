package dto

type CreateUserRequest struct {
	Name string `json:"name"`
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}
