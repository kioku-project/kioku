package main

import (
	"context"
	"github.com/apex/log"
	"github.com/kioku-project/kioku/internal/service/login"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	srv := login.New()

	go func() {
		if err := srv.Listen(":80"); err != nil {
			stop()
		}
	}()

	select {
	case <-ctx.Done():
		if err := srv.Shutdown(); err != nil {
			log.WithError(err).Fatal("Could not shutdown the server gracefully")
		}
	}
}
