package models

import "time"

type CatModel struct {
	ID              *int64     `json:"id"`
	Name            string     `json:"name" binding:"required"`
	ExperienceYears int        `json:"experience_years" binding:"required"`
	Breed           string     `json:"breed" binding:"required"`
	Salary          float64    `json:"salary" binding:"required"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type CatBreed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
