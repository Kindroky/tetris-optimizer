package lib

import (
	"fmt"
)

// Define a palette of colors for displaying tetrominoes
var colors = []string{
	"\033[38;2;237;48;38m",   // Red
	"\033[38;2;245;126;34m",  // Orange
	"\033[38;2;235;211;61m",  // Yellow
	"\033[38;2;51;189;162m",  // Light Blue
	"\033[38;2;130;199;66m",  // Light Green
	"\033[38;2;255;255;255m", // White
	"\033[38;2;240;62;192m",  // Pinky Pink
	"\033[38;2;178;40;199m",  // Lavender
	"\033[38;2;17;125;69m",   // Turquoise
	"\033[38;2;14;72;99m",    // Dark Blue
	"\033[38;2;12;36;89m",    // Very Dark Blue
	"\033[38;2;82;34;163m",   // Purple
}

// FindSmallestSquare determines the smallest grid size that fits all tetrominoes
func FindSmallestSquare(tetrominoes []Tetromino) [][]rune {
	size := 2 // Start with a 2x2 grid (smallest possible size)

	for {
		board := createEmptyGrid(size)                // Create an empty grid of the current size
		trimmedtetros := trimTetrominoes(tetrominoes) // Trim unnecessary rows/columns from tetrominoes
		solvedGrid, success := solve(board, trimmedtetros, 0)
		if success {
			return solvedGrid // Return the solved grid if the tetrominoes fit
		}
		size++ // Increment the grid size if the current one is insufficient
	}
}

// createEmptyGrid generates an empty grid of the given size filled with dots
func createEmptyGrid(size int) [][]rune {
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
		for j := range grid[i] {
			grid[i][j] = '.' // Fill each cell with '.'
		}
	}
	return grid
}

// PrintColorfulGrid displays the solved grid with each tetromino in a unique color
func PrintColorfulGrid(grid [][]rune, tetrominoes []Tetromino) {
	gridSize := len(grid)
	identifierColors := map[rune]string{} // Map to associate tetromino identifiers with colors

	// Assign a color to each tetromino based on its identifier
	for i := range tetrominoes {
		identifierColors[rune('A'+i)] = colors[i%len(colors)]
	}

	// Print the grid with colored tetrominoes
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] >= 'A' && grid[i][j] <= 'Z' { // If cell contains a tetromino
				fmt.Print(identifierColors[grid[i][j]] + string(grid[i][j]) + "\033[0m") // Print with color
			} else {
				fmt.Print(".") // Print empty cells as '.'
			}
		}
		fmt.Println() // Move to the next row
	}
}
