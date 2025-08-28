package models

import (
	"time"
	"gorm.io/gorm"
)

type Content struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Title     string         `json:"title" gorm:"type:varchar(255);not null"`
	Text      string         `json:"text" gorm:"type:mediumtext;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ContentRequest struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
	Text  string `json:"text" binding:"required,min=1"`
}

type ContentResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Excerpt   string    `json:"excerpt"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
