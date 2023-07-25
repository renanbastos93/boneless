package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/renanbastos93/boneless/internal"
	"github.com/renanbastos93/boneless/pkg/cli"
)

func init() {
	internal.ValidateLatestVersion()
}

func main() {
	flag.Usage = func() { fmt.Fprint(os.Stderr, cli.Usage) }
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprint(os.Stderr, cli.Usage)
		os.Exit(1)
	}

	if execute, ok := cli.CmdToRun[flag.Arg(0)]; ok {
		execute()
	} else {
		flag.Usage()
	}
}
