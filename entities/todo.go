package entities

// Todo entity
type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at"`
}
