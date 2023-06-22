package internal

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
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

func RunSqlcGenerate(componentName string) {
	dir := findComponentPath(componentName)
	cmd := exec.Command("sqlc", "generate", "-f", dir+"/db/sqlc.yaml")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func RunWeaverGenerate() {
	cmd := exec.Command("weaver", "generate", "./...")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func RunMakeMigrate(componentName string, name string) {
	dir := findComponentPath(componentName)
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", dir+"/db/migrations/", name)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func GenerateQueryByEntity() {}
