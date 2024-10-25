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
	if exists := folderExists(appFolderPath); !exists {
		return
	}

	// Delete the app folder
	if err := os.RemoveAll(appFolderPath); err != nil {
		panic("Error deleting app folder: " + err.Error())
	}

	fmt.Printf("App deleted successfully: %s\n", appName)
}

func isAppNameValid(appName string) bool {
	if appName == "app" {
		fmt.Println("You can't delete the app folder: it's the example component")
		return false
	}

	if appName == "bff" {
		fmt.Println("You can't delete the bff folder: it's required to run the application")
		return false
	}

	return true
}

func folderExists(pathToDelete string) bool {
	if _, err := os.Stat(pathToDelete); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("App folder not found")
			return false
		}

		panic("Error checking app folder: " + err.Error())
	}
	return true
}

func getAppFolderPath(appName string) string {
	return fmt.Sprintf("internal/%s", appName)
}
