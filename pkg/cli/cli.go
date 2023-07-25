package cli

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/renanbastos93/boneless/pkg/tools"
	"github.com/renanbastos93/boneless/templates"
)

type CLI interface {
	Build()
}

type contentTpl struct {
	ComponentName string
	Module        string
	IsSQLLite3    bool
	IsSQL         bool
}

type implCli struct {
	whichDatabase    DBType
	whichTemplate    TemplateType
	whichTemplateStr string
	pwd              string
	projectName      string
	appName          string
}

type Options struct {
	AppName       string
	WhichTemplate string
	WhichDatabase string
}

func New(opt Options) CLI {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	e := &implCli{
		whichDatabase:    StrToDBType[opt.WhichDatabase],
		whichTemplate:    StrToTemplateType[opt.WhichTemplate],
		whichTemplateStr: opt.WhichTemplate,
		pwd:              pwd,
		projectName:      tools.GetProjectName(pwd),
		appName:          opt.AppName,
	}

	return e
}

func (e *implCli) Build() {
	if e.whichTemplate == 0 {
		panic(fmt.Sprintf("failed to found templates for this type: %s", e.whichTemplateStr))
	}

	if e.whichDatabase == 0 {
		e.whichDatabase = SQLite3Type
	}

	content := &contentTpl{
		ComponentName: e.appName,
		Module:        e.projectName,
		IsSQLLite3:    false,
		IsSQL:         false,
	}

	content.setDatabaseToUse(e.whichDatabase)
	templateFiles := TemplateTypeToFilePath[e.whichTemplate]

	for _, tf := range templateFiles {
		fp := filepath.Join(e.pwd, tf.Filepath)
		if strings.Contains(fp, "%s") {
			fp = fmt.Sprintf(fp, e.appName)
		}
		fd := e.generateFile(fp)
		t, err := template.ParseFS(templates.CreateTemplateFS, tf.Pattern)
		if err != nil {
			panic(err)
		}
		_ = t.Execute(fd, content)
		_ = fd.Close()
		println("created file: " + fd.Name())
	}
}

func (e *implCli) generateFile(filepath string) (fd *os.File) {
	info, err := os.Stat(filepath)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		if idx := strings.LastIndex(filepath, "/"); idx != -1 {
			err = os.MkdirAll(filepath[:idx], os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}

	// It will ensure that don't replace files in case you repeat `boneless new` command
	if info != nil {
		fd, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}
		return fd
	}

	fd, err = os.Create(filepath)
	if err != nil {
		panic(err)
	}

	return fd
}

func (e *contentTpl) setDatabaseToUse(whichSql DBType) {
	switch whichSql {
	case SQLType, MySQLType:
		e.IsSQL = true
	default:
		e.IsSQLLite3 = true
	}
}
