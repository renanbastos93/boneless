package internal

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/renanbastos93/boneless/pkg/tools"
)

const msgRunning = `
    __                           __                                ______    __     ____
   / /_   ____    ____   ___    / /  ___    _____   _____         / ____/   / /    /  _/
  / __ \ / __ \  / __ \ / _ \  / /  / _ \  / ___/  / ___/        / /       / /     / /  
 / /_/ // /_/ / / / / //  __/ / /  /  __/ (__  )  (__  )        / /___    / /___ _/ /   
/_.___/ \____/ /_/ /_/ \___/ /_/   \___/ /____/  /____/         \____/   /_____//___/  

running...
`

func Start() {
	mainFile := tools.FindMainFile()

	// Set SERVICEWEAVER_CONFIG environment variable if not already set
	if _, ok := os.LookupEnv("SERVICEWEAVER_CONFIG"); !ok {
		err := os.Setenv("SERVICEWEAVER_CONFIG", "./weaver.toml")
		if err != nil {
			panic("failed to set env var SERVICEWEAVER_CONFIG: " + err.Error())
		}
	}

	cmd := tools.NewCmd("go", "run", mainFile)
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
