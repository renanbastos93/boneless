package internal

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"text/template"
)

func BuildStartProject(name string, path ...string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	p := pwd + "/" + name
	if len(path) == 1 {
		p = path[0] + "/" + name
	}

	var fds = make([]fd, 0, 12)
	filepath.Walk(pwd, func(path string, info fs.FileInfo, _ error) error {
		if !info.IsDir() && strings.Contains(path, "templates") {
			currentPath := strings.Replace(path, "tpl", "go", -1)
			_, currentPath, _ = strings.Cut(currentPath, "templates")
			content, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			fmt.Println(currentPath, strings.Count(currentPath, "/"))
			fds = append(fds, fd{
				name:      info.Name(),
				content:   content,
				finalPath: currentPath,
				tpl:       template.New(info.Name()),
			})
			return nil
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
			if idx := strings.LastIndex(path, "/"); idx != -1 {
				err = os.MkdirAll(path[:idx], os.ModePerm)
				if err != nil {
					panic(err)
				}
			}
		}

		fd, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		if err = t.Execute(fd, ct); err != nil {
			panic(err)
		}
		err = fd.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("created file:", path)
	}
}

type fd struct {
	name      string
	content   []byte
	finalPath string
	tpl       *template.Template
}

type contentTpl struct {
	ComponentName string
	Module        string
}
