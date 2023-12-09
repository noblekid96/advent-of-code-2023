package main

import (
	"fmt"
	"os"
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
	input := string(bytes)
	lines := strings.Split(input, "\n")
	part1(lines)
	part2(lines)
}

// Point struct to represent x, y coordinates
type Point struct {
    X int
    Y int
}

type Num struct {
	value int
	id Point
}

func IsSymbol(r rune) bool{
	return !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) && !(r == '.')
}

func IsStar(r rune) bool{
	return r == '*'
}


func part1(lines []string){
	numbers := make(map[Point]Num)
	symbols := []Point{}

	for i := 0; i < len(lines); i++ {
		line := []rune(lines[i])
		num := 0
		for j := 0 ; j < len(line); j ++ {
			if unicode.IsDigit(line[j]){
				if _, ok := numbers[Point{i,j}]; !ok {
					k := j
					// fmt.Printf("initial k %v || ", k)
					num = int(line[j] - '0')
					numbers[Point{i,k}] = Num{0, Point{i,j}}
					for k+1 < len(line){
						k+=1
						if unicode.IsDigit(line[k]){
							numbers[Point{i,k}] = Num{0, Point{i,j}}
							num = num * 10 + int(line[k] - '0')
						} else{
							break
						}
					}
				}
				// fmt.Printf("%v %v : %v || ", i,j, num)
				numbers[Point{i,j}] = Num{num, numbers[Point{i,j}].id}
			} else{
				num = 0
				if ( line[j] == '.' ){
					continue
				} else {
					if (IsSymbol(line[j])){
						// fmt.Printf("Symbol: %v\n", line[j])
						symbols = append(symbols, Point{i,j})
					}
				}
			}
		}
	}

	dir := []Point{
		Point{-1,-1}, Point{-1,0}, Point{-1,1},
		Point{0,-1}, Point{0,1},
		Point{1,-1}, Point{1,-0}, Point{1,1},
	}
	sum := 0

	for _,point := range symbols {
		x := point.X
		y := point.Y
		dx := x
		dy := x
		encounted := map[Point]bool{}

		for _, d := range dir{
			dx = x+d.X
			dy = y+d.Y
			// fmt.Printf("dx dy: {%v,%v}\n", dx,dy)
			if val , ok := numbers[Point{dx,dy}]; ok {
				// fmt.Printf("Adding num: %v at dx dy: {%v,%v}\n", val, dx,dy)
				if _, ok := encounted[val.id]; !ok {
					encounted[val.id] = true
					sum += val.value
				}
			}
		}
	}

	// fmt.Printf("Numbers: %v\n", numbers)
	// fmt.Printf("Symbols: %v\n", symbols)
	fmt.Printf("Sum: %v\n", sum)
}

func part2(lines []string){
	numbers := make(map[Point]Num)
	symbols := []Point{}

	for i := 0; i < len(lines); i++ {
		line := []rune(lines[i])
		num := 0
		for j := 0 ; j < len(line); j ++ {
			if unicode.IsDigit(line[j]){
				if _, ok := numbers[Point{i,j}]; !ok {
					k := j
					// fmt.Printf("initial k %v || ", k)
					num = int(line[j] - '0')
					numbers[Point{i,k}] = Num{0, Point{i,j}}
					for k+1 < len(line){
						k+=1
						if unicode.IsDigit(line[k]){
							numbers[Point{i,k}] = Num{0, Point{i,j}}
							num = num * 10 + int(line[k] - '0')
						} else{
							break
						}
					}
				}
				// fmt.Printf("%v %v : %v || ", i,j, num)
				numbers[Point{i,j}] = Num{num, numbers[Point{i,j}].id}
			} else{
				num = 0
				if (IsSymbol(line[j])){
					// fmt.Printf("Symbol: %v\n", line[j])
					symbols = append(symbols, Point{i,j})
				} else {
					continue
				}
			}
		}
	}

	dir := []Point{
		Point{-1,-1}, Point{-1,0}, Point{-1,1},
		Point{0,-1}, Point{0,1},
		Point{1,-1}, Point{1,-0}, Point{1,1},
	}
	sum := 0

	for _,point := range symbols {
		x := point.X
		y := point.Y
		dx := x
		dy := x
		encounted := map[Point]bool{}
		adjacent := []int{}

		for _, d := range dir{
			dx = x+d.X
			dy = y+d.Y
			// fmt.Printf("dx dy: {%v,%v}\n", dx,dy)
			if val , ok := numbers[Point{dx,dy}]; ok {
				// fmt.Printf("Adding num: %v at dx dy: {%v,%v}\n", val, dx,dy)
				if _, ok := encounted[val.id]; !ok {
					encounted[val.id] = true
					adjacent = append(adjacent, val.value)
				}
			}
		}
		if len(adjacent) == 2 {
			ratio := 1
			for _, n := range adjacent{
				ratio *= n
			}
			sum += ratio
		}
	}

	// fmt.Printf("Numbers: %v\n", numbers)
	// fmt.Printf("Symbols: %v\n", symbols)
	fmt.Printf("Sum: %v\n", sum)
}
