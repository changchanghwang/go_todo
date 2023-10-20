package dto

type TodoCreateDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
