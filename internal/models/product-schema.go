package models

import (
	"time"
)


type Product struct {
	ID          string    `json:"id" gorm:"primaryKey;type:varchar(26)"` 
	Slug        string    `json:"slug" gorm:"uniqueIndex;not null"`
	Name        string    `json:"name" gorm:"not null"`
	ImageURL    string    `json:"imageUrl" gorm:"not null"`
	Price       int       `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}