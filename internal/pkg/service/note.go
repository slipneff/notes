package note

import (
	"context"

	"github.com/slipneff/notes/internal/pkg/models"
)

func (s *Service) CreateNote(ctx context.Context, m models.Note) (*models.Note, error) {
	note, err := s.storage.CreateNote(ctx, m)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *Service) FindANoteById(ctx context.Context, noteId string) (*models.Note, error) {
	note, err := s.storage.FindNoteById(ctx, noteId)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (s *Service) GetAllNotes(ctx context.Context) ([]models.Note, error) {
	notes, err := s.storage.GetAllNotes(ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *Service) UpdateNote(ctx context.Context, m models.Note) (*models.Note, error) {
	note, err := s.storage.UpdateNote(ctx, m)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (s *Service) DeleteNote(ctx context.Context, noteId string) (*models.Note, error) {
	note, err := s.storage.DeleteNote(ctx, noteId)
	if err != nil {
		return nil, err
	}

	return note, nil
}