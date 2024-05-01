package entities

type Todo struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
