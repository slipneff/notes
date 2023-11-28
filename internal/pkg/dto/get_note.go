package dto

type GetNote struct {
	NoteId string `json:"note_id" binding:"required"`
}
