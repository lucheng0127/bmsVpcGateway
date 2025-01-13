package signals

import (
	"context"
	"os"
	"os/signal"
)

func SetupSignalHandler() context.Context {
	c := make(chan os.Signal, 2)
	ctx, cancel := context.WithCancel(context.Background())

	signal.Notify(c, shutdownSignals...)

	go func() {
		<-c
		cancel()
		<-c
		os.Exit(1)
	}()

	return ctx
}
