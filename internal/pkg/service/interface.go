package note

import (
	"context"

	"github.com/slipneff/notes/internal/pkg/models"
)

type noteStorage interface {
	CreateNote(ctx context.Context, m models.Note) (*models.Note, error)
	FindNoteById(ctx context.Context, id string) (*models.Note, error)
	GetAllNotes(ctx context.Context) ([]models.Note, error)
	UpdateNote(ctx context.Context, m models.Note) (*models.Note, error)
	DeleteNote(ctx context.Context, id string) (*models.Note, error)
}

type Service struct {
	storage noteStorage
}

func New(storage noteStorage) *Service {
	return &Service{storage: storage}
}
