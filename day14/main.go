package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	input := string(bytes)

	lines := strings.Fields(input)

	part1(lines)
	part2(lines)
}


func RotateLines(lines []string) []string {

	matrix := [][]rune{}

	for _, l := range lines{
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

	return newBlock
}

func rotateMatrix(lines []string) []string {
	matrix := [][]rune{}

	for _, l := range lines{
		runes := []rune(l)
		matrix = append(matrix, runes)
	}

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i<j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	newBlock := []string{}
	for i := range matrix {
		newBlock = append(newBlock, string(matrix[i]))
	}

	return newBlock
}

func part1(lines []string) {
	rotated := RotateLines(lines)
	// fmt.Println(PrettyFormat(rotated))
	tilted := []string{}

	for _, l := range rotated{
		movable := strings.Split(l, "#")
		sorted := []string{}
		for _, m := range movable {
			sortedM := SortString(m)
			sorted = append(sorted, sortedM)
		}

		moved := strings.Join(sorted, "#")
		fmt.Println(moved)
		tilted = append(tilted, moved)
	}

	total := 0
	for _, t := range tilted {
		for i, c := range t {
			if c == 'O'{
				total += i+1
			}
		}
	}

	fmt.Printf("Total %v\n", total)
}


func part2(lines []string) {
	seen := map[string]bool{}
	seenArr := []string{PrettyFormat(lines)}
	seen[PrettyFormat(lines)] = true
	modulo := 0
	cycleStart := 0
	total := 0

	for i := 0 ; i < 1000000000; i ++ {
		for j := 0 ; j < 4 ; j ++ {
			lines = rotateMatrix(lines)
			// fmt.Println(PrettyFormat(lines))
			tilted := []string{}

			for _, l := range lines{
				movable := strings.Split(l, "#")
				sorted := []string{}
				for _, m := range movable {
					sortedM := SortString(m)
					sorted = append(sorted, sortedM)
				}

				moved := strings.Join(sorted, "#")
				// fmt.Println(moved)
				tilted = append(tilted, moved)
			}

			lines = tilted
		}
		tiltedString := PrettyFormat(lines)

		if _, exists := seen[tiltedString]; exists {
			modulo = i+1
			cycleStart = slices.Index(seenArr, tiltedString)
			fmt.Printf("Modulo found %v\n", modulo)
			break
		}
		seen[tiltedString] = true
		seenArr = append(seenArr, tiltedString)
	}

	fmt.Printf("Modulo: %v cyclestart %v\n", modulo, cycleStart)

	jsonString :=  seenArr[((1000000000-cycleStart) % (modulo-cycleStart)) + cycleStart]
	// fmt.Println(jsonString)
	json.Unmarshal([]byte(jsonString), lines)
	// fmt.Println(lines)

	for _, t := range lines {
		for i, c := range t {
			if c == 'O'{
				total += len(lines)-i
			}
		}
	}

	fmt.Printf("P2 Total %v\n", total)

}
