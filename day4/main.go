package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func winningNumbers(line string) (int) {

	if len(line) == 0 {
		return 0
	}
	nums := strings.Split(line, "|")
	winning := nums[0]
	ours := nums[1]

	winningNumsString := strings.Split(winning," ")
	winningNums := make(map[int]bool)
	// fmt.Printf("Winning numbers: %v len(%v) \n", winningNumsString, len(winningNumsString))

	ourNumsString := strings.Split(ours," ")
	ourNums := []int{}
	// fmt.Printf("Our numbers: %v\n", ourNumsString)
	for _, n := range winningNumsString {
		nn, err := strconv.Atoi(n)
		if err == nil {
			winningNums[nn] = true
		}
	}
	// fmt.Printf("Winning numbers: %v\n", winningNums)


	for _, n := range ourNumsString {
		if len(n) > 0 {
			nn, err := strconv.Atoi(n)
			if err == nil {
				ourNums = append(ourNums, nn)
			}
		}
	}
	// fmt.Printf("Our numbers: %v\n", ourNums)

	wins := 0

	for _, n := range ourNums {
		if _, ok := winningNums[n]; ok {
			wins += 1
		}
	}

	return wins
}

func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	input := string(bytes)
	lines := strings.Split(input, "\n")
	part1(lines)
	part2(lines)
}


func part1(lines []string){
	sum := 0
	for _, line := range lines{
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, "|")
		winning := nums[0]
		ours := nums[1]

		winningNumsString := strings.Split(winning," ")
		winningNums := make(map[int]bool)
		// fmt.Printf("Winning numbers: %v len(%v) \n", winningNumsString, len(winningNumsString))

		ourNumsString := strings.Split(ours," ")
		ourNums := []int{}
		// fmt.Printf("Our numbers: %v\n", ourNumsString)
		for _, n := range winningNumsString {
			nn, err := strconv.Atoi(n)
			if err == nil {
				winningNums[nn] = true
			}
		}
		// fmt.Printf("Winning numbers: %v\n", winningNums)


		for _, n := range ourNumsString {
			if len(n) > 0 {
				nn, err := strconv.Atoi(n)
				if err == nil {
					ourNums = append(ourNums, nn)
				}
			}
		}
		// fmt.Printf("Our numbers: %v\n", ourNums)

		value := 0.0
		wins := 0

		for _, n := range ourNums {
			if _, ok := winningNums[n]; ok {
				// fmt.Printf("Winning number: %v\n", n)
				value += math.Pow(2.0, float64(wins))
				if wins > 0{
					value -= math.Pow(2.0, float64(wins-1))
				}
				wins += 1
			}
		}
		// fmt.Printf("Value: %v\n", value)
		sum += int(value)
	}

	fmt.Printf("Sum: %v\n", sum)
}

func part2(lines []string){
	copies := map[int]int{}
	copies[0] = 1

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		if _, ok := copies[i]; !ok {
			copies[i] = 1
		}
		wins := winningNumbers(line)
		// fmt.Printf("Winning numbers: %v\n", wins)
		for j := 1; j < wins+1 && i+j < len(lines); j ++ {
			if _, ok := copies[i+j]; !ok {
				copies[i+j] = 1
			}
			copies[i+j] += copies[i]
		}
	}
	result := 0

	for _,v := range copies {
		result += v
	}

	// fmt.Printf("Copies %v\n", copies)
	fmt.Printf("Sum %v\n", result)
}
