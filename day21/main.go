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

func part1(lines []string){
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
	ans := map[Point]bool{}
	seen := map[Point]bool{sp: true}
	q := []State{{sp, 64}}

	for len(q) > 0 {
		state := q[0]
		q = q[1:]

		x := state.p.X
		y := state.p.Y
		s := state.s

		if s % 2 == 0 {
			ans[Point{x,y}] = true
		}

		if s == 0 {
			continue
		}

		for _, d := range dirs {
			dx := d.X
			dy := d.Y

			nx := x+dx
			ny := y+dy

			if nx < 0 || nx >= m || ny < 0 || ny >= n || lines[nx][ny] == '#' {
				continue
			}
			if _, exists := seen[Point{nx,ny}]; exists {
				continue
			}

			seen[Point{nx,ny}] = true
			q = append(q, State{Point{nx,ny}, s-1})
		}
	}

	// fmt.Println("Ans", ans)
	fmt.Println("P1 ans", len(ans))
}

func part2(lines []string){

}
