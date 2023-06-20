package main

import (
	"context"
	"log"

	"github.com/ServiceWeaver/weaver"
	_ "github.com/renanbastos93/boneless/templates/bff"
)

func main() {
	if err := weaver.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
