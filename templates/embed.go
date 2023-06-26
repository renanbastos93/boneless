package templates

import "embed"

//go:embed files/*.tpl
var CreateTemplateFS embed.FS
