package internal

import (
	"io/fs"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const msgRunning = `
    __                           __                                ______    __     ____
   / /_   ____    ____   ___    / /  ___    _____   _____         / ____/   / /    /  _/
  / __ \ / __ \  / __ \ / _ \  / /  / _ \  / ___/  / ___/        / /       / /     / /  
 / /_/ // /_/ / / / / //  __/ / /  /  __/ (__  )  (__  )        / /___    / /___ _/ /   
/_.___/ \____/ /_/ /_/ \___/ /_/   \___/ /____/  /____/         \____/   /_____//___/  

running...
`

func findMainFile() (filePath string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	_ = filepath.Walk(pwd, func(path string, info fs.FileInfo, _ error) error {
		if !info.IsDir() {
			if info.Name() == "main.go" {
				filePath = path
			}
		}
		return nil
	})

	return filePath
}

func Start() {
	mainFile := findMainFile()
	cmd := exec.Command("go", "run", mainFile)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		panic("cmd run failed: " + err.Error())
	}
	time.Sleep(time.Second)
	println(msgRunning)

	go CloseBySignal(cmd.Process.Pid)

	err = cmd.Wait()
	if err != nil {
		println(err.Error())
	}
}

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
