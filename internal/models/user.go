package models

import "time"


type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:varchar(26)"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	
	
	FullName  *string   `json:"fullName"` 
	

	Password  *Password `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` 
	
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}


type Password struct {
	ID        string    `json:"id" gorm:"primaryKey;type:varchar(26)"`
	Hash      string    `json:"-" gorm:"not null"` 
	
	
	UserID    string    `json:"userId" gorm:"uniqueIndex;not null;type:varchar(26)"`
	
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}