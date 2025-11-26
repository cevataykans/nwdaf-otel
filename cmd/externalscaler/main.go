package main

import (
	"context"
	"log"
	"nwdaf-otel/server/externalscaler"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	shutdownChn := make(chan struct{})
	srv := externalscaler.NewExternalScalerServer()
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Stop(ctx)
		close(shutdownChn)
	}()
	srv.Start()

	<-shutdownChn
	log.Println("External Scaler Finished!")
}
