package main

import (
	"fmt"
	"time"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	startTime := time.Now()
	input, err := getFiles("example")
	if err != nil {
		fmt.Println(err)
		return
	}
	partTwo(input)
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed)
}

func partOne(input []string) {
	splits := 0
	locs := []int{}
	for i, v := range input[0] {
		if v == 'S' {
			locs = append(locs, i)
		}
	}
	for _, row := range input[1:] {
		newLocs := []int{}
		for _, loc := range locs {
			if row[loc] == '^' {
				splits++
				if !contains(newLocs, loc-1) {
					newLocs = append(newLocs, loc-1)
				}
				if !contains(newLocs, loc+1) {
					newLocs = append(newLocs, loc+1)
				}
			} else {
				if !contains(newLocs, loc) {
					newLocs = append(newLocs, loc)
				}
			}
		}
		locs = newLocs
	}
	fmt.Println(locs)
	fmt.Println(splits)
}

func contains(row []int, val int) bool {
	for _, v := range row {
		if v == val {
			return true
		}
	}
	return false
}

type locData struct {
	loc   int
	count int
}

func (l locData) locationInArray(row []locData) int {
	for i, v := range row {
		if v.loc == l.loc {
			return i
		}
	}
	return -1
}

func partTwo(input []string) {
	locs := []locData{}
	for i, v := range input[0] {
		if v == 'S' {
			newLoc := locData{loc: i, count: 1}
			locs = append(locs, newLoc)
		}
	}
	for _, row := range input[1:] {
		newLocs := []locData{}
		for _, loc := range locs {
			diffs := []int{0}
			if row[loc.loc] == '^' {
				diffs = []int{-1, 1}
			}
			for _, diff := range diffs {
				newLoc := locData{loc: loc.loc + diff, count: loc.count}
				containment := newLoc.locationInArray(newLocs)
				if containment == -1 {
					newLocs = append(newLocs, newLoc)
				} else {
					newLocs[containment].count += newLoc.count
				}
			}
		}
		locs = newLocs
	}
	total := 0
	for _, loc := range locs {
		total += loc.count
	}
	fmt.Println(total)
}

func getFiles(version string) ([]string, error) {
	file, err := utils.LoadFile("day_7", version)
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
