package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listOfPosts)

	w.WriteHeader(http.StatusOK)
}

type blogPost struct {
	Title         string
	Author        string
	Content       string
	DatePublished time.Time
}

// hardcoded posts
var listOfPosts = []blogPost{
	{
		Title:         "First Post",
		Author:        "John Doe",
		Content:       "This is the first post on this blog. I hope you enjoy it!",
		DatePublished: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		Title:         "Second Post",
		Author:        "Jane Doe",
		Content:       "This is the second post on this blog. I hope you enjoy it!",
		DatePublished: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
	},
	{
		Title:         "Third Post",
		Author:        "John Doe",
		Content:       "This is the third post on this blog. I hope you enjoy it!",
		DatePublished: time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
	},
}
