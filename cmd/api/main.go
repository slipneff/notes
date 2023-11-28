package main

import (
	"context"
	"log"

	"github.com/slipneff/notes/internal/di"
	"github.com/slipneff/notes/internal/utils/config"
	"github.com/slipneff/notes/internal/utils/flags"
)

func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	container := di.New(context.Background(), cfg)

	err := container.GetHttpServer().ListenAndServe()
	if err != nil {
		log.Panic(err, "Fail serve HTTP")
	}
}
