package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	startTime := time.Now()
	input, err := getFiles("real")
	if err != nil {
		fmt.Println(err)
		return
	}
	partTwo(input)
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed)
}

func partOne(input []string) {
	total := 0
	ranges := make([][]int, 0)
	lasti := 0
	for i, v := range input {
		if v == "" {
			break
		}
		lasti = i
		myRange := strings.Split(v, "-")
		beginning, _ := strconv.Atoi(myRange[0])
		end, _ := strconv.Atoi(myRange[1])
		newRange := []int{beginning, end}
		ranges = append(ranges, newRange)
	}
	for _, v := range input[lasti+2:] {
		for _, r := range ranges {
			val, _ := strconv.Atoi(v)
			if val >= r[0] && val <= r[1] {
				total++
				break
			}
		}
	}
	fmt.Println(total)
}

func partTwo(input []string) {
	ranges := make([][]int, 0)
	for _, v := range input {
		if v == "" {
			break
		}
		myRange := strings.Split(v, "-")
		beginning, _ := strconv.Atoi(myRange[0])
		end, _ := strconv.Atoi(myRange[1])
		newRange := []int{beginning, end}
		ranges = append(ranges, newRange)
	}
	ranges = SortRangesArray(ranges)
	total := 0
	for len(ranges) > 0 {
		if len(ranges) == 1 || ranges[0][1] < ranges[1][0] {
			total += ranges[0][1] - ranges[0][0] + 1
		} else {
			ranges[1][0] = ranges[0][0]
			if ranges[0][1] > ranges[1][1] {
				ranges[1][1] = ranges[0][1]
			}
		}
		ranges = ranges[1:]
	}
	fmt.Println(total)
}

func getFiles(version string) ([]string, error) {
	file, err := utils.LoadFile("day_5", version)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// parsed, err := utils.GetNumbers(file, "")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, nil
	// }
	return file, nil
}

// It seemed fun to implement my own merge sort since I wasn't sure what go's builtin was
func SortRangesArray(input [][]int) [][]int {
	result := make([][]int, 0, len(input))
	if len(input) == 0 || len(input) == 1 {
		return input
	}
	split := len(input) / 2
	left := SortRangesArray(input[:split])
	right := SortRangesArray(input[split:])
	l, r := 0, 0
	for l < len(left) || r < len(right) {
		if l == len(left) {
			result = append(result, right[r])
			r++
		} else if r == len(right) {
			result = append(result, left[l])
			l++
		} else if left[l][0] < right[r][0] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	return result
}
