package main

import (
	"context"
	"log"

	"github.com/ServiceWeaver/weaver"
	_ "{{.Module}}/bff"
)

func main() {
	if err := weaver.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
