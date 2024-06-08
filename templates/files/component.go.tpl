package {{.ComponentName}}

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ServiceWeaver/weaver"

	"{{.Module}}/internal/{{.ComponentName}}/store"
{{if .IsSQL}}
    _ "github.com/go-sql-driver/mysql"
	{{- end}}
	{{- if .IsSQLLite3}}
	_ "github.com/mattn/go-sqlite3"
	{{- end}}
)

type Component interface {
	AllExamples(ctx context.Context) (out AllExamplesOut, err error)
	GetOneExampleById(ctx context.Context, id int32) (out ExampleOut, err error)
	CreateExample(ctx context.Context, in ExampleIn) (ok bool, err error)
}

type Config struct {
	Driver string
	Source string
}

type implapp struct {
	weaver.Implements[Component]
	weaver.WithConfig[Config]
	dbConn *sql.DB
	db     *store.Queries
}

func (e *implapp) Init(ctx context.Context) error {
	db, err := sql.Open(e.Config().Driver, e.Config().Source)
	if err != nil {
		return fmt.Errorf("not open: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping: %w", err)
	}

	e.dbConn = db
	e.db = store.New(db)
	return nil
}


func (e *implapp) Shutdown(ctx context.Context) error {
	// TODO: create your logic to shutdown the component
	return e.dbConn.Close()
}

func (e implapp) AllExamples(ctx context.Context) (out AllExamplesOut, err error) {
	examples, err := e.db.ListExamples(ctx)
	if err != nil {
		return out, err
	}
	return out.FromStore(examples), nil
}

func (e implapp) GetOneExampleById(ctx context.Context, id int32) (out ExampleOut, err error) {
	example, err := e.db.GetExampleById(ctx, id)
	if err != nil {
		return out, err
	}
	return out.FromStore(example), nil
}

func (e implapp) CreateExample(ctx context.Context, in ExampleIn) (ok bool, err error) {
	_, err = e.db.CreateExample(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}
