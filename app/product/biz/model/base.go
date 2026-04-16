package model

import "time"

type Base struct {
	ID 				int `gorm:"primaykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}