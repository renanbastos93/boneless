package tools

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func NewCmd(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func GetProjectName(dir string) (moduleName string) {
	modFile, err := os.Open(dir + "/go.mod")
	if err != nil {
		panic(fmt.Sprintf("go.mod does not exist: %v", err))
	}

	defer modFile.Close()
	_, err = fmt.Fscanf(modFile, "module %s", &moduleName)
	if err != nil {
		panic(fmt.Sprintf("read go mod error: %v", err))
	}

	return moduleName
}

func readToml(componentName string) (qsConn string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// TODO: improve that same internal/run.go:46
	weaverToml, err := ioutil.ReadFile(pwd + "/weaver.toml")
	if err != nil {
		panic(err)
	}

	var cfgWeaver map[string]interface{}
	err = toml.Unmarshal(weaverToml, &cfgWeaver)
	if err != nil {
		panic(err)
	}

	projectName := GetProjectName(pwd)
	componentModPath := projectName + "/internal/" + componentName + "/Component"
	for k, v := range cfgWeaver {
		if k != componentModPath {
			continue
		}
		if value, ok := v.(map[string]interface{}); ok {
			var driver = ""
			var source = ""
			for nk, nv := range value {
				switch nk {
				case "Driver":
					driver = nv.(string)
				case "Source":
					source = nv.(string)
				}
			}
			qsConn = fmt.Sprintf("%s://%s", driver, source)
		}
	}

	if qsConn == "" {
		panic("not found settings to running migrate verify your weaver.toml file")
	}

	return qsConn
}

func findComponentPath(componentName string) (dir string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filepath.Walk(pwd, func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() && info.Name() == componentName {
			dir = path
		}
		return nil
	})

	return dir
}

func FindMainFile() (filePath string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// TODO: improve that to ensure that the file is the package main
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

func SqlcGenerate(componentName ...string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if len(componentName) > 0 {
		dir := findComponentPath(componentName[0])
		sqlcGenerateByComponent(dir + "/db/sqlc.yaml")
		return
	}

	filepath.Walk(pwd, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), "sqlc.yaml") {
			_ = NewCmd("sqlc", "generate", "-f", path).Run()
		}
		return nil
	})
}

func WeaverGenerate() {
	err := NewCmd("weaver", "generate", "./...").Run()
	if err != nil {
		panic(err)
	}
}

func RunMakeMigrate(componentName string, name string) {
	dir := findComponentPath(componentName)
	err := NewCmd("migrate", "create", "-seq", "-ext", "sql", "-dir", dir+"/db/migrations/", name).Run()
	if err != nil {
		panic(err)
	}
}

func ModTidy() {
	err := NewCmd("go", "mod", "tidy").Run()
	if err != nil {
		panic("failed to run `go mod tidy`")
	}
}

func RunMigrate(componentName, upDown string) {
	wasInstalledDriversDB()
	var dir = findComponentPath(componentName)

	queryConn := readToml(componentName)
	_ = NewCmd("migrate", "-path", dir+"/db/migrations/", "-database", queryConn, "-verbose", upDown).Run()
}

func wasInstalledDriversDB() {
	cmd := NewCmd("migrate", "-h")
	var errb bytes.Buffer
	cmd.Stderr = &errb
	_ = cmd.Run()

	out := errb.String()
	if i := strings.LastIndex(out, "Database drivers: "); i > -1 {
		out = out[i:]
		out = out[:strings.LastIndex(out, "\n")]
		if !strings.Contains(out, "mysql") || !strings.Contains(out, "sqlite3") {
			msg := `not supported drivers: mysql or sqlite3.
Install go-migrate with tags: go install -tags 'sql mysql sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
Or read documentation from github.com/golang-migrate/migrate`
			println(msg)
		}
	}
}

func sqlcGenerateByComponent(filePath string) {
	if stat, _ := os.Stat(filePath); stat == nil || stat.IsDir() {
		return
	}
	_ = NewCmd("sqlc", "generate", "-f", filePath).Run()
}
