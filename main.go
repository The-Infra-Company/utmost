// Package main acts as the entrypoint for utmost.
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/The-Infra-Company/utmost/cmd"
)

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	go listenForInterrupt(stopChan)

	cmd.Execute()
}

func listenForInterrupt(stopScan chan os.Signal) {
	<-stopScan
	log.Fatal("Interrupt signal received, shutting down...")
}
