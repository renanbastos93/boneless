package app

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/renanbastos93/boneless/templates/app/store"
)

type Component interface {
	AllExamples(ctx context.Context) (out AllExamplesOut, err error)
	GetOneExampleById(ctx context.Context, id int32) (out ExampleOut, err error)
	CreateExample(ctx context.Context, in ExampleIn) (err error)
}

type Config struct {
	Driver string
	Source string
}

type implAudit struct {
	weaver.Implements[Component]
	weaver.WithConfig[Config]
	db *store.Queries
}

func (e *implAudit) Init(ctx context.Context) error {
	db, err := sql.Open(e.Config().Driver, e.Config().Source)
	if err != nil {
		return fmt.Errorf("not open: %w", err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping: %w", err)
	}
	e.db = store.New(db)
	return nil
}

func (e implAudit) AllExamples(ctx context.Context) (out AllExamplesOut, err error) {
	examples, err := e.db.ListExamples(ctx)
	if err != nil {
		return out, err
	}
	return out.FromStore(examples), nil
}

func (e implAudit) GetOneExampleById(ctx context.Context, id int32) (out ExampleOut, err error) {
	example, err := e.db.GetExampleById(ctx, id)
	if err != nil {
		return out, err
	}
	return out.FromStore(example), nil
}

func (e implAudit) CreateExample(ctx context.Context, in ExampleIn) (err error) {
	_, err = e.db.CreateExample(ctx, in.ToStore())
	if err != nil {
		return err
	}
	return nil
}
