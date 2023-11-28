package testsql

import (
	"testing"

	"github.com/google/uuid"
	"github.com/slipneff/notes/internal/pkg/models"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func MustFindNote(t *testing.T, db *gorm.DB, id uuid.UUID) *models.Note {
	addCard := new(models.Note)
	err := db.First(addCard, "id = ?", id).Error
	require.NoError(t, err)

	return addCard
}

func MustCreateNote(t *testing.T, db *gorm.DB, m models.Note) *models.Note {
	err := db.Create(&m).Error
	require.NoError(t, err)

	return &m
}

func MustNotFindNote(t *testing.T, db *gorm.DB, id uuid.UUID) {
	addCard := new(models.Note)
	err := db.First(addCard, "id = ?", id).Error
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
