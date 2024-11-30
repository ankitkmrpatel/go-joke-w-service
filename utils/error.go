// HandleError logs the error with context and terminates the application if it's critical
package utils

import (
	"log"
	"os"
)

// HandleError is a centralized error handler that logs and terminates the program if needed
func HandleError(err error, context string, critical bool) {
	if err != nil {
		log.Printf("Error occurred while %s: %v", context, err)
		if critical {
			log.Fatal("Shutting down application due to critical error.")
			os.Exit(1)
		}
	}
}
