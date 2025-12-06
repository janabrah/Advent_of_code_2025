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
	fmt.Println(input)
	parsed := make([][]int, 0)
	operators := []string{}
	for i, v := range input {
		stringParse := parseStrings(strings.Trim(v, " "))
		if i == len(input)-1 {
			operators = stringParse
		} else {
			nums := convToInts(stringParse)
			parsed = append(parsed, nums)
		}
	}
	fmt.Println(parsed)
	fmt.Println(operators)
	total := 0
	ress := []int{}
	for _, v := range operators {
		if v == "+" {
			ress = append(ress, 0)
		} else {
			ress = append(ress, 1)
		}
	}
	for i := range parsed {
		for j := range parsed[0] {
			if operators[j] == "+" {
				ress[j] += parsed[i][j]
			} else {
				ress[j] *= parsed[i][j]
			}
		}
	}
	for _, v := range ress {
		total += v
	}
	fmt.Println(total)
}

func partTwo(input []string) {
	pivot := strings.Split(input[0], "")
	for _, v := range input[1 : len(input)-1] {
		runes := []rune(v)
		for j := range input[1] {
			pivot[j] = pivot[j] + string(runes[j])
		}
	}
	operators := parseStrings(strings.Trim(input[len(input)-1], " "))
	total := 0
	current := 0
	if operators[0] == "*" {
		current = 1
	}
	operatorIndex := 0
	for _, v := range pivot {
		if strings.Trim(v, " ") == "" {
			fmt.Println(current)
			total += current
			operatorIndex++
			if operators[operatorIndex] == "+" {
				current = 0
			} else {
				current = 1
			}
		} else {
			if operators[operatorIndex] == "+" {
				val, _ := strconv.Atoi(strings.Trim(v, " "))
				current += val
			} else {
				val, _ := strconv.Atoi(strings.Trim(v, " "))
				current *= val
			}
		}
	}
	total += current
	fmt.Println(total)
}

func parseStrings(input string) []string {
	for i := 0; i < len(input)-1; {
		if input[i+1] == ' ' && input[i] == ' ' {
			runes := []rune(input)
			runes = append(runes[:i], runes[i+1:]...)
			input = string(runes)
		} else {
			i++
		}
	}
	return strings.Split(input, " ")
}

func convToInts(numStrings []string) []int {
	nums := []int{}
	for _, v := range numStrings {
		conv, _ := strconv.Atoi(v)
		nums = append(nums, conv)
	}
	return nums
}

func getFiles(version string) ([]string, error) {
	file, err := utils.LoadFile("day_6", version)
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
