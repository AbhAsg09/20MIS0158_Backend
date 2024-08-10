package models

import "time"

type Video struct {
	Title        string
	Description  string
	PublishedAt  time.Time
	ThumbnailURL string // Change this to string
}
