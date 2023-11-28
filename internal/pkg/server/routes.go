package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slipneff/notes/internal/pkg/dto"
	"github.com/slipneff/notes/internal/pkg/models"
)

func (h *Handler) CreateNote(c *gin.Context) {
	var req dto.CreateNote
	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Error when parse JSON")
		return
	}
	note, err := h.noteService.CreateNote(c, models.Note{
		Header:  req.Header,
		Content: req.Content,
	})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Internal Error")
		return
	}
	c.JSON(http.StatusOK, note)
}
func (h *Handler) GetNote(c *gin.Context) {
	id := c.Param("id")
	note, err := h.noteService.FindANoteById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Internal Error")
		return
	}
	c.JSON(http.StatusOK, note)
}
func (h *Handler) GetNotes(c *gin.Context) {
	notes, err := h.noteService.GetAllNotes(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Internal Error")
		return
	}
	c.JSON(http.StatusOK, notes)
}
func (h *Handler) UpdateNote(c *gin.Context) {
	var req dto.CreateNote
	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Error when parse JSON")
		return
	}
	id := c.Param("id")
	note, err := h.noteService.UpdateNote(c, models.Note{
		Id:      uuid.MustParse(id),
		Header:  req.Header,
		Content: req.Content,
	})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Internal Error")
		return
	}
	c.JSON(http.StatusOK, note)
}
func (h *Handler) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	note, err := h.noteService.DeleteNote(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Internal Error")
		return
	}
	c.JSON(http.StatusOK, note)
}
