package main

import (
	"fmt"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	partTwo()
}

func partOne() {
	count := 0
	data := getFiles()
	for _, row := range data {
		if checkRow(row) {
			count++
		}
	}
	fmt.Println(count)
}

func partTwo() {
	count := 0
	data := getFiles()
	for _, row := range data {
		if checkRow(row) {
			count++
		} else {
			for i, _ := range row {
				newRow := []int{}
				newRow = append(newRow, row[:i]...)
				newRow = append(newRow, row[i+1:]...)
				if checkRow(newRow) {
					count++
					break
				}
			}
		}
	}
	fmt.Println(count)
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

func getFiles() [][]int {
	file, err := utils.LoadFile("./inputs/2024_day_2_real.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	utils.PrettyPrint1D(file)
	parsed, err := utils.GetNumbers(file, " ")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	utils.PrettyPrint2DInt(parsed, "")
	return parsed
}
