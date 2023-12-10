package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

var One2ThreeMAP = map[string]string{
	"|": ".#..#..#.",
	"-": "...###...",
	"L": ".#..##...",
	"J": ".#.##....",
	"7": "...##..#.",
	"F": "....##.#.",
	".": ".........",
	"S": "....#....",
}

var dirs = []Point{
	Point{-1,0},
	Point{0,-1},
	Point{0,1},
	Point{1,-0},
}

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

func splitLinebreak(c rune) bool {
	return c == '\n'
}

func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
}

func PrettyFormatGrid(g [][]string) string {
	out := ""
	for _, r := range g {
		out += strings.Join(r,"")
		out += "\n"
	}
	return out
}

func ExpandGrid(g [][]string) [][]string{
	n := len(g)
	m := len(g[0])

	newGrid := make([][]string, 3*n)
	for i := range newGrid {
		newGrid[i] = make([]string, 3*m)
	}

	for x, r := range g {
		for y, c := range r {
			for dx := 0; dx < 3; dx ++ {
				for dy := 0; dy < 3; dy ++ {
					newGrid[x*3 + dx][y*3 + dy] = string(One2ThreeMAP[c][dx*3 + dy])
				}
			}
		}
	}

	return newGrid
}

func part1(lines []string){
	grid := [][]string{}
	for _, l := range lines {
		row := []string{}
		for _, c := range l {
			row = append(row, string(c))
		}
		grid = append(grid, row)
	}

	r_start, c_start := 0, 0

	for i, r := range grid {
		for j, c := range r {
			if c == "S" {
				r_start, c_start = i,j
			}
		}
	}


	fmt.Println(PrettyFormatGrid(grid))
	newGrid := ExpandGrid(grid)
	fmt.Println("New Grid")
	fmt.Println(PrettyFormatGrid(newGrid))

	// Connect S to neighbouring openings

	r, c := r_start*3 +1 , c_start*3 + 1

	if newGrid[r + 2][c] == "#" {
		newGrid[r + 1][c] = "#"
	}
	if newGrid[r - 2][c] == "#" {
		newGrid[r - 1][c] = "#"
	}
	if newGrid[r][c + 2] == "#"{
		newGrid[r][c + 1] = "#"
	}
	if newGrid[r][c - 2] == "#"{
		newGrid[r][c - 1] = "#"
	}

	fmt.Println("Connected New Grid")
	fmt.Println(PrettyFormatGrid(newGrid))

	q := []Point{}
	seen := map[Point]bool{}

	q = append(q, Point{r,c})


	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for _, dir := range dirs {
			dx := dir.X
			dy := dir.Y

			nx := p.X + dx
			ny := p.Y + dy

			if 0 <= nx && nx < len(newGrid) && 0 <= ny && ny < len(newGrid[0]) && newGrid[nx][ny] == "#"{
				if _, exists := seen[Point{nx,ny}]; !exists {
					seen[Point{nx,ny}] = true
					q = append(q, Point{nx,ny})
				}
			}
		}
	}

	println("Part 1 solution: ", len(seen)/6)
}

func part2(lines []string){
	grid := [][]string{}
	for _, l := range lines {
		row := []string{}
		for _, c := range l {
			row = append(row, string(c))
		}
		grid = append(grid, row)
	}

	r_start, c_start := 0, 0

	for i, r := range grid {
		for j, c := range r {
			if c == "S" {
				r_start, c_start = i,j
			}
		}
	}


	fmt.Println(PrettyFormatGrid(grid))
	newGrid := ExpandGrid(grid)
	fmt.Println("New Grid")
	fmt.Println(PrettyFormatGrid(newGrid))

	// Connect S to neighbouring openings

	r, c := r_start*3 +1 , c_start*3 + 1

	if newGrid[r + 2][c] == "#" {
		newGrid[r + 1][c] = "#"
	}
	if newGrid[r - 2][c] == "#" {
		newGrid[r - 1][c] = "#"
	}
	if newGrid[r][c + 2] == "#"{
		newGrid[r][c + 1] = "#"
	}
	if newGrid[r][c - 2] == "#"{
		newGrid[r][c - 1] = "#"
	}

	fmt.Println("Connected New Grid")
	fmt.Println(PrettyFormatGrid(newGrid))

	q := []Point{}
	seen := map[Point]bool{}

	q = append(q, Point{r,c})


	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for _, dir := range dirs {
			dx := dir.X
			dy := dir.Y

			nx := p.X + dx
			ny := p.Y + dy

			if 0 <= nx && nx < len(newGrid) && 0 <= ny && ny < len(newGrid[0]) && newGrid[nx][ny] == "#" {
				if _, exists := seen[Point{nx,ny}]; !exists {
					seen[Point{nx,ny}] = true
					q = append(q, Point{nx,ny})
				}
			}
		}
	}

	for r := 0 ; r < len(newGrid) ; r ++ {
		for  c := 0 ; c < len(newGrid[0]); c ++ {
			if _, exists := seen[Point{r,c}]; !exists {
				newGrid[r][c] = "."
			}
		}
	}

	q = []Point{}
	q = append(q, Point{0,0})
	seen = map[Point]bool{}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		newGrid[p.X][p.Y] = " "

		for _, dir := range dirs {
			nx, ny := p.X + dir.X, p.Y + dir.Y
			if 0 <= nx && nx < len(newGrid) && 0 <= ny && ny < len(newGrid[0]) && newGrid[nx][ny] == "." {
				if _, exists := seen[Point{nx,ny}]; !exists {
					seen[Point{nx,ny}] = true
					q = append(q, Point{nx,ny})
				}
			}
		}
	}

	p2 := 0

	for r := 0; r < len(grid); r ++ {
		for c := 0; c < len(grid[0]); c ++ {
			allDots := true
			zoomed:
			for dr := 0; dr < 3; dr ++ {
				for dc := 0; dc < 3; dc ++ {
					if newGrid[r*3 + dr][c*3 + dc] != "." {
						allDots = false
						break zoomed
					}
				}
			}
			if allDots {
				p2 += 1
			}
		}
	}

	println("Part 2 solution: ", p2)
}
