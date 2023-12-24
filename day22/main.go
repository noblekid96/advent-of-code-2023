package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func splitCommaFn (c rune) bool {
	return c == ','
}

type Point struct {
	X int
	Y int
	Z int
}

func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
}


func Max(i , j int) int {
	if i > j {
		return i
	}
	return j
}

func Min(i , j int) int {
	if i < j {
		return i
	}
	return j
}

func RemoveFromSlice(a, b []int) []int {
	newSlice := []int{}

	for _, i := range a {
		if slices.Contains(b, i){
			continue
		}
		newSlice = append(newSlice, i)
	}
	return newSlice
}

func Subset(a, b []int) bool {
	for _, i := range b {
		if slices.Contains(a, i){
			continue
		} else {
			return false
		}
	}
	return true
}


func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	input := string(bytes)

	splitFn := func(c rune) bool {
		return c == '\n'
	}

	input = strings.ReplaceAll(input, "~", ",")
	lines := strings.FieldsFunc(input, splitFn)
	part1(lines)
}

func overlaps(a, b []int) bool {
	return Max(a[0],b[0]) <= Min(a[3],b[3]) && Max(a[1], b[1]) <= Min(a[4], b[4])
}

func part1(lines []string){
	brickCoords := [][]int{}

	for _, l := range lines {
		coordStr := strings.FieldsFunc(l, splitCommaFn)
		coords := []int{}
		for _, s := range coordStr{
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			coords = append(coords, i)
		}

		brickCoords = append(brickCoords, coords)

		// startPoint := Point{coords[0], coords[1], coords[2]}
		// endPoint := Point{coords[3], coords[4], coords[5]}

		// brick := Brick{startPoint, endPoint, []Brick{}, []Brick{}}
		// bricks = append(bricks, brick)
	}

	sort.Slice(brickCoords, func(i ,j int) bool {
		return brickCoords[i][2] < brickCoords[j][2]
	})
	// fmt.Println("Brick coords", brickCoords)

	for i := 0 ; i < len(brickCoords); i ++ {
		coord := brickCoords[i]
		highest := 1
		for _, below := range brickCoords[:i]{
			if overlaps(coord, below){
				highest = Max(highest, below[5] + 1)
			}
		}
		// Assign new height
		coord[5] -= coord[2] - highest
		coord[2] = highest
	}

	sort.Slice(brickCoords, func(i ,j int) bool {
		return brickCoords[i][2] < brickCoords[j][2]
	})

	// fmt.Println("Brick coords", brickCoords)

	supporting := map[int][]int{}
	supported := map[int][]int{}

	for i, upper := range brickCoords {
		for j, lower := range brickCoords[:i]{
			if overlaps(upper,lower) && upper[2] == lower[5] + 1{
				if _, exists := supporting[j]; ! exists {
					supporting[j] = []int{}
				}
				supporting[j] = append(supporting[j], i)
				if _, exists := supported[i]; ! exists {
					supported[i] = []int{}
				}
				supported[i] = append(supported[i], j)
			}
		}
	}
	// fmt.Println("Supporting", PrettyFormat(supporting))
	// fmt.Println("Supported", PrettyFormat(supported))

	p1Total := 0

	for i := 0 ; i < len(brickCoords) ; i ++ {
		multiSupport := true
		for _, j := range supporting[i] {
			if len(supported[j]) < 2 {
				multiSupport = false
				break
			}
		}
		if multiSupport {
			p1Total += 1
		}
	}
	fmt.Println("P1 total", p1Total)

	p2Total := 0

	for i := 0 ; i < len(brickCoords) ; i ++ {
		q := []int{}
		falling := map[int]bool{i: true}
		for _, j := range supporting[i] {
			if len(supported[j]) == 1 {
				q = append(q, j)
				falling[j] = true
			}
		}

		for len(q) > 0 {
			next := q[0]
			q = q[1:]

			fallingKeys := []int{}
			for k, _ := range falling {
				fallingKeys = append(fallingKeys, k)
			}

			for _, support := range RemoveFromSlice(supporting[next], fallingKeys){
				if Subset(fallingKeys, supported[support]) {
					q = append(q, support)
					falling[support] = true
				}
			}
		}

		p2Total += len(falling) - 1
	}

	fmt.Println("P2 Total", p2Total)
}
