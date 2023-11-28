package server

import (
	"github.com/gin-gonic/gin"
	note "github.com/slipneff/notes/internal/pkg/service"
	"github.com/slipneff/notes/internal/utils/config"
)

type Handler struct {
	cfg         *config.Config
	noteService *note.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	note := router.Group("/note")
	{
		note.POST("/", h.CreateNote)
		note.GET("/:id", h.GetNote)
		note.PATCH("/:id", h.UpdateNote)
		note.DELETE("/:id", h.DeleteNote)
	}
	router.GET("/notes", h.GetNotes)

	return router
}

func New(cfg *config.Config, noteService *note.Service) *Handler {

	return &Handler{cfg: cfg, noteService: noteService}

}
