package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input_file := os.Args[1]
	bytes, err := os.ReadFile(input_file)
	if err != nil {
		err_msg := fmt.Sprintf("%s not found or readable\n", input_file)
		panic(err_msg)
	}
	input := string(bytes)
	// lines := strings.Split(input, "\n")
	part1(input)
	part2(input)
}

func PrettyPrint(i interface{}) string {
      s, _ := json.MarshalIndent(i, "", "\t")
      // fmt.Println(string(s))
      return string(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i,j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func SortRankedHands(h []Hand) []Hand {

	sort.Slice(h, func(i int,j int) bool {
		return StrongerHand(h[i],h[j])
	})

	return h
}


type Hand struct {
	Grade int
	Cards []int
	Score int
}

func StrongerHand(i Hand, j Hand) bool{
	if j.Grade > i.Grade {
		return true
	} else if j.Grade < i.Grade {
		return false
	}

	for k := 0 ; k < 5; k ++ {
		if j.Cards[k] > i.Cards[k] {
			return true
		} else if j.Cards[k] < i.Cards[k] {
			return false
		}
	}
	return false
}

func Max(i,j int) int {
	if i < j {
		return j
	}
	return i
}

func CardsToNums(hand string) []int {
	nums := []int{}
	for _, c := range hand {
		if (c == 'A') {
			nums = append(nums, 14)
		} else if (c == 'K') {
			nums = append(nums, 13)
		} else if (c == 'Q') {
			nums = append(nums, 12)
		} else if (c == 'J') {
			nums = append(nums, 1)
		} else if (c == 'T') {
			nums = append(nums, 10)
		} else {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println("Error: ", err)
			}
			nums = append(nums, n)
		}
	}
	return nums
}

func GradeHand(hand string) int{
	cards := SortString(hand)

	counts := map[rune]int{}

	for _, c := range cards {
		if _, ok := counts[c]; !ok {
			counts[c] = 0
		}
		counts[c] = counts[c] + 1
	}

	inverse := map[int]int{}

	for r, c := range counts {
		if r != 'J' {
			if _ , ok := inverse[c]; !ok {
				inverse[c] =  0
			}
			inverse[c] += 1
		}
	}

	jcount := 0
	if c , ok := counts['J']; ok {
		jcount = c
	}

	if jcount == 5 {
		return 7
	}

	grade := 0


	if _, ok := inverse[4]; ok {
		if jcount >= 1 {
			inverse[4+jcount] = 1
			inverse[4] -= 1
		}
	} else if _, ok := inverse[3]; ok {
		if jcount >= 1 {
			inverse[3+jcount] = 1
			inverse[3] -= 1
		}
	} else if _, ok := inverse[2]; ok {
		if jcount >= 1 {
			inverse[2+jcount] = 1
			inverse[2] -= 1
		}
	} else if _, ok := inverse[1]; ok {
		if jcount >= 1 {
			inverse[1+jcount] = 1
			inverse[1] -= 1
		}
	}

	// fmt.Printf("Hand %v Inverser counts: %v, jcount %v\n", hand, PrettyPrint(inverse), jcount)


	if _, ok := inverse[5]; ok {
		grade = 7
	} else if _, ok := inverse[4]; ok {
		grade = 6
	} else if _, ok := inverse[3]; ok {
		if c, okok := inverse[2]; okok && c >= 1{
			grade = 5
		} else {
			grade = 4
		}
	} else if c, ok := inverse[2]; ok{
		if c == 2 {
			grade = 3
		} else if c == 1 {
			grade = 2
		}
	} else if _, ok := inverse[1]; ok {
		grade = 1
	}
	return grade
}

func part1(input string){
	lines := strings.Split(input, "\n")
	hands := []string{}
	scores := []int{}

	for _ , l := range lines {
		fields := strings.Fields(l)
		if len(fields) == 0 {
			break
		}
		// fmt.Println(fields)
		hand := fields[0]
		score, _ := strconv.Atoi(fields[1])

		hands = append(hands, hand)
		scores = append(scores, score)
	}

	ranks := []Hand{}

	for i, h := range hands {
		grade := GradeHand(h)
		// fmt.Printf("Grade of hand %v is %v\n", h, grade)
		rankedHand := Hand{grade, CardsToNums(h), scores[i]}
		// fmt.Printf("Ranked Hand: %v\n", rankedHand)
		ranks = append(ranks, rankedHand)
	}

	ranks = SortRankedHands(ranks)
	sum := 0

	// fmt.Printf("Ranked Hands: %v\n", PrettyPrint(ranks))

	for i, h := range ranks {
		// fmt.Printf("h.Score * i+1 %v\n", h.Score*(i+1))
		sum = sum + ((i+1) * h.Score)
	}

	// fmt.Printf("Hands: %v Scores: %v\n", hands, scores)
	fmt.Println("Sum: ", sum)

}

func part2(input string){

}
