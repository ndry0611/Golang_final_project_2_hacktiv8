package entity

import "time"

type User struct {
	ID        int    `gorm:"primaryKey;not null;type:int" json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
