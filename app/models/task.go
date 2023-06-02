package models

import "time"

type Task struct {
	ID              int    `json:"id" gorm:"primary_key"`
	Task            string `json:"task"`
	TaskDescription string `json:"task_description"`
	UserID          int    `json:"user_id" gorm:"primaryKey"`
	User            User
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
