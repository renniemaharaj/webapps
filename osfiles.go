package webapps

import (
	"bufio"
	"fmt"
	"os"
)

// Function to write content to a text file
func WriteToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	//Remove this for productions
	fmt.Println("Content written to", filename)
	return nil
}

// Function to read content from a text file
func ReadFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

// FileExists checks if a file exists at the given path
func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
