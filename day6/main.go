package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func LowerBound(arr []int, target int) int {
    low, high := 0, len(arr)
    for low < high {
        mid := low + (high-low)/2
        if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return low
}

func UpperBound(arr []int, target int) int {
    low, high := 0, len(arr)
    for low < high {
        mid := low + (high-low)/2
        if arr[mid] <= target {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return low
}

func RaceLower(time int, distance int) int {
	options := []int{}
	for i := 0; i <= time; i ++ {
		options = append(options, i)
	}

	low, high := 0, len(options)
	for low < high {
		mid := low + (high-low)/2
		if options[mid]*(time-options[mid]) < distance+1 {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func RaceHigher(time int, distance int) int {
	options := []int{}
	for i := 0; i <= time; i ++ {
		options = append(options, i)
	}

	low, high := 0, len(options)
	for low < high {
		mid := low + (high-low)/2
		if options[mid]*(time-options[mid]) > distance {
			low = mid+1
		} else {
			high = mid
		}
	}
	return low-1
}


func part1(input string){
	lines := strings.Split(input, "\n")
	raceDurationStrings := strings.Fields(lines[0])[1:]
	raceDistanceStrings := strings.Fields(lines[1])[1:]
	n := len(raceDistanceStrings)
	ways := 1

	for i := 0 ; i < n; i ++ {
		raceDuration, _ := strconv.Atoi(raceDurationStrings[i])
		raceDistance, _ := strconv.Atoi(raceDistanceStrings[i])
		lower := RaceLower(raceDuration, raceDistance)
		higher := RaceHigher(raceDuration, raceDistance)
		fmt.Printf("raceDuration lower: %v raceDuration higher: %v", lower, higher)
		ways *= (higher - lower + 1)
	}

	fmt.Printf("raceDuration: %v raceDistance: %v ways: %v\n", PrettyPrint(raceDurationStrings), PrettyPrint(raceDistanceStrings), ways)
}

func part2(input string){
	lines := strings.Split(input, "\n")
	raceDurationStrings := strings.Fields(lines[0])[1:]
	raceDistanceStrings := strings.Fields(lines[1])[1:]

	raceDuration, _ := strconv.Atoi(strings.Join(raceDurationStrings, ""))
	raceDistance, _ := strconv.Atoi(strings.Join(raceDistanceStrings, ""))
	// n := len(raceDistanceStrings)
	ways := 1

	// for i := 0 ; i < n; i ++ {
	// raceDuration, _ := strconv.Atoi(raceDurationStrings[i])
	// raceDistance, _ := strconv.Atoi(raceDistanceStrings[i])
	lower := RaceLower(raceDuration, raceDistance)
	higher := RaceHigher(raceDuration, raceDistance)
	fmt.Printf("raceDuration lower: %v raceDuration higher: %v", lower, higher)
	ways *= (higher - lower + 1)
	// }

	fmt.Printf("raceDuration: %v raceDistance: %v ways: %v\n", PrettyPrint(raceDurationStrings), PrettyPrint(raceDistanceStrings), ways)
}
