# ✋ Welcome to Boneless CLI
A CLI (Command Line Interface) to create your apps with Service Weaver, using gomigrate, sqlc, and Fiber.


## Why was the name 'Boneless' chosen for this CLI tool?
The choice of the name "Boneless" for this CLI tool draws inspiration from the historical figure Ivar the Boneless. Ivar the Boneless was a legendary Viking leader known for his strategic prowess, adaptability, and agility on the battlefield.
In a similar vein, the name "Boneless" for the CLI tool reflects these qualities in the context of software development. It symbolizes the tool's ability to navigate through complex tasks and challenges effortlessly, just as Ivar the Boneless maneuvered through battles with agility and cunning.
By associating the CLI tool with Ivar the Boneless, the name not only captures the spirit of adaptability and flexibility but also adds a touch of historical significance to the tool's identity.
In summary, the name "Boneless" pays homage to Ivar the Boneless and serves as a metaphor for the CLI tool's ability to conquer development obstacles with ease and grace.

## Which are the tools used behind the Boneless?
1. **[Service Weaver](https://serviceweaver.dev/)**: Service Weaver is a tool that facilitates the creation and management of microservices. It helps in defining service boundaries, handling service discovery, and managing communication between microservices.
2. **[go-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)**: go-migrate is a database migration tool for Go applications. It allows developers to manage database schema changes and versioning in a structured and automated manner, ensuring smooth database migrations across different environments.
3. **[SQLC](https://docs.sqlc.dev/en/stable/overview/install.html)**: SQLC is a tool that generates type-safe Go code based on SQL queries. It helps in writing database code by automating the process of generating Go code from SQL queries, reducing boilerplate code and enhancing type safety.
4. **[Fiber](https://gofiber.io/)**: Fiber is a fast and efficient web framework for Go. It provides a lightweight and easy-to-use foundation for building web applications. Fiber offers features like routing, middleware support, and performance optimizations, making it a popular choice for developing high-performance web services.

These tools, when used alongside Boneless, contribute to various aspects of the development process. Service Weaver aids in managing microservices, go-migrate simplifies database migrations, SQLC enhances database code generation, and Fiber provides a robust framework for building web applications. Together, they enhance the functionality and development experience of Boneless.


## 💡 About the architecture
Let's dive into the architecture used in Boneless. In this section, we will explore the intricacies and key details of the architecture that powers the project created by Boneless.

### What is the architecture used as the foundation for Boneless?
Boneless was inspired by Clean Architecture and offers templates that follow its core concepts. The templates provided by Boneless facilitate the creation of modular, maintainable, and testable codebases by promoting a clear separation of concerns and the independence of business logic from external dependencies. By using Boneless templates, developers can jumpstart their projects with a well-organized structure that aligns with Clean Architecture, enabling them to focus on implementing the domain-specific logic while adhering to best practices. Boneless empowers developers to build robust and scalable applications, leveraging the benefits of Clean Architecture for easier understanding, maintenance, and evolution over time.

![image](https://github.com/renanbastos93/boneless/assets/8202898/89162036-1352-4cbd-bff9-511cbbeb1021)

Based on Service Weaver, which generates connections between components using gRPC, has greatly facilitated the development of our applications.

### Here is the current repository structure, reflecting the adopted architecture:
```
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── component.go
│   │   ├── db
│   │   │   ├── migrations
│   │   │   │   └── schema.sql
│   │   │   ├── query.sql
│   │   │   └── sqlc.yaml
│   │   ├── entity.go
│   │   ├── store
│   │   │   ├── db.go
│   │   │   ├── models.go
│   │   │   └── query.sql.go
│   │   └── weaver_gen.go
│   └── bff
│       ├── bff.go
│       ├── router.go
│       └── weaver_gen.go
└── weaver.toml
```

In this structure, we can observe the organization of directories and files in the repository. The `cmd` directory contains the `main.go` file, which is responsible for starting the application. The `go.mod` and `go.sum` files are used to manage project dependencies.
The `internal` the directory is where the main implementation of the application resides. Inside, we have the `app` directory, which contains components related to the application's domain, such as entities and business logic. The `db` directory is used to store files related to the database layer, such as migrations and SQL queries.
The `store` directory contains files related to data storage, such as the implementation of database access. The `bff` directory contains files related to the implementation of the Backend for the Frontend layer.
Lastly, the `weaver.toml` file is used to configure Service Weaver, which facilitates generating connections between application components using gRPC.
This directory and file structure reflects the adopted architecture in the repository, following the principles of [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and facilitating the organization and maintenance of the application's source code.

## ✨ Getting started
Let's create our first project from scratch using Boneless!

### Installing dependencies
First, we need to install the binary of Service Weaver, Go Migrate, SQLC, and Boneless. Even so, I suggest you read the documentation of all of them on your official websites.
```sh
$ brew install golang-migrate
$ brew install sqlc
$ go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
$ go install github.com/renanbastos93/boneless/cmd/boneless@latest
```
To ensure a smooth setup of your Boneless project, let's install all the dependencies, including Boneless itself. If you're not using macOS, you can access their website for detailed instructions on installing Boneless and its dependencies specific to your operating system.

### Now that we have installed it, let's use the 'help' command to display the available options.
```sh
$ boneless help
Usage: boneless [target]

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
```
After that is installed, let's create our first project from scratch using Boneless, a framework based on clean architecture, to efficiently generate and organize the project structure, implement functionality, and deploy it with ease.

## 👷‍♂️ Creating Project
Let's dive right in and create a project from scratch, and then run it to see Boneless in action!

For starting we've been thinking of a simple API that uses CRUD.
```sh
# First, we need to create go.mod file
$ go mod init bone

# After that let's call command boneless to start a project
# As a default, it was created using SQLite3
$ boneless new

# Here we going to execute all migrations of the `internal/app`
$ boneless migrate app up

# Now, we can use the command `run`
boneless run

    __                           __                                ______    __     ____
   / /_   ____    ____   ___    / /  ___    _____   _____         / ____/   / /    /  _/
  / __ \ / __ \  / __ \ / _ \  / /  / _ \  / ___/  / ___/        / /       / /     / /  
 / /_/ // /_/ / / / / //  __/ / /  /  __/ (__  )  (__  )        / /___    / /___ _/ /   
/_.___/ \____/ /_/ /_/ \___/ /_/   \___/ /____/  /____/         \____/   /_____//___/  

running...

PID: <process-id>
BFF listener available on 127.0.0.1:8090

 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.48.0                   │ 
 │               http://127.0.0.1:8090               │ 
 │                                                   │ 
 │ Handlers ............. 6  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 40506 │ 
 └───────────────────────────────────────────────────┘ 

╭───────────────────────────────────────────────────╮
│ app        : main                                 │
│ deployment : <uuid>                               │
╰───────────────────────────────────────────────────╯
```

Wow, here we created your first app using Boneless. Now, we can be using it, let's use cURL for testing.
```sh
# Get All Examples
$ curl --location --request GET 'http://127.0.0.1:8090/examples'

# returns empty because we haven't data yet
# []

# Create an Example
$ curl --location --request POST 'http://127.0.0.1:8090/examples/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "Hello everyone"
}'

# After you created anything we can get an example just
# but you need to know the ID for it
$ curl --location --request GET 'http://127.0.0.1:8090/examples/:id'

# example data
{
        "id": 2104895800, # id is random
        "created_at": "2023-07-20T03:22:37.289-03:00",
        "message": "Hello everyone"
}
```

That's all folks!
