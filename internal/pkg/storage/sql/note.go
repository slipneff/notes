package sql

import (
	"context"
	"errors"

	"github.com/slipneff/notes/internal/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Storage) CreateNote(ctx context.Context, m models.Note) (*models.Note, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Create(&m).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return nil, ErrEntityExists
		}

		return nil, err
	}

	return &m, nil
}

func (s *Storage) FindNoteById(ctx context.Context, id string) (*models.Note, error) {
	var note models.Note

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.First(&note, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrEntityNotFound
		}

		return nil, err
	}

	return &note, nil
}

func (s *Storage) GetAllNotes(ctx context.Context) ([]models.Note, error) {
	var notes []models.Note
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Find(&notes).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrEntityNotFound
		}
		return nil, err
	}
	return notes, nil
}

func (s *Storage) UpdateNote(ctx context.Context, m models.Note) (*models.Note, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	result := tr.Model(&m).Where("id = ?", m.Id).Updates(m)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrEntityNotFound
		}

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, ErrEntityExists
		}

		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrEntityNotFound
	}

	return &m, nil
}

func (s *Storage) DeleteNote(ctx context.Context, id string) (*models.Note, error) {
	var note models.Note

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	result := tr.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&note)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrEntityNotFound
		}

		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrEntityNotFound
	}

	return &note, nil
}
