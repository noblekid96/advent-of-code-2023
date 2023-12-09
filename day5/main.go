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
	// lines := strings.Split(input, "\n")
	part1(input)
	part2(input)
}

func PrettyPrint(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      fmt.Println(string(s))
      return string(s)
}

func splitLinebreak(c rune) bool {
	return c == '\n'
}

func part1(input string){
	slice := strings.Split(input,"\n\n")
    // PrettyPrint(slice)
	mappings := [][][]int{}
	seedStrings := strings.Fields(slice[0])[1:]
	seeds := []int{}
	for _, s := range seedStrings {
		n, err := strconv.Atoi(s)
		if err == nil {
			seeds = append(seeds,n)
		}
	}
    // PrettyPrint(seeds)

	for _, mapping := range slice[1:] {
		mapEntries := [][]int{}
		mapEntriesString := strings.FieldsFunc(mapping, splitLinebreak)[1:]
		for _, e := range mapEntriesString {
			m := strings.Fields(e)
			entries := []int{}
			for _, ns := range m {
				n, err := strconv.Atoi(ns)
				if err == nil {
					entries = append(entries, n)
				}
			}
			mapEntries = append(mapEntries, entries)
		}
		mappings = append(mappings, mapEntries)
	}

	// PrettyPrint(mappings)
	// PrettyPrint(mappings[0])

	smallest := 9999999999

	for _, seed := range seeds {
		searchFor := seed


		for _, m := range mappings {
			// fmt.Printf("Finding %v in mapping %v\n", searchFor, i)
			for _, e := range m {
				// fmt.Printf("Finding %v in mapping entry %v\n", searchFor, j)
				if searchFor >= e[1] && (searchFor < e[1] + e[2]){
					// fmt.Printf("%v matches map entry %v\n", searchFor, e)
					searchFor = e[0] + (searchFor-e[1])
					break
				}
			}
		}
		// fmt.Printf("Location: %v\n", searchFor)
		if searchFor < smallest {
			smallest = searchFor
		}
	}
	fmt.Println("Smallest location: ", smallest)

}

func part2(input string){
	slice := strings.Split(input,"\n\n")
    // PrettyPrint(slice)
	mappings := [][][]int{}
	seedStrings := strings.Fields(slice[0])[1:]
	seeds := []int{}
	for _, s := range seedStrings {
		n, err := strconv.Atoi(s)
		if err == nil {
			seeds = append(seeds,n)
		}
	}
    // PrettyPrint(seeds)

	for _, mapping := range slice[1:] {
		mapEntries := [][]int{}
		mapEntriesString := strings.FieldsFunc(mapping, splitLinebreak)[1:]
		for _, e := range mapEntriesString {
			m := strings.Fields(e)
			entries := []int{}
			for _, ns := range m {
				n, err := strconv.Atoi(ns)
				if err == nil {
					entries = append(entries, n)
				}
			}
			mapEntries = append(mapEntries, entries)
		}
		mappings = append(mappings, mapEntries)
	}

	// PrettyPrint(mappings)
	// PrettyPrint(mappings[0])

	smallest := 9999999999

	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedEnd := seeds[i+1]
		for j := seedStart; j < seedStart + seedEnd; j ++ {
			searchFor := j


			for _, m := range mappings {
				// fmt.Printf("Finding %v in mapping %v\n", searchFor, i)
				for _, e := range m {
					// fmt.Printf("Finding %v in mapping entry %v\n", searchFor, j)
					if searchFor >= e[1] && (searchFor < e[1] + e[2]){
						// fmt.Printf("%v matches map entry %v\n", searchFor, e)
						searchFor = e[0] + (searchFor-e[1])
						break
					}
				}
			}
			// fmt.Printf("Location: %v\n", searchFor)
			if searchFor < smallest {
				smallest = searchFor
			}
		}
	}
	fmt.Println("Smallest location: ", smallest)
}
