package database

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	OriginalURL string `gorm:"uniqueIndex"`
	ShortURL    string `gorm:"uniqueIndex"`
}
