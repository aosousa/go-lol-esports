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
