package models

type UpdateSalaryRequest struct {
	ID     int64   `json:"id" binding:"required"`
	Salary float64 `json:"salary" binding:"required"`
}
