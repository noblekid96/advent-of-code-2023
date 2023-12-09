package main

import (
	"fmt"
	"math"
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

func part1(lines []string) {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	output := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			break
		}
		possible := true
		data := strings.Split(line, ": ")

		sets := strings.Split(data[1], "; ")
	inner:
		for j := 0; j < len(sets); j++ {
			set := sets[j]
			num_color := strings.Split(set, ", ")
			for k := 0; k < len(num_color); k++ {
				num_and_color := strings.Split(num_color[k], " ")
				num, err := strconv.Atoi(num_and_color[0])
				if err != nil {
					panic(fmt.Sprintf("Failed to convert %s to integer", num_and_color[0]))
				}
				color := num_and_color[1]
				if num > limits[color] {
					possible = false
					break inner
				}
			}
		}
		if possible {
			output += i + 1
		}
	}
	fmt.Printf("part 1: %v\n", output)
}

func part2(lines []string) {
	output := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			break
		}
		counts := []float64{0, 0, 0}
		data := strings.Split(line, ": ")

		sets := strings.Split(data[1], "; ")
		for j := 0; j < len(sets); j++ {
			set := sets[j]
			num_color := strings.Split(set, ", ")
			for k := 0; k < len(num_color); k++ {
				num_and_color := strings.Split(num_color[k], " ")
				num, err := strconv.Atoi(num_and_color[0])
				num_f := float64(num)
				if err != nil {
					panic(fmt.Sprintf("Failed to convert %s to integer", num_and_color[0]))
				}
				color := num_and_color[1]
				switch color {
				case "red":
					counts[0] = math.Max(counts[0], num_f)
				case "green":
					counts[1] = math.Max(counts[1], num_f)
				case "blue":
					counts[2] = math.Max(counts[2], num_f)
				}
			}
		}
		power := counts[0] * counts[1] * counts[2]
		output += int(power)
	}
	fmt.Printf("part 2: %v\n", output)
}
