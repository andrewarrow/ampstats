package parse

import (
	"encoding/json"
	"time"
)

type Result struct {
	NextPageToken string   `json:"nextPageToken"`
	Items         []Item   `json:"items"`
	PageInfo      PageInfo `json:"pageInfo"`
}

type PageInfo struct {
	TotalResults int `json:"totalResults"`
}

type Item struct {
	Id      Id      `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type Id struct {
	VideoId string `json:"videoId"`
}

type Snippet struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:publishedAt"`
}

func ParseJson(jsonString string) *Result {
	var result Result
	json.Unmarshal([]byte(jsonString), &result)
	return &result
}
