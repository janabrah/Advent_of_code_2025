package main

import (
	"fmt"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	// startTime := time.Now()
	input, err := getFiles("example")
	if err != nil {
		fmt.Println(err)
		return
	}
	partOne(input)
	// elapsed := time.Since(startTime)
	// fmt.Println("Time taken:", elapsed)
}

func partOne(input []string) {
	fmt.Println(input)
}

func partTwo(input []string) {
	fmt.Println(input)
}

func getFiles(version string) ([]string, error) {
	file, err := utils.LoadFile("day_8", version)
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
