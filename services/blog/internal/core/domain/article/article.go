package article

import "time"

type Article struct {
	ID          int64
	Title       string `gorm:"uniqueIndex"`
	Description string
	Content     string
	CoverImage  string
	PublishedAt time.Time
}
