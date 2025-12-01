package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	partTwo()
}

func partOne() {
	total := 0
	data := getFiles()
	for _, line := range data {
		total += getLineTotal(line)
	}
	fmt.Println(total)
}

func partTwo() {
	total := 0
	data := getFiles()
	total += handleDos(strings.Join(data, ""))
	fmt.Println(total)
}

func handleDonts(input string) int {
	total := 0
	dos := strings.Split(input, "do()")
	for i, do := range dos {
		if i > 0 {
			total += handleDos(do)
		}
	}
	return total

}

func handleDos(input string) int {
	total := 0
	donts := strings.Split(input, "don't()")
	fmt.Println(donts)
	for i, dont := range donts {
		if i == 0 {
			total += getLineTotal(dont)
		} else {
			total += handleDonts(dont)
		}
	}
	return total
}

func getLineTotal(line string) int {
	total := 0
	split1 := strings.Split(line, "mul")
	left := "("
	//	right := ")"
	for _, v := range split1 {
		option := strings.Split(v, ")")
		if len(option) > 0 && len(option[0]) > 0 && option[0][0] == left[0] {
			candidate := strings.Split(option[0], "(")
			nums := strings.Split(candidate[1], ",")
			if len(nums) == 2 {
				first, err := strconv.Atoi(nums[0])
				if err != nil {
					continue
				}
				second, err := strconv.Atoi(nums[1])
				if err != nil {
					continue
				}
				total += first * second
			}
		}
	}
	return total
}

func checkRow(row []int) bool {
	sign := row[0] < row[1]
	for i, _ := range row {
		if i == len(row)-1 {
			return true
		}
		if (row[i] < row[i+1]) != sign {
			return false
		}
		if sign && ((row[i+1]-row[i] < 1) || (row[i+1]-row[i] > 3)) {
			return false
		} else if !sign && ((row[i]-row[i+1] < 1) || (row[i]-row[i+1] > 3)) {
			return false
		}
	}
	return true
}

func getFiles() []string {
	file, err := utils.LoadFile("./inputs/2024_day_3_real.txt")
	if err != nil {
		fmt.Println("error: ", err)
		return nil
	}
	return file
}
