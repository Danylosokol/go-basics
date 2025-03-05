package main

import (
	"fmt"
	"os"

	"frontendmasters.com/go/files/fileutils"
)

func main() {
	currentWorkingDirectory, _ := os.Getwd()
	filePath := currentWorkingDirectory + "/data/text.txt"
	contents, err := fileutils.ReadTextFile(filePath)
	if err == nil {
		fmt.Println(contents)
		newContent := fmt.Sprintf("Original: %s", contents)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("Error: %v", err)
	}
}
