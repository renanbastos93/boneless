package internal

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/renanbastos93/boneless/templates"
)

type contentTpl struct {
	ComponentName string
	Module        string
}

func generateFile(filePath string) (fd *os.File) {
	_, err := os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		if idx := strings.LastIndex(filePath, "/"); idx != -1 {
			err = os.MkdirAll(filePath[:idx], os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}

	fd, err = os.Create(filePath)
	if err != nil {
		panic(err)
	}

	return fd
}

func whatFilesGen(tpl []tplFiles, kind string) (data []tplFiles) {
	if kind == KindAll {
		return tpl
	}

	for _, t := range tpl {
		if t.kind == kind {
			data = append(data, t)
		}
	}

	return data
}

func Build(appName string, which string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ct := contentTpl{
		ComponentName: appName,
		Module:        GetProjectName(pwd),
	}

	tpls := whatFilesGen(templateFiles, which)
	for _, tf := range tpls {
		filePath := pwd + tf.filePath
		if tf.kind == KindComponent {
			filePath = fmt.Sprintf(filePath, appName)
		}
		fd := generateFile(filePath)
		t, err := template.ParseFS(templates.CreateTemplateFS, tf.pattern)
		if err != nil {
			panic(err)
		}
		t.Execute(fd, ct)
		_ = fd.Close()
		fmt.Println("created file:", filePath)
	}
}
