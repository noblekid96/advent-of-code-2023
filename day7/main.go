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
			nums = append(nums, 11)
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

	if cards[0] == cards[4]{
		return 7 // Five of a kind
	} else if cards[0] == cards[3] || cards[1] == cards[4] {
		return 6 // Four of a kind
	} else if (cards[0] == cards[1] && cards[2] == cards[4]) || (cards[0] == cards[2] && cards[3] == cards[4]){
		return 5 // Full house
	} else if (cards[0] == cards[2] && cards[3] != cards[4]) || (cards[1] == cards[3] && cards[0] != cards[4]) || (cards[2] == cards[4] && cards[0] != cards[1]){
		return 4 // Three of a kind
	} else if (cards[0] == cards[1] && cards[2] == cards[3]) || (cards[0] == cards[1] && cards[3] == cards[4]) || (cards[1] == cards[2] && cards[3] == cards[4]){
		return 3 // Two pair
	} else if (cards[0] == cards[1] || cards[1] == cards[2] || cards[2] == cards[3] || cards[3] == cards[4]){
		return 2 // One pair
	} else {
		return 1
	}
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
		fmt.Println(fields)
		hand := fields[0]
		score, _ := strconv.Atoi(fields[1])

		hands = append(hands, hand)
		scores = append(scores, score)
	}

	ranks := []Hand{}

	for i, h := range hands {
		grade := GradeHand(h)
		fmt.Printf("Grade of hand %v is %v\n", h, grade)
		rankedHand := Hand{grade, CardsToNums(h), scores[i]}
		fmt.Printf("Ranked Hand: %v\n", rankedHand)
		ranks = append(ranks, Hand{grade, CardsToNums(h), scores[i]})
	}

	ranks = SortRankedHands(ranks)
	sum := 0

	fmt.Printf("Ranked Hands: %v\n", PrettyPrint(ranks))

	for i, h := range ranks {
		// fmt.Printf("h.Score * i+1 %v\n", h.Score*(i+1))
		sum = sum + ((i+1) * h.Score)
	}

	// fmt.Printf("Hands: %v Scores: %v\n", hands, scores)
	fmt.Println("Sum: ", sum)

}

func part2(input string){

}
