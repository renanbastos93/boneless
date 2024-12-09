package internal

import (
	"bytes"
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

func sqlcGenerateByComponent(filePath string) {
	if stat, _ := os.Stat(filePath); stat == nil || stat.IsDir() {
		return
	}

	_ = runCmd(sqlcCLI, "generate", "-f", filePath)
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
			_ = runCmd(sqlcCLI, "generate", "-f", path)
		}
		return nil
	})
}

func WeaverGenerate() {
	err := runCmd(weaverCLI, "generate", "./...")
	if err != nil {
		panic(err)
	}
}

func RunMakeMigrate(componentName string, name string) {
	dir := findComponentPath(componentName)
	err := runCmd(migrateCLI, "create", "-ext", "sql", "-dir", dir+"/db/migrations/", name)
	if err != nil {
		panic(err)
	}
}

func ModTidy() {
	err := runCmd(goCLI, "mod", "tidy")
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
	wasInstalledDriversDB()
	var dir = findComponentPath(componentName)

	queryConn := ReadToml(componentName)
	_ = runCmd(migrateCLI, "-path", dir+"/db/migrations/", "-database", queryConn, "-verbose", upDown)
}

func wasInstalledDriversDB() {
	cmd := exec.Command(migrateCLI, "-h")
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
