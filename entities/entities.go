package entities

type Todo struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
