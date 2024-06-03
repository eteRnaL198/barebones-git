package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func CreateBlob(srcPath, destPath string) error {
	sourceFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	fileName := filepath.Base(srcPath)
	scanner := bufio.NewScanner(sourceFile)
	writer := bufio.NewWriter(destinationFile)

	if _, err := writer.WriteString(fmt.Sprintf("blob %s\n", fileName)); err != nil {
		return err
	}

	for scanner.Scan() {
		line := scanner.Text()
		if _, err := writer.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// func CreateTree() error {

// }
