[single]
listeners.bff = {address = "localhost:8090"}

["{{.Module}}/internal/{{.ComponentName}}/Component"]
{{if .IsSQL -}}
Driver = "mysql"
Source = "root:@tcp(localhost:3306)/mydatabase"
{{- else -}}
Driver = "sqlite3"
Source = "sqlite3.db"
{{- end}}
