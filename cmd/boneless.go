package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/renanbastos93/boneless/internal"
)

const usage = `Usage: boneless [target]
  help                                // show commands for use
  version                             // show version
  create-scratch                      // create project from scratch using Weaver + sqlc + go-migrate
  build                               // build Weaver component with SQLC
  make-migrate <app-name> <name>      // create a new migrate from app
  migrate <app-name>                  // run migrate from app
  create-app <app-name>               // create a new app based on app for example later you can change that
  build-app <app-name>                // build using Weaver + SQLC
  run                                 // running project using Weaver single
`

const (
	cmdhelp          = "help"
	cmdVersion       = "version"
	cmdCreateScratch = "create-scratch"
	cmdBuild         = "build"
	cmdMakeMigrate   = "make-migrate"
	cmdMigrate       = "migrate"
	cmdCreateApp     = "create-app"
	cmdBuildApp      = "build-app"
	cmdRun           = "run"

	DefaultComponentName = "app"
)

func main() {
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case cmdhelp:
		n := len(flag.Args())
		command := flag.Arg(1)
		switch n {
		case 1:
			fmt.Fprint(os.Stdout, usage)
		case 2:
			// TODO: implement more details about help
			fmt.Fprintln(os.Stdout, "boneless help "+command)
		}
	case cmdVersion:
		fmt.Fprintln(os.Stdout, internal.Version)
	case cmdCreateScratch:
		internal.Build(DefaultComponentName, internal.KindAll)
		internal.SqlcGenerate()
		internal.ModTidy()
		internal.WeaverGenerate()
	case cmdCreateApp:
		internal.Build(flag.Arg(1), internal.KindComponent)
		internal.SqlcGenerateByComponent(flag.Arg(1))
		internal.WeaverGenerate()
	case cmdBuild:
		internal.SqlcGenerate()
		internal.WeaverGenerate()
	case cmdBuildApp:
		internal.WeaverGenerate()
		internal.SqlcGenerateByComponent(flag.Arg(1))
	case cmdMakeMigrate:
		internal.SqlcGenerate()
		internal.RunMakeMigrate(flag.Arg(1), flag.Arg(2))
	case cmdMigrate:
		// TODO: use this command
		fmt.Fprintln(os.Stdout, `not implemented`)
	case cmdRun:
		internal.Start()
	}
}
