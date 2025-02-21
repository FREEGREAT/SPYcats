package models

import "time"

type TargetModel struct {
	ID          int64     `json:"id"`
	MissionID   int64     `json:"mission_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Country     string    `json:"country" binding:"required"`
	IsCompleted bool      `json:"is_completed" `
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
