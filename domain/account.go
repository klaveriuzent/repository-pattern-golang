package domain

import "time"

type Account struct {
	Id          string `json:"id"`
	UserId      string `json:"user_id"`
	CreatedAt   time.Time
	IsActive    bool `gorm:"type:bool;default:true" json:"status"`
	LastLoginAt *time.Time
}
