package models

// Category is the core datatype this app is built around.
type Category struct {
	BaseModel
	Name string `json:"name" binding:"required"`
}
