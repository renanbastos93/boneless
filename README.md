# boneless
Inspired by the character 'The Boneless,' a prominent figure in the Vikings saga, known for his exceptional agility and battlefield flexibility, our CLI project bears his name. The central idea behind this choice stems from his leadership characterized by strategic brilliance. Our primary objective is to enhance your daily workflow by facilitating the creation of a monolith that can be effortlessly decoupled and transformed into microservices. Leveraging open-source tools like Weaver, SQLC, gofiber and GOMIGRATE, we aim to provide comprehensive assistance throughout the process.


![image](https://github.com/renanbastos93/boneless/assets/8202898/46918810-9564-4ab6-b96a-933bca50fd94)


### Dependencies
 - [Service Wevar](https://serviceweaver.dev/)
 - [Fiber](https://gofiber.io/)
 - [SQLC](https://docs.sqlc.dev/en/stable/overview/install.html)
 - [Go Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)
 

## How to use that
```
$ go install github.com/renanbastos93/boneless/cmd/boneless@latest
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

## Architecture
We envisioned adopting an architecture similar to Hexagonal/Clean Architecture, which allows for the easy creation of applications in a decoupled and concise manner.

![image](https://github.com/renanbastos93/boneless/assets/8202898/b2ca8d54-46a6-4a19-bebc-94938f438cd5)


## Directory Structure
```sh
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── component.go
│   │   ├── db
│   │   │   ├── migrations
│   │   │   │   └── schema.sql
│   │   │   ├── query.sql
│   │   │   └── sqlc.yaml
│   │   ├── entity.go
│   │   ├── store
│   │   │   ├── db.go
│   │   │   ├── models.go
│   │   │   └── query.sql.go
│   │   └── weaver_gen.go
│   └── bff
│       ├── bff.go
│       ├── router.go
│       └── weaver_gen.go
└── weaver.toml
```


## Links
- [Boneless: a CLI to create your apps with Go](https://dev.to/renanbastos93/boneless-a-cli-to-create-your-apps-with-go-31kh)
- Building a Project from Scratch with Boneless CLI (WIP)
