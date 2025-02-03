package main

import (
	"bufio"
	"fmt"
	"os"
)

// removeDuplicates reads a file, removes duplicate lines, and writes unique lines to a new file.
func removeDuplicates(inputFile, outputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	// Using a map to track unique strings
	uniqueLines := make(map[string]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		uniqueLines[line] = true
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Writing unique lines to the output file
	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for line := range uniqueLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	writer.Flush()

	fmt.Printf("Successfully removed duplicates. Output saved to %s\n", outputFile)
	return nil
}
