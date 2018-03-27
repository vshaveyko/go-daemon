// Package main provides ...
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/takama/daemon"
	"github.com/vshaveyko/test-go-daemon/dlog"
	"github.com/vshaveyko/test-go-daemon/process"
)

const (
	name        = "luigitired"
	description = "Get pipelines to run from database, schedule, loop"
	cadence     = time.Second
)

// Service has embedded daemon
type Service struct {
	daemon.Daemon
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {

	usage := "Usage: myservice start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	go (func() {
		for {
			process.Entrypoint()

			time.Sleep(cadence)
		}
	})()

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case killSignal := <-interrupt:
			dlog.Stdlog.Println("Got signal:", killSignal)

			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}

			return "Daemon was killed", nil
		}
	}

	// never happen, but need to complete code
	return usage, nil
}

func main() {
	srv, err := daemon.New(name, description)
	if err != nil {
		dlog.Errlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		dlog.Errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
