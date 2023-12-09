package models

type Domain struct {
	BaseModel
	Domain string `gorm:"column:domain;not null" json:"domain"`
	UserID int
	User   User
}
