package internal

import (
	"fmt"
	"os"
)

func DeleteApp(appName string) {
	validateAppName(appName)

	appFolderPath := getAppFolderPath(appName)
	checkIfAppFolderExists(appFolderPath)

	// Delete the app folder
	err := os.RemoveAll(appFolderPath)
	if err != nil {
		fmt.Println("Error deleting app folder: ", err)
		panic(err)
	}

	fmt.Printf("App deleted successfully: %s\n", appName)
}

func validateAppName(appName string) {
	if appName == "app" {
		fmt.Println("You can't delete the app folder, it's the example component")
		os.Exit(0)
	}

	if appName == "bff" {
		fmt.Println("You can't delete the bff folder, it's required to run the application")
		os.Exit(0)
	}
}

func checkIfAppFolderExists(pathToDelete string) {
	if _, err := os.Stat(pathToDelete); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("App folder not found")
		} else {
			fmt.Println("Error checking app folder: ", err)
		}
		panic(err)
	}
}

func getAppFolderPath(appName string) string {
	return fmt.Sprintf("internal/%s", appName)
}
