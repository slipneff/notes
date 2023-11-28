package dto

type CreateNote struct {
	Header  string `json:"header" binding:"required"`
	Content string `json:"content" binding:"required"`
}
