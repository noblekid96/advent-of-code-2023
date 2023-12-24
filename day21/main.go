package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

type State struct {
	p Point
	s int
}


var LEFT, RIGHT, UP, DOWN Point = Point{0,-1}, Point{0,1}, Point{-1,0}, Point{1,0}

var dirs = map[string]Point{
	"U": UP,
	"R": RIGHT,
	"D": DOWN,
	"L": LEFT,
}

func PrintGrid(grid []string){
	m := len(grid)
	n := len(grid[0])
	for i := 0 ; i < m ; i ++ {
		for j := 0 ; j < n; j ++ {
			fmt.Printf(string(grid[i][j]))
		}
		fmt.Println()
	}
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

	lines := strings.FieldsFunc(input, splitFn)
	part1(lines)
	part2(lines)
}

func possibleSpots(lines []string, stop int) int {
	m := len(lines)
	n := len(lines[0])
	var sp Point
	for i := 0 ; i < m; i ++ {
		for j := 0 ; j < n; j ++ {
			if lines[i][j] == 'S' {
				sp = Point{i,j}
				break
			}
		}
	}
	ans := 0
	seen := map[Point]bool{sp: true}
	q := []State{{sp, 0}}

	for len(q) > 0 {
		state := q[0]
		q = q[1:]

		x := state.p.X
		y := state.p.Y
		s := state.s

		if (s % 2) == (stop % 2) {
			ans += 1
		}

		if s == stop+1 {
			continue
		}

		for _, d := range dirs {
			dx := d.X
			dy := d.Y

			nx := x+dx
			ny := y+dy

			// if nx < 0 || nx >= m || ny < 0 || ny >= n || lines[nx % 131][ny % 131] == '#' {
			// 	continue
			// }
			modx := nx % 131
			mody := ny % 131
			if modx < 0 {
				modx += 131
			}

			if mody < 0 {
				mody += 131
			}
			if lines[modx][mody] == '#' {
				continue
			}
			if _, exists := seen[Point{nx,ny}]; exists {
				continue
			}

			seen[Point{nx,ny}] = true
			q = append(q, State{Point{nx,ny}, s+1})
		}
	}

	// fmt.Println("Ans", ans)
	// fmt.Println("P1 ans", len(ans))
	return ans
}

func part1(lines []string){
	p1 := possibleSpots(lines, 64)
	fmt.Println("P1 ans", p1)
}

func quad(y []int, n int) int {
	a := (y[2] - (2 * y[1]) + y[0]) / 2
	b := y[1] - y[0] - a
	c := y[0]
    return (a * n*n) + (b * n) + c
}

func part2(lines []string){
    goal := 26501365
    size := len(lines)
    edge := size / 2

	y := []int{}

	for i := 0 ; i < 3 ; i ++ {
		calc := possibleSpots(lines, (edge + i * size))
		y = append(y, calc)
	}

	// fmt.Println("Y", y)
    ans := quad(y, ((goal - edge) / size))
	fmt.Println("Ans", ans)
}
