package lib

// Tetromino represents a 2D grid of strings, where '#' represents filled parts and '.' represents empty parts
type Tetromino [][]string

// canPlace checks if a tetromino can be placed at position (x, y) in the grid without overlap or out-of-bounds issues
func canPlace(grid [][]rune, tetromino Tetromino, x, y int) bool {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] == "#" && (x+i >= len(grid) || y+j >= len(grid[x]) || grid[x+i][y+j] != '.') {
				return false // Return false if placement is out of bounds or overlaps with another tetromino
			}
		}
	}
	return true
}

// placeTetromino places a tetromino at position (x, y) in the grid, marking it with a specific letter
func placeTetromino(grid [][]rune, tetromino Tetromino, x, y int, letter rune) {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] == "#" {
				grid[x+i][y+j] = letter // Fill grid cells with the specified letter
			}
		}
	}
}

// removeTetromino removes a tetromino from the grid by resetting its cells to '.'
func removeTetromino(grid [][]rune, tetromino Tetromino, x, y int) {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] == "#" {
				grid[x+i][y+j] = '.' // Reset grid cells to empty
			}
		}
	}
}

// trimTetrominoes removes empty rows and columns from tetrominoes to optimize their representation
func trimTetrominoes(tetrominoes []Tetromino) []Tetromino {
	var trimmedTetrominoes []Tetromino

	for _, tetromino := range tetrominoes {
		nonEmptylines := [][]string{}
		// Remove fully empty rows
		for _, line := range tetromino {
			emptyCount := 0
			for _, char := range line {
				if char == "." {
					emptyCount++
				}
			}
			if emptyCount != len(line) {
				nonEmptylines = append(nonEmptylines, line)
			}
		}

		// Skip tetrominoes with no meaningful content
		if len(nonEmptylines) == 0 {
			continue
		}

		// Determine the range of non-empty columns
		minCol, maxCol := len(nonEmptylines[0]), 0
		for _, line := range nonEmptylines {
			for colIndex, char := range line {
				if char == "#" {
					if colIndex < minCol {
						minCol = colIndex
					}
					if colIndex > maxCol {
						maxCol = colIndex
					}
				}
			}
		}

		// Trim columns to include only the meaningful content
		trimmedTetromino := [][]string{}
		for _, line := range nonEmptylines {
			trimmedline := line[minCol : maxCol+1]
			trimmedTetromino = append(trimmedTetromino, trimmedline)
		}

		trimmedTetrominoes = append(trimmedTetrominoes, trimmedTetromino)
	}

	return trimmedTetrominoes
}

// solve recursively attempts to fit all tetrominoes into the grid starting from the given index
func solve(grid [][]rune, tetrominoes []Tetromino, index int) ([][]rune, bool) {
	if index == len(tetrominoes) {
		return grid, true // All tetrominoes have been successfully placed
	}

	// Try placing the current tetromino at every position in the grid
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if canPlace(grid, tetrominoes[index], x, y) {
				placeTetromino(grid, tetrominoes[index], x, y, rune('A'+index)) // Place the tetromino

				// Recursively attempt to solve with the next tetromino
				solvedGrid, success := solve(grid, tetrominoes, index+1)
				if success {
					return solvedGrid, true // Return the solved grid if successful
				}

				// Backtrack: remove the tetromino if the placement doesn't lead to a solution
				removeTetromino(grid, tetrominoes[index], x, y)
			}
		}
	}

	return grid, false // No valid placement found, return failure
}
