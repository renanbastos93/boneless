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
  help                                // show commands for use
  version                             // show version
  create-scratch                      // create project from scratch using Weaver + sqlc + go-migrate
  build                               // build Weaver component with SQLC
  make-migrate <app-name> <name>      // create a new migrate from app
  migrate <app-name>                  // run migrate from app
  create-app <app-name>               // create a new app based on the app for example later you can change that
  build-app <app-name>                // build using Weaver + SQLC
  run                                 // running project using Weaver single
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
