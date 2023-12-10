package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	// lines := strings.Split(input, "\n")
	part1(input)
	// part2(input)
}

func splitLinebreak(c rune) bool {
	return c == '\n'
}

func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
}


func part1(input string){
	blocks := strings.Split(input, "\n\n")
	directions := blocks[0]
	networkBlock := strings.FieldsFunc(blocks[1], splitLinebreak)
	network := map[string][]string{}
	start := "AAA"

	for _, n := range networkBlock {
		fields := strings.Fields(n)
		src := fields[0]
		leftDst := strings.Trim(fields[2],"(,")
		rightDst := strings.Trim(fields[3],")")

		network[src] = []string{leftDst,rightDst}

	}
	fmt.Printf("Directions: %v \n\n Network: %v\n", directions, PrettyFormat(network))

	steps := 0
	current := start
	fmt.Println("Current", current)

	for current != "ZZZ" {
		if directions[0] == 'L'{
			current = network[current][0]
		} else {
			current = network[current][1]
		}
		directions = directions[1:] + string(directions[0])
		// fmt.Printf("Current %v directionsLeft %v", current, directions[i:])
		steps += 1
	}

	fmt.Println("Number of steps taken to reach ZZZ", steps)
}


func endsWith(loc string, r string) bool{
	if loc[len(loc)-1:] != string(r) {
		return false
	}
	return true
}


func allendsWith(locs []string, r string) bool{
	for _, l := range locs{
		if endsWith(l, r) == false{
			return false
		}
	}
	return true
}

func part2(input string){
	blocks := strings.Split(input, "\n\n")
	directions := blocks[0]
	networkBlock := strings.FieldsFunc(blocks[1], splitLinebreak)
	network := map[string][]string{}
	start := []string{}

	for _, n := range networkBlock {
		fields := strings.Fields(n)
		src := fields[0]
		leftDst := strings.Trim(fields[2],"(,")
		rightDst := strings.Trim(fields[3],")")

		network[src] = []string{leftDst,rightDst}
		if endsWith(src, "A"){
			start = append(start, src)
		}
	}
	fmt.Printf("Directions: %v \n\n Network: %v\n", directions, PrettyFormat(network))


	steps := 0
	current := start
	fmt.Println("Current", current)

	for allendsWith(current,"ZZZ") == false {
		for i, c := range current {
			if directions[0] == 'L'{
				current[i] = network[c][0]
			} else {
				current[i] = network[c][1]
			}
		}
		directions = directions[1:] + string(directions[0])
		steps += 1
	}

	fmt.Println("Number of steps taken to reach ZZZ", steps)
}
