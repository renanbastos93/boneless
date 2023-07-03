package internal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml/v2"
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

func ReadToml(componentName string) (qsConn string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// TODO: improve that same internal/run.go:46
	weaverToml, err := ioutil.ReadFile(pwd + "/weaver.toml")
	if err != nil {
		panic(err)
	}

	var cfgWeaver map[string]interface{}
	err = toml.Unmarshal(weaverToml, &cfgWeaver)
	if err != nil {
		panic(err)
	}

	projectName := GetProjectName(pwd)
	componentModPath := projectName + "/internal/" + componentName + "/Component"
	for k, v := range cfgWeaver {
		if k != componentModPath {
			continue
		}
		if value, ok := v.(map[string]interface{}); ok {
			var driver = ""
			var source = ""
			for nk, nv := range value {
				switch nk {
				case "Driver":
					driver = nv.(string)
				case "Source":
					source = nv.(string)
				}
			}
			qsConn = fmt.Sprintf("%s://%s", driver, source)
		}
	}

	if qsConn == "" {
		panic("not found settings to running migrate verify your weaver.toml file")
	}

	return qsConn
}
