package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hailstone struct {
	sx float64
	sy float64
	sz float64
	vx float64
	vy float64
	vz float64
	a float64
	b float64
	c float64
}

func splitCommaFn (c rune) bool {
	return c == ','
}

func splitAdFn (c rune) bool {
	return c == '@'
}

func initHailstone(sx,sy,sz,vx,vy,vz int) Hailstone {
	a := vy
	b := -1*vx
	c := vy * sx - vx * sy

	return Hailstone{float64(sx),float64(sy),float64(sz),float64(vx),float64(vy),float64(vz),float64(a),float64(b),float64(c)}
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

	hailstones := []Hailstone{}
	for _, l := range lines {
		fields := strings.FieldsFunc(l, splitAdFn)

		startPosStr := strings.FieldsFunc(fields[0], splitCommaFn)
		velocityStr := strings.FieldsFunc(fields[1], splitCommaFn)

		startPos := []int{}
		velocity := []int{}

		// fmt.Println("Start pos | Velocity:", startPosStr , "|" , velocityStr)
		for _, sps := range startPosStr {
			sps = strings.TrimSpace(sps)
			sp, err := strconv.Atoi(sps)
			if err != nil {
				panic(err)
			}
			startPos = append(startPos, sp)
		}

		for _, vs := range velocityStr {
			vs = strings.TrimSpace(vs)
			v, err := strconv.Atoi(vs)
			if err != nil {
				panic(err)
			}
			velocity = append(velocity , v)
		}
		// fmt.Println("Start pos | Velocity:", startPos, "|" , velocity)

		hailstones = append(hailstones, initHailstone(
			startPos[0], startPos[1], startPos[2],
			velocity[0], velocity[1], velocity[2],
		))
	}

	total := 0

	for i, hs1 := range hailstones {
		for _, hs2 := range hailstones[0:i] {
			a1, b1, c1 := hs1.a, hs1.b, hs1.c
			a2, b2, c2 := hs2.a, hs2.b, hs2.c

			if a1*b2 == b1*a2 {
				continue
			}

			x := (c1 * b2 - c2 * b1) / (a1 * b2 - a2 * b1)
			y := (c2 * a1 - c1 * a2) / (a1 * b2 - a2 * b1)

			if 200000000000000 <= x && x <= 400000000000000 && 200000000000000 <= y && y <= 400000000000000 {
				allPosVelocity := true

				for _, hs := range []Hailstone{hs1, hs2}{
					if (x - (hs.sx)) * (hs.vx) >= 0 && (y - (hs.sy)) * (hs.vy) >= 0 {
						continue
					} else {
						allPosVelocity = false
					}
				}

				if allPosVelocity {
					total += 1
				}
			}
		}
	}

	fmt.Println("P1 ans", total)
}

func part2(lines []string){

}
