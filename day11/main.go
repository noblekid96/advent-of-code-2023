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

	fmt.Printf("Grid:\n%v\n", PrettyFormatGrid(grid))

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

	fmt.Printf("Empty rows:\n%v\n Empty cols:\n%v\n", PrettyFormat(emptyRows), PrettyFormat(emptyCols))

	counter := 0
	for _,r := range emptyRows{
		grid = append(grid[:r+1], grid[r:]...)
		grid[r+counter] = grid[r+counter+1]
		counter++
	}
	fmt.Printf("Grid:\n%v\n", PrettyFormatGrid(grid))

	counter = 0
	for _,c := range emptyCols{
		for i := range grid {
			row := grid[i]
			grid[i] = append(row[:c+counter], append([]string{"."}, row[c+counter:]...)...)
		}
		counter++
	}
	fmt.Printf("Grid:\n%v\n", PrettyFormatGrid(grid))

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

	fmt.Printf("Points:\n%v\n", PrettyFormat(points) )

	path := map[Point]map[Point]int{}

	for len(points) > 0 {
		sp := points[0]

		path[sp] = make(map[Point]int)


		q := []PointWithSteps{}
		seen := map[Point]bool{}
		q = append(q, PointWithSteps{sp, 0})
		seen[sp] = true

		for len(q) > 0 {
			p := q[0]
			q = q[1:]

			for _, dir := range dirs {
				dx := dir.X
				dy := dir.Y

				nx := p.P.X + dx
				ny := p.P.Y + dy

				if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[0]){
					if _, exists := seen[Point{nx,ny}]; !exists {
						seen[Point{nx,ny}] = true
						q = append(q, PointWithSteps{Point{nx,ny}, p.S+1})

						if grid[nx][ny] == "#"{
							path[sp][Point{nx,ny}] = p.S+1
						}
					}
				}
			}
		}
		points = points[1:]
	}

	fmt.Println(len(path))

	for k,v := range path{
		for kk, vv := range v{
			fmt.Printf("Steps from %v to %v is : %v\n", k, kk, vv)
		}
	}


	fmt.Printf("Paths: %v\n", PrettyFormat(path))
}

func part2(grid [][]string){

}
