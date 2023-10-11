package entity

import "time"

type Photo struct {
	ID        int    `gorm:primaryKey;not null;type:int" json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserID    int    `json:"user_id"`
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}
