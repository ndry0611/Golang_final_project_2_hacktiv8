package entity

import "time"

type Comment struct {
	ID        int `gorm:primaryKey;not null;type:int" json:"id"`
	UserID    int
	PhotoID   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Photo     Photo
}
