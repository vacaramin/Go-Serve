package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	URL          string
	ShortenedURL string
}
