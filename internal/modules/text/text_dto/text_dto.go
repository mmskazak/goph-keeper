package text_dto

type TextDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TextContent string `json:"text_content"`
}
