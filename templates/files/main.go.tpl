package main

import (
	"context"
	"log"

	"github.com/ServiceWeaver/weaver"
	"{{.Module}}/internal/bff"
)

func main() {
	if err := weaver.Run(context.Background(), bff.Server); err != nil {
		log.Fatal(err)
	}
}
