package infra

import (
	"fmt"
	"log"
	"os"
)

// LogError logs an error with detailed context
func LogError(err error, context string) {
	if err != nil {
		logMessage := fmt.Sprintf("ERROR: %s - %v", context, err)
		log.Println(logMessage)
		// Optionally, write to a file
		file, _ := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		defer file.Close()
		file.WriteString(logMessage + "\n")
	}
}
