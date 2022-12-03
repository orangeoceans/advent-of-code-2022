package shared

import (
	"bufio"
	"log"
	"os"

	"golang.org/x/exp/constraints"
)

// Number is a custom type set of constraints extending the Float and Integer type set from the experimental constraints package.
type Number interface {
	constraints.Float | constraints.Integer
}

func SumSlice[T Number](slice []T) T {
	// Helper function to sum a slice of numbers
	var sum T = 0
	for _, num := range slice {
		sum += num
	}
	return sum
}

func InsertIntoSlice[T any](slice []T, item T, index int) []T {
	// Helper function to insert item of arbitrary type into slice at index.
	return append(slice[:index], append([]T{item}, slice[index:]...)...)
}

func LinesFromFile(file_path string) []string {
	// Helper function to get the lines in a file as a slice.
	f, err := os.Open(file_path)
	if err != nil { // This is necessary because Go doesn't do exceptions, which seems horrible.
		log.Fatal(err)
	}
	// Close file on function return.
	defer f.Close()

	// Read each line in the file with scanner.
	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}

	return lines
}
