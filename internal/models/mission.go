package models

import "time"

type MissionModel struct {
	ID          int       `json:"id"`
	CatID       int       `json:"cat_id"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
