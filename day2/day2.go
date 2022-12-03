package main

import (
	"advent_2022/shared"
	"fmt"
	"strings"
)

var WIN_SCORE = 6
var TIE_SCORE = 3
var LOSE_SCORE = 0
var ROCK_SCORE = 1
var PAPER_SCORE = 2
var SCISSORS_SCORE = 3
var hand_score_map = map[string]int{"X": ROCK_SCORE, "Y": PAPER_SCORE, "Z": SCISSORS_SCORE}
var outcome_score_map = map[string]int{"X": LOSE_SCORE, "Y": TIE_SCORE, "Z": WIN_SCORE}

// My solution uses hashmaps to match letter codes.
// There's an elegant way of doing this that uses mod to exploit the ring nature of the game.
// However, it's a lot less readable, and not really more efficient, so I used hashmaps.
var tie_map = map[string]string{"A": "X", "B": "Y", "C": "Z"}
var win_map = map[string]string{"A": "Y", "B": "Z", "C": "X"}
var rock_map = map[string]string{"A": "Y", "B": "X", "C": "Z"}  // Which outcomes require a rock
var paper_map = map[string]string{"A": "Z", "B": "Y", "C": "X"} // Which outcomes require a paper

func getHandsFromLine(line string) (string, string) {
	// Helper function to unpack the two letter codes from each line.
	split_string := strings.Fields(line)
	return split_string[0], split_string[1]
}

func Day2Part1() {
	total_score := 0
	for _, line := range lines {
		other_hand, own_hand := getHandsFromLine(line)
		total_score += hand_score_map[own_hand]
		if own_hand == win_map[other_hand] { // Win condition
			total_score += WIN_SCORE
		} else if own_hand == tie_map[other_hand] { // Tie condition
			total_score += TIE_SCORE
		} else { // Lose condition
			total_score += LOSE_SCORE
		}
	}
	fmt.Printf("Total score: %d\n", total_score)
}

func Day2Part2() {
	total_score := 0
	for _, line := range lines {
		other_hand, own_outcome := getHandsFromLine(line)
		total_score += outcome_score_map[own_outcome]
		if own_outcome == rock_map[other_hand] { // Give a rock
			total_score += ROCK_SCORE
		} else if own_outcome == paper_map[other_hand] { // Give a paper
			total_score += PAPER_SCORE
		} else { // Give a scissors
			total_score += SCISSORS_SCORE
		}
	}
	fmt.Printf("Total score: %d\n", total_score)
}

var lines = shared.LinesFromFile("day2_input.txt")

func main() {
	fmt.Println("Day 2 Part 1 results:")
	Day2Part1()
	fmt.Println("Day 2 Part 2 results:")
	Day2Part2()
}
