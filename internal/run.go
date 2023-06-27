package internal

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

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
	fmt.Printf("running")

	err = cmd.Wait()
	if err != nil {
		println(err.Error())
	}
}
