package utils

import (
	"os"
	"strconv"
	"strings"
)

func LoadFile(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	stringFile := string(file)
	result := strings.Split(stringFile, "\n")
	return result, nil
}

func GetNumbers(input []string, divider string) ([][]int, error) {
	result := make([][]int, 0)
	for _, v := range input {
		row := make([]int, 0, 2)
		split := strings.Split(v, divider)
		left, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		row = append(row, left)
		right, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		row = append(row, right)
		result = append(result, row)
	}
	return result, nil
}

// It seemed fun to implement my own merge sort since I wasn't sure what go's builtin was
func SortIntArray(input []int) []int {
	result := make([]int, 0, len(input))
	if len(input) == 0 || len(input) == 1 {
		return input
	}
	split := len(input) / 2
	left := SortIntArray(input[:split])
	right := SortIntArray(input[split:])
	l, r := 0, 0
	for l < len(left) || r < len(right) {
		if l == len(left) {
			result = append(result, right[r])
			r++
		} else if r == len(right) {
			result = append(result, left[l])
			l++
		} else if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	return result
}
