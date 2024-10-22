package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(filePath string) ([][][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var tetrominoes [][][]string // The slice that will contain all tetrominoes
	var currentTetromino [][]string
	scanner := bufio.NewScanner(file)
	lineCount := 0
	consecutiveEmptyLines := 0 // Track consecutive empty lines

	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line is empty (newline)
		if strings.TrimSpace(line) == "" {
			// ERROR: MORE THAN ONE NEWLINE
			if consecutiveEmptyLines > 0 {
				return nil, fmt.Errorf("invalid format: too many newlines between tetrominoes")
			}
			consecutiveEmptyLines++ // Increment the counter for empty lines

			// If we have finished a tetromino block, append it to tetrominoes
			if len(currentTetromino) > 0 {
				// Check if tetromino has exactly 4 rows
				if len(currentTetromino) != 4 {
					return nil, fmt.Errorf("invalid format: tetromino does not have exactly 4 rows")
				}
				tetrominoes = append(tetrominoes, currentTetromino)
				currentTetromino = nil // Reset for the next tetromino
			}
			continue
		}

		// Reset consecutive empty lines counter since we encountered a non-empty line
		consecutiveEmptyLines = 0

		// Split the line into individual characters (row of the tetromino)
		row := strings.Split(line, "")

		// ERROR: MORE THAN 4 COLUMNS
		if len(row) != 4 {
			return nil, fmt.Errorf("invalid format: tetromino does not have exactly 4 columns")
		}

		currentTetromino = append(currentTetromino, row)
		lineCount++

		// ERROR: MORE THAN 4 ROWS
		if lineCount > 4 {
			return nil, fmt.Errorf("invalid format: tetromino does not have exactly 4 rows")
		}

		// After reading 4 lines, the next line must be an empty line
		if lineCount == 4 {
			if scanner.Scan() {
				nextLine := scanner.Text()
				// Check if the next line is a newline (empty)
				if strings.TrimSpace(nextLine) != "" {
					return nil, fmt.Errorf("invalid format - please modify the input")
				}
			} else {
				// End of file reached, reset for next tetromino
				break
			}
			tetrominoes = append(tetrominoes, currentTetromino)
			currentTetromino = nil // Reset for the next tetromino
			lineCount = 0
		}
	}

	// Add the last tetromino if it wasn't already added
	if len(currentTetromino) > 0 {
		// Check if the last tetromino has exactly 4 rows
		if len(currentTetromino) != 4 {
			return nil, fmt.Errorf("invalid format: last tetromino does not have exactly 4 rows")
		}
		tetrominoes = append(tetrominoes, currentTetromino)
	}

	// SCANNING ERROR
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return tetrominoes, nil
}

func TransformType(tetros [][][]string) []Tetromino {
	tetrominoes := []Tetromino{}
	for _, tetro := range tetros {
		tetrominoes = append(tetrominoes, tetro)
	}
	return tetrominoes
}

func CheckValidity(tetromino Tetromino) bool {
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
