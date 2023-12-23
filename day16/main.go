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
	D Direction
}

type Direction struct {
	X int
	Y int
}


func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}


var LEFT, RIGHT, UP, DOWN Direction = Direction{0,-1}, Direction{0,1}, Direction{-1,0}, Direction{1,0}


func MovePoint(p Point) Point {
	d := p.D
	np := Point{p.X+d.X, p.Y+d.Y, d}
	return np
}


func Rotate90Left(p Point) Point {
	d := p.D
	result := d
	switch d {
	case LEFT:
		result = DOWN
	case RIGHT:
		result = UP
	case DOWN:
		result = RIGHT
	case UP:
		result = LEFT
	default:
		panic(fmt.Sprintf("Unexpected point %v\n", p))
	}
	return Point{p.X,p.Y,result}
}


func Rotate90Right(p Point) Point {
	d := p.D
	result := d
	switch d {
	case LEFT:
		result = UP
	case RIGHT:
		result = DOWN
	case DOWN:
		result = LEFT
	case UP:
		result = RIGHT
	default:
		panic(fmt.Sprintf("Unexpected point %v\n", p))
	}
	return Point{p.X,p.Y,result}
}

func SplitterSplit(p Point, s rune) []Point {
	result := []Point{}
	d := p.D

	if s == '|' {
		if d == LEFT || d == RIGHT {
			return []Point{{p.X,p.Y,UP}, {p.X,p.Y,DOWN}}
		} else {
			return []Point{p}
		}
	} else if s == '-' {
		if d == UP || d == DOWN {
			return []Point{{p.X,p.Y,LEFT}, {p.X,p.Y,RIGHT}}
		} else {
			return []Point{p}
		}
	}
	// fmt.Println("Something wrong with split")
	return result
}

func MarkAsSeen(ps Point, m map[Point]bool) map[Point]bool {
	m[ps] = true
	return m
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

func energize(lines []string, sp Point) int {
	q := []Point{sp}

	m := len(lines)
	n := len(lines[0])
	// fmt.Println(m, n)

	seen := map[Point]bool{}


	for len(q) > 0 {

		p := q[0]
		// fmt.Printf("P : %v len q %v\n", PrettyFormat(p), len(q))
		q = q[1:]
		if _, exists := seen[p]; exists {
			// fmt.Printf("Point %v has been seen before\n", p)
			continue
		}
		// fmt.Printf("Seen: %v\n", PrettyFormat(seen))
		// for k,_ := range seen{
		// 	fmt.Printf("%v,", k)
		// }
		// fmt.Printf("\n\n")

		pX := p.X
		pY := p.Y
		pD := p.D

		if (pX < 0 || pX >= m || pY < 0 || pY >= n){
			// Out of bounds
			// fmt.Println("Reached 0")
			continue
		}
		// fmt.Printf("Pushing p %v\n", p)
		seen = MarkAsSeen(p, seen)

		current := lines[pX][pY]
		np := p

		if current == '.' {
			// fmt.Println("Reached 1")
			np = MovePoint(p)
			// fmt.Printf("NP : %v len q %v\n", PrettyFormat(np), len(q))
			q = append(q,np)
			continue
		}

		if current == '/'{
			// fmt.Println("Reached 2")
			switch pD {
			case RIGHT:
				np = Rotate90Left(p)
			case LEFT:
				np = Rotate90Left(p)
			case UP:
				np = Rotate90Right(p)
			case DOWN:
				np = Rotate90Right(p)
			default:
				panic(fmt.Sprintf("Unexpected point %v\n", p))
			}
			// fmt.Printf("New point after rotate: %v\n", np)
			np := MovePoint(np)
			q = append(q,np)
			continue
		} else if current == '\\'{
			// fmt.Println("Reached 3")
			switch pD {
			case RIGHT:
				np = Rotate90Right(p)
			case LEFT:
				np = Rotate90Right(p)
			case UP:
				np = Rotate90Left(p)
			case DOWN:
				np = Rotate90Left(p)
			default:
				panic(fmt.Sprintf("Unexpected point %v\n", p))
			}
			// fmt.Printf("New point after rotate: %v\n", np)
			np := MovePoint(np)
			q = append(q,np)
			continue
		} else if current == '|' || current == '-' {
			// fmt.Println("Reached splitter")
			new_points := SplitterSplit(p, rune(current))
			for _, new := range new_points {
				// fmt.Printf("Adding new point: %v", new)
				q = append(q,MovePoint(new))
			}
		}
	}

	grid := [][]string{}
	for i := 0 ; i < m ; i ++ {
		row := []string{}
		for j := 0 ; j < n; j ++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	total := 0
	for p := range seen {
		grid[p.X][p.Y] = "#"
	}

	for i := 0 ; i < m ; i ++ {
		for j := 0 ; j < n; j ++ {
			// fmt.Printf(grid[i][j])
			if grid[i][j] == "#"{
				total +=1
			}
		}
		// fmt.Println()
	}

	return total
}

func part1(lines []string){
	sp := Point{0,0, Direction{0,1}}

	total := energize(lines, sp)

	fmt.Printf("P1 Total: %v\n", total)
}

func part2(lines []string){

	m := len(lines)
	n := len(lines[0])

	total := 0

	for i := 0 ; i < n ; i ++ {
		for _, d := range []Direction{UP, DOWN, LEFT, RIGHT}{
			total = Max(total, energize(lines, Point{0, i, d}))
			total = Max(total, energize(lines, Point{m-1, i, d}))
		}
	}

	for i := 0 ; i < m ; i ++ {
		for _, d := range []Direction{UP, DOWN, LEFT, RIGHT}{
			total = Max(total, energize(lines, Point{i, 0, d}))
		// total = Max(total, energize(lines, Point{i, 0, Direction{0, -1}}))
			total = Max(total, energize(lines, Point{i, n-1, d}))
		// total = Max(total, energize(lines, Point{i, n-1, Direction{0, -1}}))
		}
	}

	fmt.Printf("P2 Total: %v\n", total)
}
