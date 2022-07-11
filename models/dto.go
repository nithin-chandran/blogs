package models

// BlogRequest create request data for blog
type BlogRequest struct {
	ID          int    `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}
