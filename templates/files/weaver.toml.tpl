[single]
listeners.bff = {address = "localhost:8090"}

["{{.Module}}/internal/{{.ComponentName}}/Component"]
Driver = "mysql"
Source = "root:@tcp(localhost:3306)/mydatabase"