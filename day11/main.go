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

type PointWithSteps struct {
	P Point
	S int
}

var dirs = []Point{
	Point{-1,0},
	Point{0,-1},
	Point{0,1},
	Point{1,-0},
}

func Min(i,j int) int {
	if i > j {
		return j
	}
	return i
}

func Max(i,j int) int {
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

	lines := strings.Fields(input)

	grid := [][]string{}
	for _, l := range lines {
		row := []string{}
		for _, c := range l {
			row = append(row, string(c))
		}
		grid = append(grid, row)
	}

	// fmt.Printf("Grid:\n%v\n", PrettyFormatGrid(grid))

	part1(grid)
	part2(grid)
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

func IndexOf[T comparable](collection []T, el T) int {
    for i, x := range collection {
        if x == el {
            return i
        }
    }
    return -1
}

func part1(grid [][]string){
	emptyCols := []int{}
	emptyRows := []int{}
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i ++ {
		empty := true
		for j := 0; j < n; j ++ {
			if grid[i][j] == "#"{
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows,i)
			continue
		}
	}

	for i := 0; i < n; i ++ {
		empty := true
		for j := 0; j < m; j ++ {
			if grid[j][i] == "#"{
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols,i)
			continue
		}
	}

	// fmt.Printf("Empty rows:\n%v\n Empty cols:\n%v\n", PrettyFormat(emptyRows), PrettyFormat(emptyCols))

	m = len(grid)
	n = len(grid[0])
	points := []Point{}

	for i := 0; i < m; i ++ {
		for j := 0; j < n; j ++ {
			if grid[i][j] == "#" {
				points = append(points, Point{i,j})
			}
		}
	}

	// fmt.Printf("Points:\n%v\n. Len points: %v\n", PrettyFormat(points), len(points) )

	scale := 2
	total := 0

	for i, sp := range points{
		for _, ep := range points[i+1:]{
			x1 := sp.X
			y1 := sp.Y

			x2 := ep.X
			y2 := ep.Y

			for x := Min(x1,x2); x < Max(x1,x2); x ++ {
				if slices.Contains(emptyRows, x){
					total += scale
				} else {
					total += 1
				}
			}

			for y := Min(y1,y2); y < Max(y1,y2); y ++ {
				if slices.Contains(emptyCols, y){
					total += scale
				} else {
					total += 1
				}
			}
		}
	}

	fmt.Printf("P1 Total steps: %v\n", total)
}

func part2(grid [][]string){
	emptyCols := []int{}
	emptyRows := []int{}
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i ++ {
		empty := true
		for j := 0; j < n; j ++ {
			if grid[i][j] == "#"{
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows,i)
			continue
		}
	}

	for i := 0; i < n; i ++ {
		empty := true
		for j := 0; j < m; j ++ {
			if grid[j][i] == "#"{
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols,i)
			continue
		}
	}
	// fmt.Printf("Empty rows:\n%v\n Empty cols:\n%v\n", PrettyFormat(emptyRows), PrettyFormat(emptyCols))

	m = len(grid)
	n = len(grid[0])
	points := []Point{}

	for i := 0; i < m; i ++ {
		for j := 0; j < n; j ++ {
			if grid[i][j] == "#" {
				points = append(points, Point{i,j})
			}
		}
	}

	// fmt.Printf("Points:\n%v\n. Len points: %v\n", PrettyFormat(points), len(points) )

	scale := 1000000
	total := 0

	for i, sp := range points{
		for _, ep := range points[i+1:]{
			x1 := sp.X
			y1 := sp.Y

			x2 := ep.X
			y2 := ep.Y

			for x := Min(x1,x2); x < Max(x1,x2); x ++ {
				if slices.Contains(emptyRows, x){
					total += scale
				} else {
					total += 1
				}
			}

			for y := Min(y1,y2); y < Max(y1,y2); y ++ {
				if slices.Contains(emptyCols, y){
					total += scale
				} else {
					total += 1
				}
			}
		}
	}

	fmt.Printf("P2 Total steps: %v\n", total)

}
