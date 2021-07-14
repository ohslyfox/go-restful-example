package db

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string    `gorm:"unique" json:"title"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"publish-date"`
	Rating      uint      `json:"rating"`
	Status      bool      `json:"checked-in:"`
}
