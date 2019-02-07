package utils

import (
	"fmt"
	"os"
)

// HandleError prints out an error that occurred and exits the application
func HandleError(err error) {
	fmt.Printf("%s", err)
	os.Exit(1)
}

// FindInSlice detects if a string is in a slice. Used to remove matches from ignored leagues from today's matches results
func FindInSlice(str string, list []string) bool {
	for _, el := range list {
		if str == el {
			return true
		}
	}

	return false
}
