package entity

import "time"

type SocialMedia struct {
	ID             int `gorm:primaryKey;not null;type:int" json:"id"`
	Name           string
	SocialMediaUrl string
	UserID         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           User
}
