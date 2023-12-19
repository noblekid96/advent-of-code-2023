package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/elliotchance/orderedmap/v2"
)


func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
}



func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	input := string(bytes)

	part1(input)
	part2(input)
}

func hash(s string) int {
	v := 0

	for _, c := range s {
		v += int(c)
		v *= 17
		v %= 256
	}

	return v
}

func part1(input string){
	inputs := strings.Split(strings.TrimSpace(input), ",")
	fmt.Println(inputs)

	total := 0
	for _, i := range inputs {
		total += hash(i)
	}

	fmt.Println("P1 total: ", total)
}

func part2(input string){
	inputs := strings.Split(strings.TrimSpace(input), ",")
	fmt.Println(inputs)

	total := 0
	hashmap := make([]*orderedmap.OrderedMap[string, int], 256)
	for _, i := range inputs {

		last := i[len(i)-1]
		if last == '-'{
			label := i[:len(i)-1]
			hash := hash(label)
			box := hashmap[hash]
			if box == nil {
				hashmap[hash] = orderedmap.NewOrderedMap[string, int]()
				box = hashmap[hash]
			}
			if _, exists := box.Get(label); exists {
				box.Delete(label)
			}
			continue
		}

		label := i[:len(i)-2]
		hash := hash(label)
		box := hashmap[hash]
		if box == nil {
			hashmap[hash] = orderedmap.NewOrderedMap[string, int]()
			box = hashmap[hash]
		}
		focus, err := strconv.Atoi(string(i[len(i)-1]))
		if err != nil {
			panic(err)
		}

		box.Set(label, focus)
	}

	for i, m := range hashmap {
		if m != nil {
			for j, key := range m.Keys() {
				value, _:= m.Get(key)
				// fmt.Println(i, key, value, j)

				total += (i+1)*(j+1)*value
			}
		}

	}

	// fmt.Printf("hashmap %v\n", PrettyFormat(hashmap))
	fmt.Println("P2 total: ", total)
}
