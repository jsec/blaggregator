package dto

type CreateFeedRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CreateFollowRequest struct {
	FeedID string `json:"feed_id"`
}
