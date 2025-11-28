package main

import (
	"fmt"
	"math"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	partTwo()
}

func partOne() {
	total := 0
	left, right := getFiles()
	sortedLeft := utils.SortIntArray(left)
	sortedRight := utils.SortIntArray(right)
	for i := range sortedLeft {
		total += int(math.Abs(float64(sortedLeft[i] - sortedRight[i])))
	}
	fmt.Println(total)
}

func partTwo() {
	total := 0
	left, right := getFiles()
	rightCounts := countAppearances(right)
	for i := range left {
		total += left[i] * rightCounts[left[i]]
	}
	fmt.Println(total)
}

func countAppearances(input []int) map[int]int {
	result := map[int]int{}
	for _, v := range input {
		result[v] = result[v] + 1
	}
	fmt.Println(result)
	return result
}

func getFiles() ([]int, []int) {
	file, err := utils.LoadFile("./inputs/2024_day_1_real.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	parsed, err := utils.GetNumbers(file, "   ")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	left, right := make([]int, 0, len(parsed)), make([]int, 0, len(parsed))
	for _, v := range parsed {
		left = append(left, v[0])
		right = append(right, v[1])
	}
	return left, right
}
