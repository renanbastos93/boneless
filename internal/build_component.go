package internal

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"text/template"
)

func BuildComponent(name string, path ...string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	p := pwd + "/" + name
	if len(path) == 1 {
		p = path[0] + "/" + name
	}

	var fds = make([]fd, 0, 8)
	filepath.Walk(pwd, func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() && strings.Contains(path, "app") {
			fmt.Println("path", path)
			_, currentPath, _ := strings.Cut(path, "app")
			if len(currentPath) > 0 {
				err := os.MkdirAll(p+currentPath, os.ModePerm)
				if err != nil {
					panic(err)
				}
			}
			return nil
		}
		if !info.IsDir() && strings.Contains(path, "app") {
			switch {
			case strings.HasSuffix(path, ".tpl"),
				strings.HasSuffix(path, ".go"),
				strings.HasSuffix(path, ".yaml"),
				strings.HasSuffix(path, ".sql"),
				strings.HasSuffix(path, ".toml"):
				currentPath := strings.Replace(path, ".tpl", ".go", -1)
				currentPath = strings.Replace(currentPath, "app.go", name+".go", -1)
				_, currentPath, _ = strings.Cut(currentPath, "app")
				content, err := ioutil.ReadFile(path)
				if err != nil {
					panic(err)
				}
				fds = append(fds, fd{
					name:      info.Name(),
					content:   content,
					finalPath: currentPath,
					tpl:       template.New(info.Name()),
				})
			}
		}
		return nil
	})

	ct := contentTpl{
		ComponentName: name,
		// TODO: get mod
		Module: "gomodinit",
	}

	for _, f := range fds {
		t, err := f.tpl.Parse(string(f.content))
		if err != nil {
			panic(err)
		}
		path := p + f.finalPath
		_, err = os.Stat(path)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(err)
			}
		}

		fd, err := os.Create(path)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(err)
			}
		}
		defer fd.Close()
		if err = t.Execute(fd, ct); err != nil {
			panic(err)
		}
		fmt.Println("created file:", path)
	}
}
