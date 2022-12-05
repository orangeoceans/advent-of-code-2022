package main

import (
	"advent_2022/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func stringPairToInts(slice []string) []int {
	// Helper function to convert a pair of strings to a pair of ints.
	out_slice := make([]int, 2)
	for i := range out_slice {
		item_as_int, err := strconv.Atoi(slice[i])
		if err != nil {
			log.Fatal(err)
		}
		out_slice[i] = item_as_int
	}
	return out_slice
}

func checkFullyInside(range_1 []int, range_2 []int) bool {
	// Helper function to check if either range is fully inside the other.
	if (range_1[0] <= range_2[0] && range_1[1] >= range_2[1]) ||
		(range_2[0] <= range_1[0] && range_2[1] >= range_1[1]) {
		return true
	}
	return false
}

func checkOverlap(range_1 []int, range_2 []int) bool {
	// Helper function to check if the given ranges overlap.
	if (range_1[0] <= range_2[0] && range_1[1] >= range_2[0]) ||
		(range_2[0] <= range_1[0] && range_2[1] >= range_1[0]) {
		return true
	}
	return false
}

func Day4Part1() {
	num_fully_contained := 0
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		range_1 := stringPairToInts(strings.Split(assignments[0], "-"))
		range_2 := stringPairToInts(strings.Split(assignments[1], "-"))
		if checkFullyInside(range_1, range_2) {
			num_fully_contained += 1
		}
	}
	fmt.Printf("Number of fully contained pairs: %d\n", num_fully_contained)
}

func Day4Part2() {
	num_overlapping := 0
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		range_1 := stringPairToInts(strings.Split(assignments[0], "-"))
		range_2 := stringPairToInts(strings.Split(assignments[1], "-"))
		if checkOverlap(range_1, range_2) {
			num_overlapping += 1
		}
	}
	fmt.Printf("Number of overlapping pairs: %d\n", num_overlapping)
}

var lines = shared.LinesFromFile("day4_input.txt")

func main() {
	fmt.Println("Day 4 Part 1 result:")
	Day4Part1()
	fmt.Println("Day 4 Part 2 result:")
	Day4Part2()
}
