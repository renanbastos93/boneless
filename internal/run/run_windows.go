//go:build windows
// +build windows

package run

import (
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
)

func CloseBySignal(pid int) {
	println("PID:", pid)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	select {
	default:
		switch <-ch {
		case syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM:
			err := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(pid)).Run()
			if err != nil {
				panic(err)
			}
		}
	}
}
