# Tetris Optimizer - Go Project

## Project Overview

This project is a Golang-based program that reads a list of tetrominoes from a text file and attempts to assemble them into the smallest possible square grid. Each tetromino is represented by a unique uppercase letter. The program uses algorithms to efficiently place the tetrominoes, aiming to minimize the grid size. In cases where a perfect square is not possible, the program allows spaces between tetrominoes.

## Features

- Assembles tetrominoes into the smallest square grid possible.
- Identifies each tetromino with a unique uppercase Latin letter (A for the first, B for the second, etc.).
- Handles errors gracefully by printing the corresponding error message for invalid input files or malformed tetrominoes.
- Supports at least one tetromino per input file.
- Outputs results with any remaining spaces in the grid marked with periods (`.`).

## Bonus features added

-  Provides detailed performance statistics, including solving time and the number of empty spaces in the final grid.
- Decorates all tetrominoes in different colors for clarity.

## Project Structure

The program expects a single argument: a path to a text file that contains tetrominoes in a specific format. It reads the file, processes the tetrominoes, and outputs the smallest grid to the terminal.

### Example Input File

```txt
....#
....#
....#
####

....
....
..##
..##

....
.###
....
....

....
##..
.##.
....
```

## Example Output

```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

Example output with all additional features (except colored tetrominoes, not supported in the readme) : 
```
Solving the puzzle...
Solved board in 91.0874ms :
AABKKK.
AAB.KGG
DDBFFFG
DDBCCFG
.ECCJII
EEHHJJI
EHH..JI
Found 5 empty spaces in the grid!
```




