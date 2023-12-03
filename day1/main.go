package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	example := string(bytes)
	lines := strings.Split(example,"\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	output := 0
	for i := 0; i < len(lines); i++ {
		line := []rune(lines[i])
		first := 0
		last := 0
		for j := 0; j < len(line); j ++ {
			if unicode.IsDigit(line[j]){
				first = int(line[j] - '0')
				break
			}
		}
		for j := len(line)-1; j > -1; j -- {
			if unicode.IsDigit(line[j]){
				last = int(line[j] - '0')
				break
			}
		}
		output += first*10+last
	}
	fmt.Printf("Part 1 Output: %v\n", output)
}

func part2(lines []string) {
	output := 0
	regex := "one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9"
	regex_r := reverse(regex)
	regex1, err := regexp.Compile(regex)
	if err != nil {
		panic("regex failed to compile")
	}
	regex2, err := regexp.Compile(regex_r)
	if err != nil {
		panic("regex failed to compile")
	}

	mapping := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}


	for i := 0; i < len(lines); i++ {
		line := lines[i]

		first := mapping[regex1.FindString(line)]
		last := mapping[reverse(regex2.FindString(reverse(line)))]
		output += first*10+last
	}
	fmt.Printf("Part 2 Output: %v\n", output)
}

func reverse(s string) string {
    rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }

    // return the reversed string.
    return string(rns)
}
