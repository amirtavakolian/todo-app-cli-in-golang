package dto

type Task struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description bool   `json:"description"`
	Category_id string `json:"category_id"`
}
