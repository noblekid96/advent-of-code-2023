package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Comparison struct {
	key string
	cmp string
	n int
	target string
}

type Instruction struct {
	rules []Comparison
	fallback string
}

func splitFn (c rune) bool {
	return c == '\n'
}

func splitCurlyFn (c rune) bool {
	return c == '{'
}

func splitCommaFn (c rune) bool {
	return c == ','
}

func splitEqualFn (c rune) bool {
	return c == '='
}

func splitColonFn (c rune) bool {
	return c == ':'
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}


func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	input := string(bytes)

	blocks := strings.Split(input, "\n\n")
	part1(blocks)
	part2(blocks)
}

func operation(key int, cmp string, n int ) bool {
	if cmp == "<" {
		if key < n {
			return true
		}
		return false
	} else if cmp == ">" {
		if key > n {
			return true
		}
		return false
	}

	panic("Invalid operation")
}

func acceptable(item map[string]int, name string, workflows map[string]Instruction) bool {
	if name == "R"{
		return false
	}
	if name == "A" {
		return true
	}

	// fmt.Printf("Workflows %v\n", workflows)
	// fmt.Printf("Name %v\n", name)
	if _, exists := workflows[name]; !exists {
		panic(fmt.Sprintf("Name [%v] does not exist in workflows: %v\n", name, workflows))
	}
	instruction := workflows[name]
	rules := instruction.rules
	fallback := instruction.fallback

	// fmt.Printf("Instruction %v\n", instruction)
	// fmt.Printf("Rules %v\n", rules)

	for _, r := range rules {

		result := operation(item[r.key], r.cmp, r.n)
		if result {
			return acceptable(item, r.target, workflows)
		}
	}

	return acceptable(item, fallback, workflows)
}

func part1(blocks []string ){
	block1 := blocks[0]
	block2 := blocks[1]
	workflow := map[string]Instruction{}

	for _, l := range strings.FieldsFunc(block1, splitFn){
		lineParts := strings.FieldsFunc(l, splitCurlyFn)
		name := lineParts[0]
		rest := lineParts[1][:len(lineParts[1])-1]

		rules := strings.FieldsFunc(rest, splitCommaFn)
		workflow[name] = Instruction{[]Comparison{}, rules[len(rules)-1]}
		rules = rules[:len(rules)-1]
		// fmt.Println(rules)

		for _, r := range rules {
			rParts := strings.FieldsFunc(r, splitColonFn)
			// fmt.Println(rParts)
			comparison := rParts[0]
			target := rParts[1]
			key := string(comparison[0])
			cmp := string(comparison[1])
			n, err := strconv.Atoi(comparison[2:])
			if err != nil {
				panic(err)
			}

			if entry, ok := workflow[name]; ok {
				entry.rules = append(entry.rules, Comparison{key, cmp, n, target})

				workflow[name] = entry
			}
		}

	}
	// fmt.Println("Workflow", workflow)

	total := 0

	for _, line := range strings.FieldsFunc(block2, splitFn){
		// fmt.Println(line)

		item := map[string]int{}

		for _, segment := range strings.FieldsFunc(line[1:len(line)-1], splitCommaFn){
			segmentParts := strings.FieldsFunc(segment, splitEqualFn)
			ch := segmentParts[0]
			nStr := segmentParts[1]

			n, err := strconv.Atoi(nStr)

			if err != nil {
				panic(err)
			}

			item[ch] = n
		}

		// fmt.Println("Item", item)

		if acceptable(item, "in", workflow){
			for _, v := range item {
				total += v
			}
		}
	}

	fmt.Println("P1 total", total)
}

func count(ranges map[string][]int, name string, workflows map[string]Instruction) int64 {
	// fmt.Println("Ranges", ranges)
	// fmt.Println("Name", name)
	// fmt.Println()
	if name == "R"{
		return 0
	}
	if name == "A" {
		var product int64
		product = 1

		for _, r := range ranges {
			lo, hi := r[0], r[1]
			product *= int64(hi) - int64(lo) + 1
		}
		return product
	}

	var total int64
	total = 0
	workflow := workflows[name]
	rules := workflow.rules
	fallback := workflow.fallback
	// useFallback := false

	for _, rule := range rules {
		ruleRange := ranges[rule.key]
		lo, hi := ruleRange[0], ruleRange[1]
		cmp := rule.cmp
		target := rule.target
		key := rule.key
		n := rule.n
		var T []int
		var F []int
		if cmp == "<" {
			T = []int{lo, Min(n-1,hi)}
			F = []int{Max(n, lo), hi}
		} else {
			T = []int{Max(n+1, lo), hi}
			F = []int{lo, Min(n, hi)}
		}

		if T[0] <= T[1] {
			rangeCopy := make(map[string][]int)
			for k, v := range ranges {
				rangeCopy[k] = v
			}
			rangeCopy[key] = T
			total += count(rangeCopy, target, workflows)
		}

		if F[0] <= F[1] {
			rangeCopy := make(map[string][]int)
			for k, v := range ranges {
				rangeCopy[k] = v
			}
			rangeCopy[key] = F
			ranges = rangeCopy
		} else {
			break
		}
	}
	total += count(ranges, fallback, workflows)

	return total
}

func part2(blocks []string){
	block1 := blocks[0]
	// block2 := blocks[1]
	workflow := map[string]Instruction{}

	for _, l := range strings.FieldsFunc(block1, splitFn){
		lineParts := strings.FieldsFunc(l, splitCurlyFn)
		name := lineParts[0]
		rest := lineParts[1][:len(lineParts[1])-1]

		rules := strings.FieldsFunc(rest, splitCommaFn)

		workflow[name] = Instruction{[]Comparison{}, rules[len(rules)-1]}
		rules = rules[:len(rules)-1]
		// fmt.Println(rules)

		for _, r := range rules {
			rParts := strings.FieldsFunc(r, splitColonFn)
			// fmt.Println(rParts)
			comparison := rParts[0]
			target := rParts[1]
			key := string(comparison[0])
			cmp := string(comparison[1])
			n, err := strconv.Atoi(comparison[2:])
			if err != nil {
				panic(err)
			}

			if entry, ok := workflow[name]; ok {
				entry.rules = append(entry.rules, Comparison{key, cmp, n, target})

				workflow[name] = entry
			}
		}

	}
	// fmt.Println("Workflow", workflow)

	ranges := map[string][]int{}
	word := "xmas"

	for i := 0; i < len(word); i ++ {
		ranges[string(word[i])] = []int{1,4000}
	}

	total := count(ranges, "in", workflow)


	fmt.Println("P2 total", total)
}
