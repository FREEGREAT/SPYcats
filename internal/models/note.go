package models

import "time"

type NoteModel struct {
	ID        int       `json:"id"`
	TargetID  int       `json:"target_id"`
	CatID     int       `json:"cat_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
