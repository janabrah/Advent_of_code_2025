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
			distance, err := integerPart(v)
			if err != nil {
				fmt.Println("There was en error: ", err)
				return
			}
			if v[0] == 'L' {
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
			distance, err := integerPart(v)
			if err != nil {
				fmt.Println("There was en error: ", err)
				return
			}
			if v[0] == 'L' {
				val -= distance
				if val == 0 {
					count += 1
				} else if oldval > 0 && val < 0 {
					count += 1
				}
				count += val / -100
				val %= 100
			} else {
				val += distance
				if val == 0 {
					count += 1
				} else if oldval < 0 && val > 0 {
					count += 1
				}
				count += val / 100
				val %= 100
			}
		}
	}
	fmt.Println(count)
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
	return file, nil
}
