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

type State struct {
	pos Point
	dir Point
	streak int // Streak
}

func dirLeft(p Point) Point {
	return Point{p.Y, -p.X}
}

func dirRight(p Point) Point {
	return Point{-p.Y, p.X}
}

func Min(i, j int) int {
	if i < j {
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

func PrettyFormat(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      return string(s)
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
	// PrintGrid(lines)
	part1(lines)
	part2(lines)
}

func countHeatLoss(grid []string, minStreak, maxStreak int) int {
	m := len(grid)
	n := len(grid[0])
	start, end := Point{0, 0}, Point{m-1, n-1}
	pointsToCheck := []State{{start, Point{1,0}, 0}, {start, Point{0,1}, 0}}
	visited := map[State]int{{start, Point{0,0}, 0}: 0}
	minHeatLoss := 99999999

	for len(pointsToCheck) > 0 {
		current := pointsToCheck[0]
		pointsToCheck = pointsToCheck[1:]

		if current.streak > maxStreak {
			continue
		}

		if current.pos == end && current.streak >= minStreak {
			minHeatLoss = Min(minHeatLoss, visited[current])
		}

		straight := State{Point{current.pos.X + current.dir.X, current.pos.Y + current.dir.Y}, current.dir, current.streak + 1}
		if straight.pos.X >= 0 && straight.pos.Y >= 0 && straight.pos.X < m && straight.pos.Y < n {
			totalHeatLoss := visited[current] + int(grid[straight.pos.X][straight.pos.Y] - '0')
			if v, exists := visited[straight]; !exists || v > totalHeatLoss {
				visited[straight] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, straight)
			}
		}

		leftDir := dirLeft(current.dir)
		left := State{Point{current.pos.X + leftDir.X, current.pos.Y + leftDir.Y}, leftDir, 1}
		if left.pos.X >= 0 && left.pos.Y >= 0 && left.pos.X < m && left.pos.Y < n && current.streak >= minStreak {
			totalHeatLoss := visited[current] + int(grid[left.pos.X][left.pos.Y] - '0')
			if v, exists := visited[left]; !exists || v > totalHeatLoss {
				visited[left] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, left)
			}
		}

		rightDir := dirRight(current.dir)
		right := State{Point{current.pos.X + rightDir.X, current.pos.Y + rightDir.Y}, rightDir, 1}
		if right.pos.X >= 0 && right.pos.Y >= 0 && right.pos.X < m && right.pos.Y < n && current.streak >= minStreak {
			totalHeatLoss := visited[current] + int(grid[right.pos.X][right.pos.Y] - '0')
			if v, exists := visited[right]; !exists || v > totalHeatLoss {
				visited[right] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, right)
			}
		}
	}

	return minHeatLoss
}

func part1(lines []string){
	p1 := countHeatLoss(lines, 0, 3)

	fmt.Printf("P1 result: %v\n", p1)
}

func part2(lines []string){
	p2 := countHeatLoss(lines, 4, 10)

	fmt.Printf("P2 result: %v\n", p2)
}
