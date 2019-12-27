package models

// A Task represents the tasks which a user can fulfill.
type Task struct {
	BaseModel
	Name    string      `json:"name" binding:"required"`
	Entries []TaskEntry `json:"task_entries" binding:"-"`
}
