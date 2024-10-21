package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
