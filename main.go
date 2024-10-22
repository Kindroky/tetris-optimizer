package main

import (
	"fmt"
	"os"
	"time"

	lib "tetris-optimizer/lib"
)

func main() {
	timer := time.Now()
	// Validate argument count
	if len(os.Args) != 2 {
		fmt.Println("Invalid argument entry. \nUsage: go run . <textfile.txt>")
		os.Exit(1)
	}

	// Get the file path from the argument
	filePath := os.Args[1]

	// Read tetrominoes from the file
	tetrominoes, err := lib.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(tetrominoes) == 0 {
		fmt.Println("Error: No tetrominoes found")
		os.Exit(1)
	}

	// Validate each tetromino
	for _, tetromino := range tetrominoes {
		if !lib.CheckValidity(lib.Tetromino(tetromino)) {
			fmt.Println("Error: Invalid tetromino found in input")
			os.Exit(1)
		}
	}

	realtetros := lib.TransformType(tetrominoes)

	// Solve the board with the smallest square possible
	fmt.Println("Solving the puzzle...")
	solvedGrid := lib.FindSmallestSquare(realtetros)

	// Print the result
	finalAnnouncement := fmt.Sprintf("Solved board in %s :", time.Since(timer))
	fmt.Println(finalAnnouncement)
	lib.PrintColorfulGrid(solvedGrid, realtetros)
	nbDots := findNbofDots(solvedGrid)
	DotsMessage := ""
	if nbDots == 1 {
		DotsMessage = "Found 1 empty space in the grid!"
	} else {
		DotsMessage = fmt.Sprintf("Found %d empty spaces in the grid!", nbDots)
	}
	fmt.Println(DotsMessage)
}

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
