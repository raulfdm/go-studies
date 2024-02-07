package data

import "time"

type BlogPost struct {
	Title         string
	Author        string
	Content       string
	DatePublished time.Time
}

func GetPosts() []BlogPost {
	return listOfPosts
}

// hardcoded posts
var listOfPosts = []BlogPost{
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
