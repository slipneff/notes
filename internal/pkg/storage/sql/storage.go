package sql

import (
	"fmt"
	"log"
	"os"
	"testing"

	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/slipneff/notes/internal/pkg/models"
	"github.com/slipneff/notes/internal/utils/config"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	db     *gorm.DB
	getter *trmgorm.CtxGetter
}

func New(db *gorm.DB, getter *trmgorm.CtxGetter) *Storage {
	return &Storage{
		db:     db,
		getter: getter,
	}
}

func buildDSN(cfg *config.Config) string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Name, cfg.DB.Sslmode)

	return dsn
}

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(buildDSN(cfg)), &gorm.Config{
		TranslateError: true,
	})
}

func MustNewPostgresDB(cfg *config.Config) *gorm.DB {
	db, err := NewPostgresDB(cfg)
	if err != nil {
		log.Panic(err, "MustNewPostgresDB")
	}

	err = db.AutoMigrate(&models.Note{})

	return db
}

func NewSQLiteDB(cfg *config.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(buildDSN(cfg)), &gorm.Config{
		TranslateError: true,
	})
}

func MustNewSQLite(cfg *config.Config) *gorm.DB {
	db, err := NewSQLiteDB(cfg)
	if err != nil {
		log.Panic(err, "MustNewSQLite")
	}

	return db
}

func MustNewTestDB(t *testing.T) *gorm.DB {
	const dbName = "test_storage.db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		require.NoError(t, err)
	}

	db.AutoMigrate(&models.Note{})

	t.Cleanup(func() {
		dbInstance, err := db.DB()
		require.NoError(t, err)
		require.NoError(t, dbInstance.Close())
		require.NoError(t, os.Remove(dbName))
	})

	return db
}
