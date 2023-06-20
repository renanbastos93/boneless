package bff

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/renanbastos93/boneless/templates/app"
)

type implBFF struct {
	weaver.Implements[weaver.Main]
	example weaver.Ref[app.Component]
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

func (e *implBFF) Main(ctx context.Context) error {
	fmt.Printf("BFF listener available on %v\n", e.bff.String())

	f := fiber.New()
	e.createRouter(ctx)
	return f.Listener(e.bff)
}
