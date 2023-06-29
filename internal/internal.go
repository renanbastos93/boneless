package internal

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func findComponentPath(componentName string) (dir string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(pwd, func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() && info.Name() == componentName {
			dir = path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return dir
}

func SqlcGenerateByComponent(componentName string) {
	dir := findComponentPath(componentName)
	err := runCmd("sqlc", "generate", "-f", dir+"/db/sqlc.yaml")
	if err != nil {
		panic(err)
	}
}

func SqlcGenerate(componentName ...string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filepath.Walk(pwd, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), "sqlc.yaml") {
			err := runCmd("sqlc", "generate", "-f", path)
			if err != nil {
				fmt.Println("warn: can't running sqlc in", path)
			}
		}
		return nil
	})
}

func WeaverGenerate() {
	err := runCmd("weaver", "generate", "./...")
	if err != nil {
		panic(err)
	}
}

func RunMakeMigrate(componentName string, name string) {
	dir := findComponentPath(componentName)
	err := runCmd("migrate", "create", "-ext", "sql", "-dir", dir+"/db/migrations/", name)
	if err != nil {
		panic(err)
	}
}

func ModTidy() {
	err := runCmd("go", "mod", "tidy")
	if err != nil {
		panic(err)
	}
}

func runCmd(name string, args ...string) (err error) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunMigrate(componentName, upDown string) {
	dir := findComponentPath(componentName)
	queryConn := ReadToml(componentName)
	err := runCmd("migrate", "-path", dir+"/db/migrations/", "-database", queryConn, "-verbose", upDown)
	if err != nil {
		panic(err)
	}
}
