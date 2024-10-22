package lib

type Tetromino [][]string

func canPlace(grid [][]rune, tetromino Tetromino, x, y int) bool {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] == "#" && (x+i >= len(grid) || y+j >= len(grid[x]) || grid[x+i][y+j] != '.') {
				return false // Out of bounds or overlap with another tetromino
			}
		}
	}
	return true
}

func placeTetromino(grid [][]rune, tetromino Tetromino, x, y int, letter rune) {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] == "#" {
				grid[x+i][y+j] = letter
			}
		}
	}
}

func removeTetromino(grid [][]rune, tetromino Tetromino, x, y int) {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] == "#" {
				grid[x+i][y+j] = '.'
			}
		}
	}
}

func trimTetrominoes(tetrominoes []Tetromino) []Tetromino {
	var trimmedTetrominoes []Tetromino

	for _, tetromino := range tetrominoes {
		nonEmptylines := [][]string{}
		for _, line := range tetromino {
			emptyCount := 0
			for _, char := range line {
				if char == "." {
					emptyCount++
				}
			}
			if emptyCount != len(line) { // only add lines that aren't full dots
				nonEmptylines = append(nonEmptylines, line)
			}
		}

		// remove empty columns
		if len(nonEmptylines) == 0 {
			continue
		}

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

		trimmedTetromino := [][]string{}
		for _, line := range nonEmptylines {
			trimmedline := line[minCol : maxCol+1]
			trimmedTetromino = append(trimmedTetromino, trimmedline)
		}

		trimmedTetrominoes = append(trimmedTetrominoes, trimmedTetromino)
	}

	return trimmedTetrominoes
}

func solve(grid [][]rune, tetrominoes []Tetromino, index int) ([][]rune, bool) {
	if index == len(tetrominoes) {
		return grid, true // Solution found, return the grid
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if canPlace(grid, tetrominoes[index], x, y) {
				placeTetromino(grid, tetrominoes[index], x, y, rune('A'+index))

				solvedGrid, success := solve(grid, tetrominoes, index+1)
				if success {
					return solvedGrid, true // Return the solved grid when successful
				}

				// Backtrack: remove the tetromino if placing it here doesn't lead to a solution
				removeTetromino(grid, tetrominoes[index], x, y)
			}
		}
	}
	return grid, false // No valid placement found, return the grid and false
}
