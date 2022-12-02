package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 1")
	// Open the text file
	inputFile, err := os.Open("Day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	var cur_count = 0
	var cal_sums []int
	var great1_value = 0
	var great2_value = 0
	var great3_value = 0
	// split by line
	lineParser := bufio.NewScanner(inputFile)
	// for each line in file
	for lineParser.Scan() {
		// get text
		var cur_line = lineParser.Text()
		//convert to int
		cur_num, err := strconv.Atoi(cur_line)
		// if not an int, then I got a new line
		if err != nil {
			if cur_count >= great3_value {
				great3_value = cur_count
			}
			if great3_value > great2_value {
				great3_value, great2_value = great2_value, great3_value
			}
			if great2_value > great1_value {
				great2_value, great1_value = great1_value, great2_value
			}
			cal_sums = append(cal_sums, cur_count)
			cur_count = 0
		}
		cur_count += cur_num
	}
	fmt.Println(" Cal1: ", great1_value, "Cal2: ", great2_value, "Cal3: ", great3_value)
	var total = great1_value + great2_value + great3_value
	fmt.Println("Sum: ", total)
}

/*package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 1")
	// Open the text file
	inputFile, err := os.Open("Day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	var cur_count = 0
	var cal_sums []int
	var greatest_index = 0
	var cur_index = 0
	var greatest_value = 0
	// split by line
	lineParser := bufio.NewScanner(inputFile)
	// for each line in file
	for lineParser.Scan() {
		// get text
		var cur_line = lineParser.Text()
		//convert to int
		cur_num, err := strconv.Atoi(cur_line)
		// if not an int, then I got a new line
		if err != nil {
			cur_index++
			if cur_count >= greatest_value {
				greatest_value = cur_count
				greatest_index = cur_index
			}
			cal_sums = append(cal_sums, cur_count)
			cur_count = 0
		}
		cur_count += cur_num
	}
	fmt.Println("Index: ", greatest_index, " Calories: ", greatest_value)
}*/
