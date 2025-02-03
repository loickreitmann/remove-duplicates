package main

import "fmt"

func main() {
	// Define input and output file names
	inputFile := "input.txt"
	outputFile := "output.txt"

	// Run the duplicate removal function
	if err := removeDuplicates(inputFile, outputFile); err != nil {
		fmt.Println("Error:", err)
	}
}
