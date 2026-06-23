package app

import (
	"GoCeleryLite/internal/server"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Run() error {
	// Ctrl + c
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := server.New(":9384")

	return srv.Run(ctx)
}
