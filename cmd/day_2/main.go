package main

import (
	"fmt"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	exampleInput, err := getFiles("example")
	if err != nil {
		fmt.Println(err)
		return
	}
	realInput, err := getFiles("real")
	if err != nil {
		fmt.Println(err)
		return
	}
	partOne(exampleInput)
	//partTwo(exampleInput)
	partOne(realInput)
	//partTwo(realInput)
}

func partOne(input []string) {
	fmt.Println(input[0])
}

func partTwo(input []string) {
	fmt.Println(input[0])
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
