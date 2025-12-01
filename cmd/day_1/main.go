package main

import (
	"fmt"
	"strconv"

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

func partOne(input []string) {
	count := 0
	val := 50
	for _, v := range input {
		if len(v) > 0 {
			distance, _ := integerPart(v)
			if startsWithL(v) {
				val -= distance
				val %= 100
				if val == 0 {
					count += 1
				}
			} else {
				val += distance
				val %= 100
				if val == 0 {
					count += 1
				}
			}
		}
	}
	fmt.Println(count)
}

func partTwo(input []string) {
	count := 0
	val := 50
	for _, v := range input {
		if len(v) > 0 {
			oldval := val
			distance, _ := integerPart(v)
			if startsWithL(v) {
				val -= distance
				if val == 0 {
					count += 1
				} else if oldval > 0 && val < 0 {
					count += 1
				}
				if val <= -100 {
					count += val / -100
				}
				val %= 100
			} else {
				val += distance
				if val == 0 {
					count += 1
				} else if oldval < 0 && val > 0 {
					count += 1
				}
				if val >= 100 {
					count += val / 100
				}
				val %= 100
			}
		}
	}
	fmt.Println(count)
}

func startsWithL(input string) bool {
	runes := []rune(input)
	return runes[0] == 'L'
}

func integerPart(input string) (int, error) {
	runes := []rune(input)
	substring := string(runes[1:])
	result, err := strconv.Atoi(substring)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func getFiles(version string) ([]string, error) {
	file, err := utils.LoadFile("day_1", version)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// parsed, err := utils.GetNumbers(file, "   ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, nil
	// }
	return file, nil
}
