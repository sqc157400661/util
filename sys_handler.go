//go:build linux || darwin
// +build linux darwin

package util

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func ExitSignalHandler(onExit func()) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, os.Interrupt, os.Kill)
	select {
	case sig := <-quit:
		fmt.Printf("Quiting process for signal=%+v, PID=%d \n", sig, os.Getpid())
		if onExit != nil {
			onExit()
		}
		return
	}
}
