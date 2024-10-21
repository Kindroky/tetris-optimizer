package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tetromino [][]string

var colors = []string{
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
	"\033[37m", // White
}

func main() {
	// Validate argument count
	if len(os.Args) != 2 {
		fmt.Println("Invalid argument entry. \nUsage: go run . <textfile.txt>")
		os.Exit(1)
	}

	// Get the file path from the argument
	filePath := os.Args[1]

	// Read tetrominoes from the file (assuming ReadFile takes a string path)
	tetrominoes, err := ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(tetrominoes) == 0 {
		fmt.Println("Error: No tetrominoes found")
		os.Exit(1)
	}

	// Validate each tetromino and print true or false
	for _, tetromino := range tetrominoes {
		valid := checkValidity(Tetromino(tetromino))
		fmt.Println(valid)
	}
}

func createEmptyGrid(size int) [][]rune {
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

// printPrettyPrincessGrid prints each tetromino in a different color.
func printColorfulGrid(grid [][]rune, tetrominoes []Tetromino) {
	gridSize := len(grid)
	identifierColors := map[rune]string{}

	for i := range tetrominoes {
		identifierColors[rune('A'+i)] = colors[i%len(colors)]
	}

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] >= 'A' && grid[i][j] <= 'Z' {
				fmt.Print(identifierColors[grid[i][j]] + string(grid[i][j]) + "\033[0m")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func ReadFile(filePath string) ([][][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	var tetrominoes [][][]string // The slice that will contain all tetrominoes
	var currentTetromino [][]string
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// If the line is empty, that means we finished a tetromino block
		if strings.TrimSpace(line) == "" {
			if len(currentTetromino) > 0 {
				tetrominoes = append(tetrominoes, currentTetromino)
				currentTetromino = nil // Reset for the next tetromino
			}
			lineCount = 0
			continue
		}

		// Split the line into individual characters
		row := strings.Split(line, "")
		currentTetromino = append(currentTetromino, row)
		lineCount++

		// If 4 lines (for example, for a 4x4 grid tetromino) have been read, start a new tetromino
		if lineCount == 4 {
			tetrominoes = append(tetrominoes, currentTetromino)
			currentTetromino = nil // Reset for the next tetromino
			lineCount = 0
		}
	}

	// Add the last tetromino if it wasn't already added
	if len(currentTetromino) > 0 {
		tetrominoes = append(tetrominoes, currentTetromino)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	return tetrominoes, nil
}

func checkValidity(tetromino Tetromino) bool {
	hashcount := 0
	adjacencyCount := 0
	content := ""

	// Flatten the tetromino grid into a single string
	for _, lines := range tetromino {
		for _, char := range lines {
			content += string(char)
		}
	}

	// Track the total number of '#' and check their adjacency
	for i := 0; i < len(content); i++ {
		if content[i] == '#' {
			hashcount++

			// Check right adjacency (ignore rightmost boundary)
			if (i%4) != 3 && content[i+1] == '#' {
				adjacencyCount++
			}

			// Check down adjacency (ignore bottom boundary)
			if i+4 < len(content) && content[i+4] == '#' {
				adjacencyCount++
			}
		}
	}

	return hashcount == 4 && adjacencyCount >= 3
}
