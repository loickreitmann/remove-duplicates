package main

import (
	"bufio"
	"fmt"
	"os"
)

// writeLines writes a slice of strings to a file, each string on a new line.
func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	writer.Flush()
	return nil
}
