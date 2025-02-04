package main

import "fmt"

func main() {
	// Define input, sorted, and output file names
	inputFile := "input.txt"
	sortedFile := "sorted.txt"
	outputFile := "output.txt"

	// Run the duplicate removal and sorting function
	if err := removeDuplicatesAndSort(inputFile, sortedFile, outputFile); err != nil {
		fmt.Println("Error:", err)
	}
}
