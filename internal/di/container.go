package di

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/slipneff/notes/internal/pkg/server"
	note "github.com/slipneff/notes/internal/pkg/service"
	"github.com/slipneff/notes/internal/pkg/storage/sql"
	"github.com/slipneff/notes/internal/utils/config"
	"gorm.io/gorm"
)

type Container struct {
	cfg                *config.Config
	ctx                context.Context
	noteService        *note.Service
	netListener        *net.Listener
	handler            *server.Handler
	storage            *sql.Storage
	db                 *gorm.DB
	transactionManager trm.Manager
	httpServer         *http.Server
}

func New(ctx context.Context, cfg *config.Config) *Container {
	return &Container{cfg: cfg, ctx: ctx}
}

func (c *Container) GetHttpServer() *http.Server {
	return get(&c.httpServer, func() *http.Server {
		return &http.Server{
			Addr:           fmt.Sprintf("%s:%d", c.cfg.Host, c.cfg.Port),
			Handler:        c.GetHttpHandler().InitRoutes(),
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		}
	})
}

func (c *Container) GetPostgresDB() *sql.Storage {
	return get(&c.storage, func() *sql.Storage {
		return sql.New(c.GetDB(), trmgorm.DefaultCtxGetter)
	})
}

func (c *Container) GetDB() *gorm.DB {
	return get(&c.db, func() *gorm.DB {
		return sql.MustNewPostgresDB(c.cfg)
	})
}

func (c *Container) GetTransactionManager() trm.Manager {
	return get(&c.transactionManager, func() trm.Manager {
		return manager.Must(trmgorm.NewDefaultFactory(c.GetDB()))
	})
}

func (c *Container) GetNoteService() *note.Service {
	return get(&c.noteService, func() *note.Service {
		return note.New(c.GetPostgresDB())
	})
}

func (c *Container) GetHttpHandler() *server.Handler {
	return get(&c.handler, func() *server.Handler {
		return server.New(c.cfg, c.GetNoteService())
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}

	*obj = builder()
	return *obj
}
