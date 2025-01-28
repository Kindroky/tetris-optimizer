package main

import (
	"fmt"
	"os"
	"time"

	lib "tetris-optimizer/lib"
)

func main() {
	timer := time.Now() // Start a timer to measure execution time

	// Validate argument count
	if len(os.Args) != 2 {
		fmt.Println("Invalid argument entry. \nUsage: go run . <textfile.txt>")
		os.Exit(1)
	}

	// Get the file path from the command-line argument
	filePath := os.Args[1]

	// Read tetrominoes data from the input file
	tetrominoes, err := lib.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Ensure the input file contains tetrominoes
	if len(tetrominoes) == 0 {
		fmt.Println("Error: No tetrominoes found")
		os.Exit(1)
	}

	// Validate each tetromino's format and structure
	for _, tetromino := range tetrominoes {
		if !lib.CheckValidity(lib.Tetromino(tetromino)) {
			fmt.Println("Error: Invalid tetromino found in input")
			os.Exit(1)
		}
	}

	// Transform tetromino data into the desired internal structure
	realtetros := lib.TransformType(tetrominoes)

	// Solve the puzzle using the smallest square possible
	fmt.Println("Solving the puzzle...")
	solvedGrid := lib.FindSmallestSquare(realtetros)

	// Announce and print the solved grid
	finalAnnouncement := fmt.Sprintf("Solved board in %s :", time.Since(timer))
	fmt.Println(finalAnnouncement)
	lib.PrintColorfulGrid(solvedGrid, realtetros)

	// Count and display the number of empty spaces (dots) in the solved grid
	nbDots := findNbofDots(solvedGrid)
	DotsMessage := ""
	if nbDots == 1 {
		DotsMessage = "Found 1 empty space in the grid!"
	} else {
		DotsMessage = fmt.Sprintf("Found %d empty spaces in the grid!", nbDots)
	}
	fmt.Println(DotsMessage)
}

// findNbofDots counts the number of empty spaces (dots) in the solved grid
func findNbofDots(solvedgrid [][]rune) int {
	count := 0
	for _, line := range solvedgrid {
		for _, char := range line {
			if char == '.' {
				count++
			}
		}
	}
	return count
}
