package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/renanbastos93/boneless/internal"
)

const usage = `Usage: boneless [target]

Targets:
  help                                     Show commands for use
  version                                  Show version
  new <sql|sqlite3>                        Create a project from scratch using Weaver, SQLC, and go-migrate
  create-scratch <sql|sqlite3>             Create a project from scratch using Weaver, SQLC, and go-migrate
  build                                    Build the Weaver component with SQLC
  make-migrate <app-name> <name>           Create a new migration for an app
  migrate <app-name> <up|down>             Run migrations for an app
  create-app <app-name>                    Create a new app based on a template
  build-app <app-name>                     Build an app using Weaver and SQLC
  delete-app <app-name>                    Delete an app created
  install-deps [package]          Installs external dependencies required by Boneless (e.g., weaver, sqlc). If no package is specified, all dependencies are updated.
  update-deps [package]           Updates the specified external dependency (e.g., weaver, migrate). If no package is specified, all dependencies are updated.
  run                                      Run the project using Weaver

Parameters:
  <app-name>                               Name of the app to create or run migrations on
  <name>                                   Name of the migration to create
  <up|down>                                Specify "up" to apply migrations or "down" to rollback migrations
  <sql|sqlite>                             Specify "sql" to use some SQL "sqlite3" to use sqlite3 and it is the default
  [package]                                Specify the package to update or install like "weaver, sqlc, golang-migrate"


Examples:
  boneless help
  boneless version
  boneless create-scratch
  boneless build
  boneless make-migrate my-app migration-name
  boneless migrate my-app up
  boneless create-app my-app
  boneless build-app my-app
  boneless delete-app my-app
  boneless install-deps
  boneless install-deps weaver
  boneless update-deps
  boneless update-deps sqlc
  boneless run

`

const (
	cmdhelp          = "help"
	cmdVersion       = "version"
	cmdNew           = "new"
	cmdCreateScratch = "create-scratch"
	cmdBuild         = "build"
	cmdMakeMigrate   = "make-migrate"
	cmdMigrate       = "migrate"
	cmdCreateApp     = "create-app"
	cmdBuildApp      = "build-app"
	cmdDeleteApp     = "delete-app"
	cmdRun           = "run"
	cmdInstallDeps   = "install-deps"
	cmdUpdateDeps    = "update-deps"

	DefaultComponentName = "app"
)

func init() {
	internal.ValidateLatestVersion()
}

func main() {
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case cmdhelp:
		fmt.Fprint(os.Stdout, usage)
	case cmdVersion:
		fmt.Fprintln(os.Stdout, internal.Version)
	case cmdCreateScratch, cmdNew:
		internal.ModInit()
		internal.Build(DefaultComponentName, internal.KindAll, flag.Arg(1))
		internal.SqlcGenerate()
		internal.ModTidy()
		internal.WeaverGenerate()
	case cmdCreateApp:
		internal.Build(flag.Arg(1), internal.KindComponent, "")
		internal.SqlcGenerate(flag.Arg(1))
		internal.WeaverGenerate()
	case cmdDeleteApp:
		internal.DeleteApp(flag.Arg(1))
	case cmdBuild:
		internal.SqlcGenerate()
		internal.WeaverGenerate()
	case cmdBuildApp:
		internal.SqlcGenerate(flag.Arg(1))
		internal.WeaverGenerate()
	case cmdMakeMigrate:
		internal.SqlcGenerate()
		internal.RunMakeMigrate(flag.Arg(1), flag.Arg(2))
	case cmdMigrate:
		internal.RunMigrate(flag.Arg(1), flag.Arg(2))
	case cmdRun:
		internal.Start()
	case cmdInstallDeps:
		internal.InstallDeps(flag.Arg(1))
	case cmdUpdateDeps:
		internal.UpdateDeps(flag.Arg(1))
	default:
		flag.Usage()
	}
}
