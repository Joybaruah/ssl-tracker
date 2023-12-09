package models

import (
	"time"
)

type BaseModel struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type User struct {
	BaseModel
	Name     string `gorm:"column:name;not null" json:"name"`
	Email    string `gorm:"column:email;not null;unique" json:"email"`
	Password string
}
