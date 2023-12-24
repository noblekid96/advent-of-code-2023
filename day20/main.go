package main

import (
	"fmt"
	"os"
	"slices"

	// "strconv"
	"strings"
)

func splitFn (c rune) bool {
	return c == '\n'
}

func splitCommaFn (c rune) bool {
	return c == ','
}

type Module struct {
	name string
	ttype string
	outputs []string
	memActive bool
	memory map[string]string
}

type Pulse struct {
	origin string
	target string
	pulse string
}

func initModule(name, ttype string, outputs []string) Module {
	var memActive bool
	var memory map[string]string
	if ttype == "%"{
		memActive = false
		memory = map[string]string{}
	} else {
		memActive = true
		memory = map[string]string{}
	}

	return Module{name, ttype, outputs, memActive, memory}
}

func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	input := string(bytes)

	lines := strings.FieldsFunc(input, splitFn)
	part1(lines)
	part2(lines)
}

func part1(lines []string){
	modules := map[string]Module{}
	broadcastTargets := []string{}

	for _, l := range lines {
		fields := strings.Split(l, " -> ")
		left := strings.Trim(fields[0]," ")
		right := strings.Trim(fields[1]," ")

		// fmt.Println("left", left)
		// fmt.Println("right", right)

		if left == "broadcaster"{
			broadcastTargets = strings.Split(right, ", ")
		} else {
			ttype := string(left[0])
			name := left[1:]
			modules[name] = initModule(name, ttype, strings.Split(right, ", "))
		}
	}
	// fmt.Println("broadcastTargets", broadcastTargets)
	// fmt.Println("modules", modules)

	for name, module := range modules {
		for _, output := range module.outputs {
			if outputMod, exists := modules[output]; exists{
				if modules[output].ttype == "&" {
					outputMod.memory[name] = "lo"
					modules[output] = outputMod
				}
			}
		}
	}
	// fmt.Println("modules after pulse", modules)

	lo := 0
	hi := 0
	for i := 0 ; i < 1000 ; i ++ {
		lo += 1

		// origin, target, pulse
		q := []Pulse{}
		for _, bt := range broadcastTargets {
			q = append(q, Pulse{"broadcaster", bt, "lo"})
		}

		for len(q) > 0 {
			p := q[0]
			q = q[1:]

			origin := p.origin
			target := p.target
			pulse := p.pulse

			if pulse == "lo"{
				lo += 1
			} else {
				hi += 1
			}

			if _, exists := modules[target]; !exists {
				continue
			}

			module := modules[target]

			var outgoing string
			if module.ttype == "%" {
				if pulse == "lo" {
					module.memActive = !module.memActive
					if module.memActive {
						outgoing = "hi"
					} else {
						outgoing = "lo"
					}

					for _, x := range module.outputs {
						q = append(q, Pulse{module.name, x, outgoing})
					}
				}
			} else {
				module.memory[origin] = pulse
				allHigh := true
				for _, m := range module.memory {
					if m == "lo"{
						allHigh = false
						break
					}
				}
				if allHigh {
					outgoing = "lo"
				} else {
					outgoing = "hi"
				}

				for _, x := range module.outputs {
					q = append(q, Pulse{module.name, x, outgoing})
				}
			}

			modules[target] = module
		}
	}

	total := lo * hi

	fmt.Println("P1 total", total)
}

func part2(lines []string){
	modules := map[string]Module{}
	broadcastTargets := []string{}

	for _, l := range lines {
		fields := strings.Split(l, " -> ")
		left := strings.Trim(fields[0]," ")
		right := strings.Trim(fields[1]," ")

		// fmt.Println("left", left)
		// fmt.Println("right", right)

		if left == "broadcaster"{
			broadcastTargets = strings.Split(right, ", ")
		} else {
			ttype := string(left[0])
			name := left[1:]
			modules[name] = initModule(name, ttype, strings.Split(right, ", "))
		}
	}
	// fmt.Println("broadcastTargets", broadcastTargets)
	// fmt.Println("modules", modules)

	for name, module := range modules {
		for _, output := range module.outputs {
			if outputMod, exists := modules[output]; exists{
				if modules[output].ttype == "&" {
					outputMod.memory[name] = "lo"
					modules[output] = outputMod
				}
			}
		}
	}
	// fmt.Println("modules after pulse", modules)
	var beforeRx string
	for _, m := range modules {
		if slices.Contains(m.outputs, "rx"){
			beforeRx = m.name
			break
		}
	}
	fmt.Println("beforeRx", beforeRx)

	cycleLengths := map[string]int{}
	seen := map[string]int{}

	for _, m := range modules {
		if slices.Contains(m.outputs, beforeRx) {
			// cycleLengths[m.name] = 0
			seen[m.name] = 0
		}
	}
	fmt.Println("cycleLengths", cycleLengths)

	presses := 0
	infinite:
	for {
		presses += 1
		// origin, target, pulse
		q := []Pulse{}
		for _, bt := range broadcastTargets {
			q = append(q, Pulse{"broadcaster", bt, "lo"})
		}

		for len(q) > 0 {
			p := q[0]
			q = q[1:]

			origin := p.origin
			target := p.target
			pulse := p.pulse

			if _, exists := modules[target]; !exists {
				continue
			}

			module := modules[target]

			if module.name == beforeRx && pulse == "hi" {
				seen[origin] += 1
				if _, exists := cycleLengths[origin]; !exists {
					cycleLengths[origin] = presses
				} else {
					if presses != seen[origin] * cycleLengths[origin] {
						panic("Cycle is off")
					}
				}


				allSeen := true
				for _, s := range seen {
					if s == 0 {
						allSeen = false
						break
					}
				}

				if allSeen {
					break infinite
				}
			}

			var outgoing string
			if module.ttype == "%" {
				if pulse == "lo" {
					module.memActive = !module.memActive
					if module.memActive {
						outgoing = "hi"
					} else {
						outgoing = "lo"
					}

					for _, x := range module.outputs {
						q = append(q, Pulse{module.name, x, outgoing})
					}
				}
			} else {
				module.memory[origin] = pulse
				allHigh := true
				for _, m := range module.memory {
					if m == "lo"{
						allHigh = false
						break
					}
				}
				if allHigh {
					outgoing = "lo"
				} else {
					outgoing = "hi"
				}

				for _, x := range module.outputs {
					q = append(q, Pulse{module.name, x, outgoing})
				}
			}

			modules[target] = module
		}
	}

	fmt.Println("cycleLengths after presses", cycleLengths)

	total := 1

	for _, s := range cycleLengths {
		total *= s
	}

	fmt.Println("P2 total", total)

}
