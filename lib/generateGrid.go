package lib

import (
	"fmt"
)

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

func FindSmallestSquare(tetrominoes []Tetromino) [][]rune {
	size := 2 // Start with the smallest possible square size (2x2)

	for {
		board := createEmptyGrid(size)
		trimmedtetros := trimTetrominoes(tetrominoes)
		solvedGrid, success := solve(board, trimmedtetros, 0)
		if success {
			return solvedGrid
		}
		size++ // Increment the size of the board and try again if the current size can't fit the tetrominoes
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
func PrintColorfulGrid(grid [][]rune, tetrominoes []Tetromino) {
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
