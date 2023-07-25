package cli

// Type is a new types inheritance from int to define iota for we use it
type (
	DBType       int
	TemplateType int
)

type TemplateFiles struct {
	Pattern  string
	Filepath string
}

// Template type to define constants using iota and string values
const (
	AllType TemplateType = 1 + iota
	BaseType
	BFFType
	ComponentType

	All       = "all"
	Base      = "base"
	BFF       = "bff"
	Component = "component"
)

// Database type to define constants using iota and string values
const (
	MySQLType DBType = 1 + iota
	SQLType
	SQLite3Type

	MySQL   = "mysql"
	SQL     = "sql"
	SQLite3 = "sqlite3"
)

// TemplateTypeStr convert TemplateType to string
var TemplateTypeStr = map[TemplateType]string{
	AllType:       All,
	BaseType:      Base,
	BFFType:       BFF,
	ComponentType: Component,
}

// StrToTemplateType convert string to TemplateType
var StrToTemplateType = map[string]TemplateType{
	All:       AllType,
	Base:      BaseType,
	BFF:       BFFType,
	Component: ComponentType,
}

// DBTypeToStr convert DBType to string
var DBTypeToStr = map[DBType]string{
	MySQLType:   MySQL,
	SQLType:     SQL,
	SQLite3Type: SQLite3,
}

// StrToDBType convert string to DBType
var StrToDBType = map[string]DBType{
	MySQL:   MySQLType,
	SQL:     SQLType,
	SQLite3: SQLite3Type,
}

// BaseFiles is a list of all template files for use
var BaseFiles = []TemplateFiles{
	{
		Pattern:  "files/main.go.tpl",
		Filepath: "/cmd/main.go",
	},
	{
		Pattern:  "files/weaver.toml.tpl",
		Filepath: "/weaver.toml",
	},
}

// BFFFiles is a list of all template files to BFF
var BFFFiles = []TemplateFiles{
	{
		Pattern:  "files/bff.go.tpl",
		Filepath: "/internal/bff/bff.go",
	},
	{
		Pattern:  "files/router.go.tpl",
		Filepath: "/internal/bff/router.go",
	},
}

// ComponentFiles is a list of all template files to Component/app
var ComponentFiles = []TemplateFiles{
	{
		Pattern:  "files/component.go.tpl",
		Filepath: "/internal/%s/component.go",
	},
	{
		Pattern:  "files/entity.go.tpl",
		Filepath: "/internal/%s/entity.go",
	},
	{
		Pattern:  "files/create_table.up.sql.tpl",
		Filepath: "/internal/%s/db/migrations/000001_create_table.up.sql",
	},
	{
		Pattern:  "files/create_table.down.sql.tpl",
		Filepath: "/internal/%s/db/migrations/000001_create_table.down.sql",
	},
	{
		Pattern:  "files/query.sql.tpl",
		Filepath: "/internal/%s/db/query.sql",
	},
	{
		Pattern:  "files/sqlc.yaml.tpl",
		Filepath: "/internal/%s/db/sqlc.yaml",
	},
}

// TemplateTypeToFilePath is a map to filter easier which template we'd should be generate
var TemplateTypeToFilePath = map[TemplateType][]TemplateFiles{
	BaseType:      BaseFiles,
	BFFType:       BFFFiles,
	ComponentType: ComponentFiles,
	AllType:       append(BaseFiles, append(BFFFiles, ComponentFiles...)...),
}
