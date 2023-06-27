package internal

import (
	"fmt"
	"os"
)

func GetProjectName(dir string) (moduleName string) {
	modFile, err := os.Open(dir + "/go.mod")
	if err != nil {
		panic(fmt.Sprintf("go.mod does not exist: %v", err))
	}

	defer modFile.Close()
	_, err = fmt.Fscanf(modFile, "module %s", &moduleName)
	if err != nil {
		panic(fmt.Sprintf("read go mod error: %v", err))
	}

	return moduleName
}
