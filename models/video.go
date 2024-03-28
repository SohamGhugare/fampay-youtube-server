package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	ID           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnail_url"`
	PublishedAt  string `json:"published_at"`
}
