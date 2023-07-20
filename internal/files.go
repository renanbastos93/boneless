package internal

type tplFiles struct {
	pattern  string
	filePath string
	kind     string
}

var (
	KindBase      = "base"
	KindBff       = "bff"
	KindComponent = "component"
	KindAll       = "all"
)

var templateFiles = []tplFiles{
	{
		pattern:  "files/main.go.tpl",
		filePath: "/cmd/main.go",
		kind:     KindBase,
	},
	{
		pattern:  "files/weaver.toml.tpl",
		filePath: "/weaver.toml",
		kind:     KindBase,
	},
	{
		pattern:  "files/bff.go.tpl",
		filePath: "/internal/bff/bff.go",
		kind:     KindBff,
	},
	{
		pattern:  "files/router.go.tpl",
		filePath: "/internal/bff/router.go",
		kind:     KindBff,
	},
	{
		pattern:  "files/component.go.tpl",
		filePath: "/internal/%s/component.go",
		kind:     KindComponent,
	},
	{
		pattern:  "files/entity.go.tpl",
		filePath: "/internal/%s/entity.go",
		kind:     KindComponent,
	},
	{
		pattern:  "files/create_table.up.sql.tpl",
		filePath: "/internal/%s/db/migrations/000001_create_table.up.sql",
		kind:     KindComponent,
	},
	{
		pattern:  "files/create_table.down.sql.tpl",
		filePath: "/internal/%s/db/migrations/000001_create_table.down.sql",
		kind:     KindComponent,
	},
	{
		pattern:  "files/query.sql.tpl",
		filePath: "/internal/%s/db/query.sql",
		kind:     KindComponent,
	},
	{
		pattern:  "files/sqlc.yaml.tpl",
		filePath: "/internal/%s/db/sqlc.yaml",
		kind:     KindComponent,
	},
}
