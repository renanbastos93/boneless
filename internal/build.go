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
	IsSQLLite3    bool
	IsSQL         bool
}

const (
	SQL      = "sql"
	SQLlite3 = "sqlite3"
)

func (e *contentTpl) setDatabaseToUse(whichSql string) {
	switch strings.ToLower(whichSql) {
	case SQL:
		e.IsSQL = true
	default:
		e.IsSQLLite3 = true
	}
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

func Build(appName string, which string, whichSql string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ct := &contentTpl{
		ComponentName: appName,
		Module:        GetProjectName(pwd),
		IsSQLLite3:    false,
		IsSQL:         false,
	}

	ct.setDatabaseToUse(whichSql)

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
