package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func createAndAppendObject(filesName string) *os.File {
	// Create New File log and close it immediately
	file, err := os.Create(filesName)
	if err != nil {
		log.Println("Error creating file:", err)
		return nil
	}
	file.Close()

	// Open the created file in append mode and save to struct
	file, err = os.OpenFile(filesName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return nil
	}

	return file
}

func fileNameGenerator(objectName string) string {
	// Generate file name based on gateName
	today := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s:%s.txt", today, objectName)
}
