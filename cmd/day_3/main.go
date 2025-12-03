package main

import (
	"fmt"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	input, err := getFiles("real")
	if err != nil {
		fmt.Println(err)
		return
	}
	partTwo(input)
}

func partOne(input [][]int) {
	total := 0
	for _, v := range input {
		total += findMaxJoltagePartOne(v)
	}
	fmt.Println(total)
}

func partTwo(input [][]int) {
	total := 0
	for _, v := range input {
		total += findMaxJoltagePartTwo(v)
	}
	fmt.Println(total)
}

func findMaxJoltagePartOne(input []int) int {
	firstDigit := 0
	firstDigitIndex := -1
	for i, v := range input[:len(input)-1] {
		if v > firstDigit {
			firstDigit = v
			firstDigitIndex = i
		}
	}
	secondDigit := 0
	for i, v := range input[firstDigitIndex+1:] {
		if i+firstDigitIndex+1 != firstDigitIndex && v > secondDigit {
			secondDigit = v
		}
	}
	fmt.Println(firstDigit, secondDigit)
	return 10*firstDigit + secondDigit
}

func findMaxJoltagePartTwo(input []int) int {
	total := 0
	digits := make([]int, 0, 12)
	digitIndices := make([]int, 1, 13)
	digitIndices[0] = -1
	for j := 0; j < 12; j++ {
		digits = append(digits, 0)
		digitIndices = append(digitIndices, 0)
		for i, v := range input[digitIndices[j]+1 : len(input)-(11-j)] {
			if v > digits[j] {
				digits[j] = v
				digitIndices[j+1] = i + digitIndices[j] + 1
			}
		}
		// I didn't know the syntax for exponents in Go so this happened
		total = 10*total + digits[j]
	}
	fmt.Println(total)
	return total
}

func getFiles(version string) ([][]int, error) {
	file, err := utils.LoadFile("day_3", version)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	parsed, err := utils.GetNumbers(file, "")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return parsed, nil
}
