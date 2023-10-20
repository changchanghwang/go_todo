package dto

type TodoCreateDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoUpdateDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
