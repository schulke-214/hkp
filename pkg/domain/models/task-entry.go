package models

import "github.com/google/uuid"

// A TaskEntry represents a done task
type TaskEntry struct {
	BaseModel
	TaskID uuid.UUID `json:"task_id" binding:"omitempty,uuid"`
	User   string    `json:"user" binding:"required"`
}
