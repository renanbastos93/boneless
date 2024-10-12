package internal

import (
	"fmt"
	"os"
)

func DeleteApp(appName string) {
	appFolderPath := getAppFolderPath(appName)
	checkIfAppFolderExists(appFolderPath)

	// Delete the app folder
	err := os.RemoveAll(appFolderPath)
	if err != nil {
		fmt.Println("Error deleting app folder: ", err)
		os.Exit(1)
	}

	fmt.Printf("App deleted successfully: %s\n", appName)
}

func checkIfAppFolderExists(pathToDelete string) {
	if _, err := os.Stat(pathToDelete); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("App folder not found")
		} else {
			fmt.Println("Error checking app folder: ", err)
		}
		os.Exit(1)
	}
}

func getAppFolderPath(appName string) string {
	return fmt.Sprintf("internal/%s", appName)
}
