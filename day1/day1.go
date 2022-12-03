package main

import (
	"advent_2022/shared"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

func updateMaxCaloriesSlice(elf_calories int, max_calories []int) []int {
	if elf_calories <= max_calories[0] {
		return max_calories
	}
	for i, max_cal := range max_calories {
		if elf_calories <= max_cal {
			max_calories = shared.InsertIntoSlice(max_calories, elf_calories, i)
			return max_calories[1:]
		}
	}
	return append(max_calories, elf_calories)[1:]
}

func maxCaloriesElvesV1(lines []string, num_elves int) []int {
	// O(N*M), where N = # lines, M = # elves wanted
	// Faster with low M
	max_calories := make([]int, num_elves)
	elf_calories := 0
	for _, line := range lines {
		if line == "" {
			max_calories = updateMaxCaloriesSlice(elf_calories, max_calories)
			elf_calories = 0
			continue
		}
		num_calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		elf_calories += num_calories
	}
	return max_calories
}

func maxCaloriesElvesV2(lines []string, num_elves int) []int {
	// O(NlogN), where N = # lines
	// Faster with high M = # of elves wanted
	var elves []int
	elf_calories := 0
	for _, line := range lines {
		if line == "" {
			elves = append(elves, elf_calories)
			elf_calories = 0
			continue
		}
		num_calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		elf_calories += num_calories
	}
	sort.Ints(elves)
	return elves[len(elves)-num_elves:]
}

func Day1Part1() {
	max_calories := maxCaloriesElvesV1(lines, 1)
	fmt.Printf("\tMax number of calories: %d\n", max_calories[0])
}

func Day2Part2V1() {
	max_calories_sum := shared.SumSlice(maxCaloriesElvesV1(lines, 3))
	fmt.Printf("\tMax number of calories: %d\n", max_calories_sum)
}

func Day2Part2V2() {
	max_calories_sum := shared.SumSlice(maxCaloriesElvesV2(lines, 3))
	fmt.Printf("\tMax number of calories: %d\n", max_calories_sum)
}

var lines = shared.LinesFromFile("day1_input.txt")

func main() {
	fmt.Println("Day 1 Part 1 results:")
	Day1Part1()
	fmt.Println("Day 1 Part 2 v1 results:")
	start_time := time.Now()
	Day2Part2V1()
	finish_v1_time := time.Now()
	fmt.Println("Time to run: ", finish_v1_time.Sub(start_time))
	fmt.Println("Day 1 Part 2 v2 results:")
	Day2Part2V2()
	finish_v2_time := time.Now()
	fmt.Println("Time to run: ", finish_v2_time.Sub(finish_v1_time))
}
