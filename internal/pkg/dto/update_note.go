package dto

type UpdateNote struct {
	Id      string `json:"id" binding:"required"`
	Header  string `json:"header" binding:"required"`
	Content string `json:"content" binding:"required"`
}
