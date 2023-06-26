package internal

var filesDefault = []string{
	"/weaver.toml",
	"/cmd/main.go",
}

var filesBff = []string{
	"/internal/bff/bff.go",
	"/internal/bff/router.go",
}

var filesComponent = []string{
	"/internal/%s/component.go",
	"/internal/%s/entity.go",
	"/internal/%s/db/migrations/schema.sql",
	"/internal/%s/db/sqlc.yaml",
	"/internal/%s/db/query.sql",
}

var templatesFilesDefault = []string{
	"templates/files/main.go.tpl",
	"templates/files/weaver.toml.tpl",
}

var templateFilesBff = []string{
	"templates/files/router.go.tpl",
	"templates/files/bff.go.tpl",
}

var templateFilesComponent = []string{
	"templates/files/component.go.tpl",
	"templates/files/entity.go.tpl",
	"templates/files/sqlc.yaml.tpl",
	"templates/files/query.sql.tpl",
	"templates/files/schema.sql.tpl",
}
