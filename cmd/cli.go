package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `usage: boneless [target]
  help                         // show commands for use
  version                      // show version
  create-scratch               // create project from scratch using Weaver + sqlc + go-migrate
  build                        // build Weaver component with SQLC
  make-migrate <app-name       // create a new migrate from app
  migrate <app-name>           // run migrate from app
  create-app <app-name>        // create a new app based on app for example later you can change that
  build-app <app-name>         // build using Weaver + SQLC
  run                          // running project using Weaver single
`

var (
	cmdhelp          = "help"
	cmdVersion       = "version"
	cmdCreateScratch = "create-scratch"
	cmdBuild         = "build"
	cmdMakeMigrate   = "make-migrate"
	cmdMigrate       = "migrate"
	cmdCreateApp     = "create-app"
	cmdBuildApp      = "build-app"
	cmdRun           = "run"
)

const version = "v0.0.1"

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
			fmt.Fprintln(os.Stdout, "boneless help "+command)
		}
	case cmdVersion:
		fmt.Fprintln(os.Stdout, version)
	case cmdCreateScratch:
		fallthrough
	case cmdBuild:
		fallthrough
	case cmdCreateApp:
		fallthrough
	case cmdBuildApp:
		fallthrough
	case cmdMakeMigrate:
		fallthrough
	case cmdMigrate:
		fallthrough
	case cmdRun:
		fmt.Fprintln(os.Stdout, `not implemented`)
	}

	code, err := run(flag.Args()[0], flag.Args()[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(code)
	}
	// boneless sqlc generate
	// boneless build start
	// internal.RunSqlcGenerate("app")
	// internal.BuildStartProject("zig")
	// internal.BuildComponent("audit")
}

func run(deployer string, args []string) (int, error) {
	return -1, nil
}
