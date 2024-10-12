package bff

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"{{.Module}}/internal/{{.ComponentName}}"
)

type implBFF struct {
	weaver.Implements[weaver.Main]
	example weaver.Ref[{{.ComponentName}}.Component]
	bff     weaver.Listener `weaver:"bff"`

	f *fiber.App
}

func (e *implBFF) createRouter(ctx context.Context) {
	router := e.f.Use(logger.New(logger.ConfigDefault))

	grpExamples := router.Group("/examples")
	grpExamples.Get("/", e.GetAllExamples)
	grpExamples.Get("/:id", e.GetExampleById)
	grpExamples.Post("/", e.CreateExample)
}

func Server(ctx context.Context, e *implBFF) (err error) {
	fmt.Printf("BFF listener available on %v\n", e.bff)

	e.f = fiber.New()
	e.createRouter(ctx)
	return e.f.Listener(e.bff)
}