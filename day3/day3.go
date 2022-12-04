package main

import (
	"advent_2022/shared"
	"fmt"
	"unicode"

	"golang.org/x/exp/maps"
)

func letterToNumber(letter rune) int {
	// Helper function to convert a letter into the desired number:
	// a-z -> 1-26
	// A-Z -> 27-52
	if unicode.IsLower(letter) {
		return int(letter - 'a' + 1)
	} else {
		return int(letter - 'A' + 27)
	}
}

func intersectInPlace[T comparable, U any](base_set map[T]U, other_set map[T]U) {
	// Helper function that removes non-shared keys in the base_set with other_set.
	for key := range base_set {
		if _, in_other_set := other_set[key]; !in_other_set {
			delete(base_set, key)
		}
	}
}

func Day3Part1() {
	priorities_sum := 0
	for _, line := range lines {
		line_runes := []rune(line)             // Split string into slice of runes
		bag_1_contents := make(map[rune]bool)  // Hashset of the first compartment
		bag_2_start_idx := len(line_runes) / 2 // Each compartment has equal # items
		for i, item := range line_runes {
			if i < bag_2_start_idx {
				bag_1_contents[item] = true // Store item into bag 1 hashset
				continue
			}
			if _, in_bag_1 := bag_1_contents[item]; in_bag_1 { // If item in bag 1 hashset
				priorities_sum += letterToNumber(item)
				break
			}
		}
	}
	fmt.Printf("Sum of priorities of shared items: %d\n", priorities_sum)
}

func Day3Part2() {
	priorities_sum := 0
	common_items := make(map[rune]bool)
	for l, line := range lines {
		line_runes := []rune(line)          // Split string into slice of runes
		bag_contents := make(map[rune]bool) // Hashset of seen letters
		for _, item := range line_runes {
			bag_contents[item] = true
		}
		if l%3 == 0 {
			common_items = bag_contents
		} else {
			intersectInPlace(common_items, bag_contents) // Update set of common items
			if l%3 == 2 {                                // Third elf, we know only one remaining common item
				badge := maps.Keys(common_items)[0]
				priorities_sum += letterToNumber(badge)
			}
		}
	}
	fmt.Printf("Sum of priorities of shared items: %d\n", priorities_sum)
}

var lines = shared.LinesFromFile("day3_input.txt")

func main() {
	fmt.Println("Day 3 Part 1 results:")
	Day3Part1()
	fmt.Println("Day 3 Part 2 results:")
	Day3Part2()
}
