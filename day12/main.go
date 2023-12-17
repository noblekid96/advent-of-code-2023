package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
}

func splitLinebreak(c rune) bool {
	return c == '\n'
}

func splitDot(c rune) bool {
	return c == '.'
}

type CombinationArg struct {
	record string
	index int
}

type CountArg struct {
	record string
	expected []int
}

// Key returns a string representation of CountArg that can be used as a map key.
func (c CountArg) Key() string {
	var sb strings.Builder
	sb.WriteString(c.record)
	sb.WriteRune(':')
	for _, v := range c.expected {
		sb.WriteString(fmt.Sprintf("%d,", v))
	}
	return sb.String()
}

var memo = map[CombinationArg][]string{}
var countMemo = map[string]int{}

func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	input := string(bytes)

	lines := strings.FieldsFunc(input, splitLinebreak)

	part1(lines)
	part2(lines)
}

func Combinations(record string, expected []int) int {
	combos := CombinationsHelper(record, 0)
	// fmt.Printf("Combos: %v\n", PrettyFormat(combos))
	total := 0

	for _, combo := range combos {
		valid := true
		springs := strings.FieldsFunc(combo, splitDot)
		if len(springs) != len(expected){
			continue
		}

		for i, s := range springs {
			if len(s) != expected[i]{
				valid = false
				break
			}
		}

		if (valid){
			total += 1
		}
	}
	return total
}

func CombinationsHelper(record string, i int) []string {
	combiArg := CombinationArg{record, i}
	if _, exists := memo[combiArg]; exists {
		return memo[combiArg]
	}

	combinations := []string{}
	if i == len(record){
		combinations = append(combinations, record)
		// memo[combiArg] = combinations
		return combinations
	}
	if record[i] == '?'{
		recordRunes := []rune(record)
		recordRunes[i] = '.'
		newRecord := string(recordRunes)
		combinations = append(combinations,CombinationsHelper(newRecord, i+1)...)

		recordRunes[i] = '#'
		newRecord = string(recordRunes)
		combinations = append(combinations,CombinationsHelper(newRecord, i+1)...)
	} else {
		combinations = append(combinations, CombinationsHelper(record, i+1)...)
	}
	memo[combiArg] = combinations
	return combinations
}

func Count(record string, expected []int) int {
	if record == ""{
		if len(expected) == 0{
			return 1
		} else {
			return 0
		}
	}

	if len(expected) == 0 {
		if strings.Contains(record, "#"){
			return 0
		} else {
			return 1
		}
	}

	countArg := CountArg{record, expected}
	countArgKey := countArg.Key()
	if _, exists := countMemo[countArgKey]; exists {
		return countMemo[countArgKey]
	}
	result := 0

	if record[0] == '.' || record[0] == '?'{
		result += Count(record[1:], expected)
	}

	if record[0] == '#' || record[0] == '?'{
		if expected[0] <= len(record) && !strings.Contains(record[:expected[0]],".") &&
			(expected[0] == len(record) || record[expected[0]] != '#'){
				newRecord := ""
				if len(record) > expected[0]+1{
					newRecord = record[expected[0]+1:]
				} else {
					newRecord = ""
				}
				result += Count(newRecord, expected[1:])
			}
	}

	countMemo[countArgKey] = result
	return result
}

func part1(lines []string){
	records := []string{}
	expected := [][]int{}

	for _, l := range lines {
		// fmt.Println(l)
		fields := strings.Fields(l)
		// fmt.Printf("Fields: %v\n", PrettyFormat(fields))

		records = append(records, fields[0])
		e := []int{}
		for _, s := range strings.Split(fields[1],","){
			i, _ := strconv.Atoi(s)
			e = append(e, i)
		}
		expected = append(expected, e)
	}
	// fmt.Printf("Records: %v Expected: %v\n", PrettyFormat(records), PrettyFormat(expected))

	total := 0
	for i := 0; i < len(records); i ++ {
		count := Count(records[i], expected[i])
		// fmt.Printf("Count : %v\n", count)
		total += count
	}
	fmt.Printf("Total: %v\n", total)
}

func part2(lines []string){
	records := []string{}
	expected := [][]int{}

	for _, l := range lines {
		// fmt.Println(l)
		fields := strings.Fields(l)
		// fmt.Printf("Fields: %v\n", PrettyFormat(fields))

		newFields := []string{}
		for i :=0 ; i < 5; i ++ {
			newFields = append(newFields, fields[0])
		}
		fields[0] = strings.Join(newFields,"?")

		records = append(records, fields[0])
		e := []int{}
		for _, s := range strings.Split(fields[1],","){
			i, _ := strconv.Atoi(s)
			e = append(e, i)
		}
		nums := []int{}
		for i := 0 ; i < 5; i ++ {
			nums = append(nums, e...)
		}

		expected = append(expected, nums)
	}
	// fmt.Printf("Records: %v Expected: %v\n", PrettyFormat(records), PrettyFormat(expected))

	total := 0
	for i := 0; i < len(records); i ++ {
		count := Count(records[i], expected[i])
		// fmt.Printf("Count : %v\n", count)
		total += count
	}
	fmt.Printf("Total: %v\n", total)
}
