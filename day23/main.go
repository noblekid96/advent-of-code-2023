package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	X int
	Y int
}

type State struct {
	p Point // coordinate
	s int   // steps
}

var LEFT, RIGHT, UP, DOWN Point = Point{0,-1}, Point{0,1}, Point{-1,0}, Point{1,0}

var udlr = map[string]Point{
	"U": UP,
	"R": RIGHT,
	"D": DOWN,
	"L": LEFT,
}

var dirs = map[string][]Point{
    "^": []Point{{-1, 0}},
    "v": []Point{{1, 0}},
    "<": []Point{{0, -1}},
    ">": []Point{{0, 1}},
    ".": []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
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

func part1(grid []string){
	m := len(grid)
	n := len(grid[0])
	start := Point{0,0}
	end := Point{m-1,n-1}


	for i, c := range grid[0]{
		if c == '.' {
			start.Y = i
			break
		}
	}

	for i, c := range grid[m-1]{
		if c == '.' {
			end.Y = i
			break
		}
	}
	points := []Point{start, end}

	// fmt.Println("Sp Ep", start, end)

	for x, row := range grid{
		for y, char := range row {
			if char == '#' {
				continue
			}

			neighbours := 0

			for _, dir := range udlr {
				nx := x + dir.X
				ny := y + dir.Y

				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] != '#' {
					neighbours += 1
				}
			}
			if neighbours >= 3 {
				points = append(points, Point{x,y})
			}
		}
	}

	// fmt.Println("Points", points)

	graph := map[Point]map[Point]int{}
	for _, p := range points {
		graph[p] = map[Point]int{}
	}

	for _, sp := range points {
		stack := []State{{sp, 0}}
		seen := map[Point]bool{sp: true}

		for len(stack) > 0 {
			state := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]

			if state.s != 0 && slices.Contains(points, state.p){
				// fmt.Println("State entering graph for p", sp, state)
				graph[sp][state.p] = state.s
				continue
			}

			for _, dir := range dirs[string(grid[state.p.X][state.p.Y])]{
				nx := state.p.X + dir.X
				ny := state.p.Y + dir.Y

				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] != '#'{
					if _, exists := seen[Point{nx,ny}]; !exists {
						stack = append(stack, State{Point{nx,ny}, state.s + 1})
						seen[Point{nx,ny}] = true
					}
				}
			}
		}
	}


	seen := map[Point]bool{}
	// fmt.Println("Graph", graph)

	// for p, x := range graph {
	// 	fmt.Println("p", p)
	// 	fmt.Println("edge", x)
	// }

	var dfs func(pt Point) int

	dfs = func(pt Point) int {

		// fmt.Println("pt", pt)
		if pt == end {
			return 0
		}

		max := -999999999999999

		seen[pt] = true
		for np, dist := range graph[pt] {
			if _, exists := seen[np]; !exists {
				max = Max(max, dfs(np) + dist)
			}
		}
		delete(seen, pt)

		return max
	}

	ans := dfs(start)

	fmt.Println("P1 ans", ans)
}

func part2(grid []string){
	m := len(grid)
	n := len(grid[0])
	start := Point{0,0}
	end := Point{m-1,n-1}


	for i, c := range grid[0]{
		if c == '.' {
			start.Y = i
			break
		}
	}

	for i, c := range grid[m-1]{
		if c == '.' {
			end.Y = i
			break
		}
	}
	points := []Point{start, end}

	// fmt.Println("Sp Ep", start, end)

	for x, row := range grid{
		for y, char := range row {
			if char == '#' {
				continue
			}

			neighbours := 0

			for _, dir := range udlr {
				nx := x + dir.X
				ny := y + dir.Y

				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] != '#' {
					neighbours += 1
				}
			}
			if neighbours >= 3 {
				points = append(points, Point{x,y})
			}
		}
	}

	graph := map[Point]map[Point]int{}
	for _, p := range points {
		graph[p] = map[Point]int{}
	}

	for _, sp := range points {
		stack := []State{{sp, 0}}
		seen := map[Point]bool{sp: true}

		for len(stack) > 0 {
			state := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]

			if state.s != 0 && slices.Contains(points, state.p){
				// fmt.Println("State entering graph for p", sp, state)
				graph[sp][state.p] = state.s
				continue
			}

			for _, dir := range udlr{
				nx := state.p.X + dir.X
				ny := state.p.Y + dir.Y

				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] != '#'{
					if _, exists := seen[Point{nx,ny}]; !exists {
						stack = append(stack, State{Point{nx,ny}, state.s + 1})
						seen[Point{nx,ny}] = true
					}
				}
			}
		}
	}


	seen := map[Point]bool{}
	// fmt.Println("Graph", graph)

	// for p, x := range graph {
	// 	fmt.Println("p", p)
	// 	fmt.Println("edge", x)
	// }

	var dfs func(pt Point) int

	dfs = func(pt Point) int {

		// fmt.Println("pt", pt)
		if pt == end {
			return 0
		}

		max := -999999999999999

		seen[pt] = true
		for np, dist := range graph[pt] {
			if _, exists := seen[np]; !exists {
				max = Max(max, dfs(np) + dist)
			}
		}
		delete(seen, pt)

		return max
	}

	ans := dfs(start)

	fmt.Println("P2 ans", ans)
}
