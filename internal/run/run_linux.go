//go:build !plan9 && !windows
// +build !plan9,!windows

package run

import (
	"os"
	"os/signal"
	"syscall"
)

func CloseBySignal(pid int) {
	println("PID:", pid)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	select {
	default:
		switch <-ch {
		case syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM:
			syscall.Kill(-pid, syscall.SIGKILL)
		}
	}
}
