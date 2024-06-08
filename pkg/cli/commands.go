package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/renanbastos93/boneless"
	"github.com/renanbastos93/boneless/internal"
	"github.com/renanbastos93/boneless/pkg/tools"
)

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
	cmdRun           = "run"

	DefaultComponentName = "app"
)

const Usage = `Usage: boneless [target]

Targets:
  help                                     Show commands for use
  version                                  Show version
  new  <sql|sqlite3>                       Create a project from scratch using Weaver, SQLC, and go-migrate
  create-scratch <sql|sqlite3>             Create a project from scratch using Weaver, SQLC, and go-migrate
  build                                    Build the Weaver component with SQLC
  make-migrate <app-name> <name>           Create a new migration for an app
  migrate <app-name> <up|down>             Run migrations for an app
  create-app <app-name>                    Create a new app based on a template
  build-app <app-name>                     Build an app using Weaver and SQLC
  run                                      Run the project using Weaver

Parameters:
  <app-name>                               Name of the app to create or run migrations on
  <name>                                   Name of the migration to create
  <up|down>                                Specify "up" to apply migrations or "down" to rollback migrations
  <sql|sqlite>                             Specify "sql" to use some SQL "sqlite3" to use sqlite3 and it is the default

Examples:
  boneless help
  boneless version
  boneless create-scratch
  boneless build
  boneless make-migrate my-app migration-name
  boneless migrate my-app up
  boneless create-app my-app
  boneless build-app my-app
  boneless run

`

// TODO: valide if we should use `` or `flag.Arg(i)`

func newScratch() {
	opt := Options{
		AppName:       DefaultComponentName,
		WhichTemplate: All,
		WhichDatabase: flag.Arg(1),
	}

	New(opt).Build()
	tools.SqlcGenerate()
	tools.ModTidy()
	tools.WeaverGenerate()
}

var CmdToRun = map[string]func(){
	cmdhelp: func() {
		fmt.Fprint(os.Stdout, Usage)
	},
	cmdVersion: func() {
		fmt.Fprintln(os.Stdout, boneless.Version)
	},
	cmdNew: func() {
		newScratch()
	},
	cmdCreateScratch: func() {
		newScratch()
	},
	cmdCreateApp: func() {
		opt := Options{
			AppName:       flag.Arg(1),
			WhichTemplate: Component,
		}

		New(opt).Build()
		tools.SqlcGenerate(flag.Arg(1))
		tools.WeaverGenerate()
	},
	cmdBuild: func() {
		tools.SqlcGenerate()
		tools.WeaverGenerate()
	},
	cmdBuildApp: func() {
		tools.SqlcGenerate(flag.Arg(1))
		tools.WeaverGenerate()
	},
	cmdMakeMigrate: func() {
		tools.SqlcGenerate()
		tools.RunMakeMigrate(flag.Arg(1), flag.Arg(2))
	},
	cmdMigrate: func() {
		tools.RunMigrate(flag.Arg(1), flag.Arg(2))
	},
	cmdRun: func() {
		internal.Start()
	},
}
