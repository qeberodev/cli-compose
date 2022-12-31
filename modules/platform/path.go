package platform

import (
	"log"
	"os"
)

const WIN32 = "windows"

// Get path for users config directory
func GetConfigDir() string {
	path, err := os.UserConfigDir()

	if err != nil {
		log.Fatalf("Unable to read current users config directory %v", err)
	}

	return path
}

// Get path for users home directory
func GetHomeDir() string {
	path, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("Unable to read current users home directory %v", err)
	}

	return path
}
