package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("AoC Day 2")
	// Open the text file
	inputFile, err := os.Open("Day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	var tot_score = 0
	// split by line
	lineParser := bufio.NewScanner(inputFile)
	// for each line in file
	for lineParser.Scan() {
		var my_val = 0
		var their_val = 0
		var cur_line = lineParser.Text()
		var my_rps = strings.Split(cur_line, " ")
		var i_need = my_rps[1]
		var they_play = my_rps[0]
		// instead of I play, its it should end as
		if i_need == "X" {
			//lose
			my_val = 0
		}
		if i_need == "Y" {
			//Tie
			my_val = 3
		}
		if i_need == "Z" {
			//Win
			my_val = 6
		}
		if they_play == "A" {
			//rock
			their_val = 1
		}
		if they_play == "B" {
			//papper
			their_val = 2
		}
		if they_play == "C" {
			//scissors
			their_val = 3
		}
		// get your play
		if my_val == 6 {
			tot_score += my_val
			if they_play == "A" {
				// add paper
				tot_score += 2
			} else if they_play == "B" {
				// add scissors
				tot_score += 3
			} else {
				tot_score += 1
			}
		} else if my_val == 0 {
			if they_play == "A" {
				// add
				tot_score += 3
			} else if they_play == "B" {
				tot_score += 1
			} else {
				tot_score += 2
			}
		} else {
			tot_score += my_val
			tot_score += their_val
		}

	}
	fmt.Println(" Total: ", tot_score)
}

/*
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func main() {
	fmt.Println("AoC Day 2")
	// Open the text file
	inputFile, err := os.Open("Day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	var tot_score = 0
	// split by line
	lineParser := bufio.NewScanner(inputFile)
	// for each line in file
	for lineParser.Scan() {
		var my_val = 0
		var their_val = 0
		var cur_line = lineParser.Text()
		var my_rps = strings.Split(cur_line, " ")
		var i_play = my_rps[1]
		var they_play = my_rps[0]
		if i_play == "X" {
			//rock
			my_val = 1
		}
		if i_play == "Y" {
			//paper
			my_val = 2
		}
		if i_play == "Z" {
			//scissors
			my_val = 3
		}
		if they_play == "A" {
			//rock
			their_val = 1
		}
		if they_play == "B" {
			//papper
			their_val = 2
		}
		if they_play == "C" {
			//scissors
			their_val = 3
		}
		// you win conditions
		if (i_play == "X" && they_play == "C") || (i_play == "Y" && they_play == "A") || (i_play == "Z" && they_play == "B") {
			tot_score += my_val
			tot_score += 6
		} else if my_val == their_val {
			tot_score += my_val
			tot_score += 3
		} else {
			tot_score += my_val
		}
	}
	fmt.Println(" Total: ", tot_score)
}
*/
