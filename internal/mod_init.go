package internal

import "fmt"

func ModInit() {
	moduleName, err := getModuleNameFromUserInput()
	if err != nil {
		panic(err)
	}

	args := []string{"mod", "init"}
	if moduleName != "" {
		args = append(args, moduleName)
	}

	err = runCmd(goCLI, args...)
	if err != nil {
		panic(err)
	}

}

func getModuleNameFromUserInput() (string, error) {
	print("What is the module name for your project? (e.g., github.com/renanbastos93/boneless) ")

	var moduleName string
	_, err := fmt.Scanf("%s\n", &moduleName)
	if err != nil {
		return "", err
	}

	return moduleName, nil
}
