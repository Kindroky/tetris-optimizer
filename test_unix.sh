#!/bin/bash

# List of test files
declare -a test_files=("test1.txt" "test2.txt" "test3.txt" "test4.txt" "test5.txt" "test6.txt" "test7.txt" "test8.txt" "test9.txt" "test10.txt" "test11.txt")

# Loop through each test file
for i in "${!test_files[@]}"; do
    test_file="${test_files[$i]}"

    echo "Running test with $test_file"

    # Run the Go program with the current test file and capture the output
    result=$(go run main.go "$test_file" 2>&1) # Capture both stdout and stderr

    # Display the result or any error message
    echo "Result for $test_file:"
    echo "$result"
    echo "----------------------------"
done

