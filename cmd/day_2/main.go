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
	entries := strings.Split(input[0], ",")
	total := 0
	for _, v := range entries {
		ends := strings.Split(v, "-")
		start, err := strconv.Atoi(ends[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		end, err := strconv.Atoi(ends[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		total += findInvalids(start, end)
	}
	fmt.Println(total)
}

func findInvalids(start int, end int) int {
	fmt.Println(start, end)
	total := 0
	for i := start; i <= end; i++ {
		if isInvalid(strconv.Itoa(i)) {
			total += i
		}
	}
	return total
}

func isInvalid(val string) bool {
	length := len(val)
	if length%2 != 0 {
		return false
	}
	if val[:length/2] == val[length/2:] {
		return true
	}
	return false
}

func findInvalids2(start int, end int) int {
	fmt.Println(start, end)
	total := 0
	for i := start; i <= end; i++ {
		if isInvalid2(strconv.Itoa(i)) {
			total += i
		}
	}
	return total
}

func isInvalid2(val string) bool {
	length := len(val)
	checked := false
loop:
	for i := 1; i <= length/2; i++ {
		if length%i == 0 {
			for j := 1; j < (length / i); j++ {
				if val[:i] != val[i*j:i*(j+1)] {
					continue loop
				}
			}
			checked = true
		}
	}
	return checked
}

func partTwo(input []string) {
	entries := strings.Split(input[0], ",")
	total := 0
	for _, v := range entries {
		ends := strings.Split(v, "-")
		start, err := strconv.Atoi(ends[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		end, err := strconv.Atoi(ends[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		total += findInvalids2(start, end)
	}
	fmt.Println(total)
}

func getFiles(version string) ([]string, error) {
	file, err := utils.LoadFile("day_2", version)
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
