package main

import "github.com/renanbastos93/boneless/internal"

func main() {
	// boneless sqlc generate
	// boneless build start
	// internal.RunSqlcGenerate("app")
	// internal.BuildStartProject("zig")
	internal.BuildComponent("audit")
}
