package db

import (
	"time"
)

type Book struct {
	ID          uint64    `gorm:"primarykey" json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"publish-date"`
	Rating      uint      `json:"rating"`
	Status      bool      `gorm:"default:true" json:"checked-in:"`
}
