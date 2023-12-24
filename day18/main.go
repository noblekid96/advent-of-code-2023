package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
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

	points := []Point{{0,0}}
	bound := 0

	for _, l := range lines {
		fields := strings.Fields(l)
		d := fields[0]
		steps, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		dir := dirs[d]

		lastPoint := points[len(points)-1]
		bound += steps
		points = append(points, Point{lastPoint.X + dir.X*steps, lastPoint.Y + dir.Y*steps})
	}

	// fmt.Printf("POints %v\n", points)

	area := 0
	for i := 0 ; i < len(points); i ++ {
		area += points[i].X * (points[(i-1+len(points))%len(points)].Y - points[(i+1)%len(points)].Y)
	}
	area = int(math.Abs((float64(area))))/2

	total := area-(bound/2)+1 + bound


	// fmt.Printf("Area %v Bound: %v\n", area, bound)
	fmt.Printf("Total %v\n", total)
}

func part2(lines []string){
	points := []Point{{0,0}}
	bound := 0

	for _, l := range lines {
		fields := strings.Fields(l)
		// d := fields[0]
		color := fields[2]
		color = color[2:len(color)-1]
		last, _ := strconv.Atoi(string(color[len(color)-1]))
		dir := dirs[string("RDLU"[last])]
		steps64, err :=  strconv.ParseInt(color[:len(color)-1], 16, 64)
		if err != nil {
			panic(err)
		}
		steps := int(steps64)
		// dir := dirs[d]

		lastPoint := points[len(points)-1]
		bound += steps
		points = append(points, Point{lastPoint.X + dir.X*steps, lastPoint.Y + dir.Y*steps})
	}

	// fmt.Printf("POints %v\n", points)

	area := 0
	for i := 0 ; i < len(points); i ++ {
		area += points[i].X * (points[(i-1+len(points))%len(points)].Y - points[(i+1)%len(points)].Y)
	}
	area = int(math.Abs((float64(area))))/2

	total := area-(bound/2)+1 + bound


	// fmt.Printf("Area %v Bound: %v\n", area, bound)
	fmt.Printf("Total %v\n", total)
}
