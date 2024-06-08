["{{.Module}}/internal/{{.ComponentName}}/Component"]
{{if .IsSQL -}}
Driver = "mysql"
Source = "root:@tcp(localhost:3306)/mydatabase"
{{- else -}}
Driver = "sqlite3"
Source = "db.sqlite3"
{{- end}}
