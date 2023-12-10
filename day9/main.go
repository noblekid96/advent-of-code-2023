package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func splitLinebreak(c rune) bool {
	return c == '\n'
}

func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
}

func AllZeroes(s []int) bool {
	for _, i := range s {
		if i != 0 {
			return false
		}
	}
	return true
}

func ExtrapolatePart1(h []int) int {
	current := h
	histories := [][]int{}
	histories = append(histories, current)

	// fmt.Printf("Current %v\n", PrettyFormat(current))
	for AllZeroes(current) != true {
		next := []int{}

		for i := 1 ; i < len(current) ; i ++ {
			next = append(next, current[i] - current[i-1])
		}
		current = next
		histories = append(histories, current)
		// fmt.Printf("Current %v\n", PrettyFormat(current))
	}

	// fmt.Printf("histories %v\n", PrettyFormat(histories))

	for i := len(histories)-1; i >= 0; i -- {
		if i == len(histories) - 1 {
			histories[i] = append(histories[i], 0)
			continue
		}
		histories[i] = append(histories[i], histories[i+1][len(histories[i+1])-1] + histories[i][len(histories[i])-1]  )
	}

	// fmt.Printf("histories after extrapolate %v\n", PrettyFormat(histories))

	return histories[0][len(histories[0])-1]
}

func ExtrapolatePart2(h []int) int {
	current := h
	histories := [][]int{}
	histories = append(histories, current)

	// fmt.Printf("Current %v\n", PrettyFormat(current))
	for AllZeroes(current) != true {
		next := []int{}

		for i := 1 ; i < len(current) ; i ++ {
			next = append(next, current[i] - current[i-1])
		}
		current = next
		histories = append(histories, current)
		// fmt.Printf("Current %v\n", PrettyFormat(current))
	}

	// fmt.Printf("histories %v\n", PrettyFormat(histories))

	for i := len(histories)-1; i >= 0; i -- {
		if i == len(histories) - 1 {
			histories[i] = append([]int{0},histories[i][0:]...)
			continue
		}
		// histories[i] = append(histories[i], histories[i+1][len(histories[i+1])-1] + histories[i][len(histories[i])-1])
		histories[i] = append([]int{histories[i][0] - histories[i+1][0]}, histories[i][0:]...)
	}

	// fmt.Printf("histories after extrapolate %v\n", PrettyFormat(histories))

	return histories[0][0]
}

func part1(lines []string){
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		start := []int{}
		nums := strings.Fields(line)
		for _ , ns := range nums {
			n, _ := strconv.Atoi(ns)
			start = append(start, n)
		}
		extrapolated := ExtrapolatePart1(start)
		// fmt.Printf("Extrapolated %v\n", extrapolated)
		total += extrapolated
	}
	fmt.Println("Part1 total ", total)
}


func part2(lines []string){
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		start := []int{}
		nums := strings.Fields(line)
		for _ , ns := range nums {
			n, _ := strconv.Atoi(ns)
			start = append(start, n)
		}
		extrapolated := ExtrapolatePart2(start)
		// fmt.Printf("Extrapolated %v\n", extrapolated)
		total += extrapolated
	}
	fmt.Println("Part2 total ", total)
}
