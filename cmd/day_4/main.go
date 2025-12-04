package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	input, err := getFiles("real")
	if err != nil {
		fmt.Println(err)
		return
	}

	startTime := time.Now()
	partTwoBetter(input)
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed)
}

func partOne(input []string) {
	myMap := make([][]string, 0)
	for _, v := range input {
		myMap = append(myMap, strings.Split(v, ""))
	}
	sides := []int{-1, 0, 1}
	total := 0
	for i, _ := range myMap {
		for j, c := range myMap[i] {
			if c == "@" {
				neighbors := 0
				for _, iplus := range sides {
					for _, jplus := range sides {
						if i+iplus >= 0 && i+iplus < len(myMap) && j+jplus >= 0 && j+jplus < len(myMap[0]) {
							if myMap[i+iplus][j+jplus] == "@" {
								neighbors++
							}
						}
					}
				}
				if neighbors <= 4 {
					total++
				}

			}
		}
	}
	fmt.Println(total)
}

func partTwo(input []string) {
	myMap := make([][]string, 0)
	for _, v := range input {
		myMap = append(myMap, strings.Split(v, ""))
	}
	sides := []int{-1, 0, 1}
	myIntMap := make([][]int, 0)
	for i, _ := range myMap {
		myIntMap = append(myIntMap, make([]int, 0))
		for j, c := range myMap[i] {
			neighbors := 0
			myIntMap[i] = append(myIntMap[i], -1)
			if c == "@" {
				for _, iplus := range sides {
					for _, jplus := range sides {
						if i+iplus >= 0 && i+iplus < len(myMap) && j+jplus >= 0 && j+jplus < len(myMap[0]) && (iplus != 0 || jplus != 0) {
							if myMap[i+iplus][j+jplus] == "@" {
								neighbors++
							}
						}
					}
				}
				myIntMap[i][j] = neighbors
			}
		}
	}
	total := -1
	removed := 1
	for removed > 0 {
		total += removed
		myIntMap, removed = removePaper(myIntMap)
	}
	fmt.Println(total)
}

func removePaper(myIntMap [][]int) ([][]int, int) {
	sides := []int{-1, 0, 1}
	removed := 0
	for i, _ := range myIntMap {
		for j, c := range myIntMap[i] {
			if c != -1 && c < 4 {
				removed++
				myIntMap[i][j] = -1
				for _, iplus := range sides {
					for _, jplus := range sides {
						if i+iplus >= 0 && i+iplus < len(myIntMap) && j+jplus >= 0 && j+jplus < len(myIntMap[0]) && (iplus != 0 || jplus != 0) {
							if myIntMap[i+iplus][j+jplus] != -1 {
								myIntMap[i+iplus][j+jplus]--
							}
						}
					}
				}
			}
		}
	}
	return myIntMap, removed
}

type Stack struct {
	items [][]int
}

func (s *Stack) Push(data []int) {
	s.items = append(s.items, data)
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() []int {
	if s.IsEmpty() {
		return nil
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func partTwoBetter(input []string) {
	myMap := make([][]string, 0)
	for _, v := range input {
		myMap = append(myMap, strings.Split(v, ""))
	}
	sides := []int{-1, 0, 1}
	myIntMap := make([][]int, 0)
	locationStack := Stack{}
	total := 0
	for i, _ := range myMap {
		myIntMap = append(myIntMap, make([]int, 0))
		for j, c := range myMap[i] {
			neighbors := 0
			myIntMap[i] = append(myIntMap[i], -1)
			if c == "@" {
				for _, iplus := range sides {
					for _, jplus := range sides {
						if i+iplus >= 0 && i+iplus < len(myMap) && j+jplus >= 0 && j+jplus < len(myMap[0]) && (iplus != 0 || jplus != 0) {
							if myMap[i+iplus][j+jplus] == "@" {
								neighbors++
							}
						}
					}
				}
				myIntMap[i][j] = neighbors
				newLoc := []int{i, j}
				if neighbors < 4 {
					locationStack.Push(newLoc)
				}
			}
		}
	}
	for len(locationStack.items) > 0 {
		loc := locationStack.Pop()
		i := loc[0]
		j := loc[1]
		if myIntMap[i][j] > -1 && myIntMap[i][j] < 4 {
			total += 1
			myIntMap[i][j] = -1
			for _, iplus := range sides {
				for _, jplus := range sides {
					if i+iplus >= 0 && i+iplus < len(myIntMap) && j+jplus >= 0 && j+jplus < len(myIntMap[0]) && (iplus != 0 || jplus != 0) {
						if myIntMap[i+iplus][j+jplus] != -1 {
							myIntMap[i+iplus][j+jplus]--
							newLoc := []int{i + iplus, j + jplus}
							if myIntMap[i+iplus][j+jplus] == 3 {
								locationStack.Push(newLoc)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(total)
}

func getFiles(version string) ([]string, error) {
	file, err := utils.LoadFile("day_4", version)
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
