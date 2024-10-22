package internal

import (
	"fmt"
	"os"
)

func DeleteApp(appName string) {
	if isValid := isAppNameValid(appName); !isValid {
		return
	}

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

func isAppNameValid(appName string) bool {
	if appName == "app" {
		fmt.Println("You can't delete the app folder, it's the example component")
		return false
	}

	if appName == "bff" {
		fmt.Println("You can't delete the bff folder, it's required to run the application")
		return false
	}

	return true
}

func checkIfAppFolderExists(pathToDelete string) {
	if _, err := os.Stat(pathToDelete); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("App folder not found")
		} else {
			fmt.Println("Error checking app folder: ", err)
		}
		panic(0)
	}
}

func getAppFolderPath(appName string) string {
	return fmt.Sprintf("internal/%s", appName)
}
