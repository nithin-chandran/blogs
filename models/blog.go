package models

// BlogModel reprsents the blog.
type BlogModel struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
