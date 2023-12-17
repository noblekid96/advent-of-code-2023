package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
}

func Min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

func splitLinebreak(c rune) bool {
	return c == '\n'
}

func FindMirror(block []string) int {
	for i := 1; i < len(block); i ++ {
		above :=  append([]string{}, block[:i]...)
		below :=  append([]string{}, block[i:]...)

		// fmt.Printf("Above before reverse %v\n", PrettyFormat(above))
		slices.Reverse(above)
		// fmt.Printf("Above after reverse %v\n", PrettyFormat(above))

		min := Min(len(below), len(above))

		above = above[:min]
		below = below[:min]

		// fmt.Printf("I: %v\nAbove: %v\nBelow: %v\n\n", i, PrettyFormat(above), PrettyFormat(below))
		if PrettyFormat(above) == PrettyFormat(below){
			return i
		}
	}
	return 0
}


func FindMirror2(block []string) int {
	for i := 1; i < len(block); i ++ {
		above :=  append([]string{}, block[:i]...)
		below :=  append([]string{}, block[i:]...)

		// fmt.Printf("Above before reverse %v\n", PrettyFormat(above))
		slices.Reverse(above)
		// fmt.Printf("Above after reverse %v\n", PrettyFormat(above))

		min := Min(len(below), len(above))

		above = above[:min]
		below = below[:min]

		// fmt.Printf("I: %v\nAbove: %v\nBelow: %v\n\n", i, PrettyFormat(above), PrettyFormat(below))
		// if PrettyFormat(above) == PrettyFormat(below){
		// 	return i
		// }

		differences := 0
		for i := 0 ; i < len(above); i ++ {
			for j := 0; j < len(above[0]); j ++ {
				if above[i][j] != below[i][j]{
					differences ++
				}
			}
		}

		if differences == 1 {
			return i
		}
	}
	return 0
}


func RotatePattern(block []string) []string {

	matrix := [][]rune{}

	for _, l := range block{
		runes := []rune(l)
		matrix = append(matrix, runes)
	}


    rowCount := len(matrix)
    colCount := len(matrix[0])
    rotated := make([][]rune, colCount)

    for i := range rotated {
        rotated[i] = make([]rune, rowCount)
        for j := range matrix {
            rotated[i][j] = matrix[rowCount-1-j][i]
        }
    }

	newBlock := []string{}
	for i := range rotated {
		newBlock = append(newBlock, string(rotated[i]))
	}

	// fmt.Println(PrettyFormat(newBlock))
	return newBlock
}

func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	input := string(bytes)

	patterns := strings.Split(input, "\n\n")
	// fmt.Println(PrettyFormat(patterns))

	part1(patterns)
	part2(patterns)
}

func part1(patterns []string){
	total := 0

	for _, p := range patterns{
		block := strings.Fields(p)

		row := FindMirror(block)

		total += row * 100
		rotatedBlock := RotatePattern(block)

		column := FindMirror(rotatedBlock)
		total += column
	}

	fmt.Printf("Total: %v\n", total)
}

func part2(patterns []string){
	total := 0

	for _, p := range patterns{
		block := strings.Fields(p)

		row := FindMirror2(block)

		total += row * 100
		rotatedBlock := RotatePattern(block)

		column := FindMirror2(rotatedBlock)
		total += column
	}

	fmt.Printf("Total: %v\n", total)
}
