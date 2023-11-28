package sql_test

import (
	"context"
	"testing"

	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/slipneff/notes/internal/pkg/storage/sql"
	"github.com/slipneff/notes/internal/test/fake"
	"github.com/slipneff/notes/internal/test/testsql"
	"github.com/stretchr/testify/require"
)

func TestStorage_CreateNote(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = sql.MustNewTestDB(t)
		storage = sql.New(db, trmgorm.DefaultCtxGetter)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		fakeNote := fake.Note()
		created, err := storage.CreateNote(ctx, fakeNote)
		require.NoError(t, err)

		found := testsql.MustFindNote(t, db, created.Id)
		created.CreatedAt = nil
		found.CreatedAt = nil
		require.Equal(t, created, found)
	})
	t.Run("error", func(t *testing.T) {
		t.Parallel()

		fakeNote := fake.Note()

		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, err := storage.CreateNote(ctx, fakeNote)
		require.Error(t, err)
	})
	t.Run("entity exists", func(t *testing.T) {
		t.Parallel()

		created1 := testsql.MustCreateNote(t, db, fake.Note())

		fakeNote := fake.Note()
		fakeNote.Id = created1.Id

		_, err := storage.CreateNote(ctx, fakeNote)
		require.ErrorIs(t, err, sql.ErrEntityExists)
	})
}
func TestStorage_FindNote(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = sql.MustNewTestDB(t)
		storage = sql.New(db, trmgorm.DefaultCtxGetter)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		fakeNote := fake.Note()
		created := testsql.MustCreateNote(t, db, fakeNote)

		found, err := storage.FindNoteById(ctx, created.Id.String())
		require.NoError(t, err)
		created.CreatedAt = nil
		found.CreatedAt = nil

		require.Equal(t, created, found)
	})
	t.Run("error", func(t *testing.T) {
		t.Parallel()

		fakeNote := fake.Note()

		addCard := testsql.MustCreateNote(t, db, fakeNote)

		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, err := storage.FindNoteById(ctx, addCard.Id.String())
		require.Error(t, err)
	})
	t.Run("entity not exists", func(t *testing.T) {
		t.Parallel()

		fakeNote := fake.Note()

		_, err := storage.FindNoteById(ctx, fakeNote.Id.String())
		require.ErrorIs(t, err, sql.ErrEntityNotFound)
	})
}

func TestStorage_UpdateNote(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = sql.MustNewTestDB(t)
		storage = sql.New(db, trmgorm.DefaultCtxGetter)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		created := testsql.MustCreateNote(t, db, fake.Note())

		fake := fake.Note()
		fake.Id = created.Id

		updated, err := storage.UpdateNote(ctx, fake)
		require.NoError(t, err)
		
		found := testsql.MustFindNote(t, db, updated.Id)
		fake.CreatedAt = nil
		found.CreatedAt = nil
		require.Equal(t, fake, *found)
	})
	t.Run("error", func(t *testing.T) {
		t.Parallel()

		fake := fake.Note()
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, err := storage.UpdateNote(ctx, fake)
		require.Error(t, err)
	})
	t.Run("entity not found", func(t *testing.T) {
		t.Parallel()
		fake := fake.Note()
		_, err := storage.UpdateNote(ctx, fake)
		require.ErrorIs(t, err, sql.ErrEntityNotFound)
	})
}

func TestStorage_DeleteNote(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = sql.MustNewTestDB(t)
		storage = sql.New(db, trmgorm.DefaultCtxGetter)
	)

	t.Run("success", func(t *testing.T) {
		fakeNote := fake.Note()
		created := testsql.MustCreateNote(t, db, fakeNote)

		deleted, err := storage.DeleteNote(ctx, created.Id.String())
		require.NoError(t, err)
		testsql.MustNotFindNote(t, db, deleted.Id)
	})
	t.Run("error", func(t *testing.T) {
		fakeNote := fake.Note()

		note := testsql.MustCreateNote(t, db, fakeNote)

		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, err := storage.DeleteNote(ctx, note.Id.String())
		require.Error(t, err)
	})
	t.Run("entity not exists", func(t *testing.T) {
		t.Parallel()

		fakeNote := fake.Note()

		_, err := storage.DeleteNote(ctx, fakeNote.Id.String())
		require.ErrorIs(t, err, sql.ErrEntityNotFound)
	})
}
