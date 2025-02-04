package main

import (
	"fmt"
	"sort"
)

// removeDuplicatesAndSort processes the input, sorts it, removes duplicates, and writes output files.
func removeDuplicatesAndSort(inputFile, sortedFile, outputFile string) error {
	// Read all lines from the file
	lines, err := readLines(inputFile)
	if err != nil {
		return err
	}

	// Sort the original list and write to sorted.txt
	sort.Strings(lines)
	if err := writeLines(sortedFile, lines); err != nil {
		return err
	}

	// Remove duplicates while keeping order
	uniqueMap := make(map[string]bool)
	var uniqueLines []string
	for _, line := range lines {
		if !uniqueMap[line] {
			uniqueMap[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}

	// Sort the deduplicated list and write to output.txt
	sort.Strings(uniqueLines)
	if err := writeLines(outputFile, uniqueLines); err != nil {
		return err
	}

	fmt.Printf("Processing complete!\n- Sorted original list -> %s\n- Sorted unique list -> %s\n", sortedFile, outputFile)
	return nil
}
