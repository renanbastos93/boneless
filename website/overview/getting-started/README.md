---
description: Let's create our first project from scratch using Boneless!
---

# ✨ Getting started

## Installing dependencies

First, we need to install the binary of Service Weaver, Go Migrate, SQLC, and Boneless. Even so, I suggest you read the documentation of all of them on your official websites.

```sh
$ brew install golang-migrate
$ brew install sqlc
$ go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
$ go install github.com/renanbastos93/boneless/cmd/boneless@latest
```

To ensure a smooth setup of your Boneless project, let's install all the dependencies, including Boneless itself. If you're not using macOS, you can access their website for detailed instructions on installing Boneless and its dependencies specific to your operating system.

#### Now that we have installed it, let's use the 'help' command to display the available options.

```sh
$ boneless help
Usage: boneless [target]

Targets:
  help                                     Show commands for use
  version                                  Show version
  create-scratch                           Create a project from scratch using Weaver, SQLC, and go-migrate
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
```

After that is installed, let's create our first project from scratch using Boneless, a framework based on clean architecture, to efficiently generate and organize the project structure, implement functionality, and deploy it with ease.&#x20;

#### Once the project is set up, we can pass to the next page to create that.

