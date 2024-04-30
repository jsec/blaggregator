package dto

type CreateFeedRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
