package main

import (
	"fmt"
	"time"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

var infinity int = 100000000000 // largest input i think is 99999,99999,99999

func main() {
	startTime := time.Now()
	input, err := getFiles("real")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(partOne(input))
	elapsed := time.Since(startTime)
	fmt.Println("Time taken for part one slow:", elapsed)

	startTime = time.Now()
	fmt.Println(partOneFast(input))
	elapsed = time.Since(startTime)
	fmt.Println("Time taken for part one fast:", elapsed)

	startTime = time.Now()
	fmt.Println(partTwo(input))
	elapsed = time.Since(startTime)
	fmt.Println("Time taken for part 2 slow:", elapsed)

	startTime = time.Now()
	fmt.Println(partTwoFast(input))
	elapsed = time.Since(startTime)
	fmt.Println("Time taken for part 2 fast:", elapsed)
}

func partOne(input [][]int) int {
	connections := map[int]map[int]bool{}
	groups := []map[int]bool{}
	for i := range input {
		connections[i] = map[int]bool{}
	}
	for range 1000 {
		i, j := findClosest(input, connections)
		connections[i][j] = true
		groups = groupInsert(groups, i, j)
	}
	sizes := []int{}
	for _, g := range groups {
		sizes = append(sizes, len(g))
	}
	sizes = utils.SortIntArray(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func partOneFast(input [][]int) int {
	distances := sortDistanceArray(findDstances(input))
	groups := []map[int]bool{}
	for i, d := range distances {
		if i == 1000 {
			break
		}
		groups = groupInsert(groups, d[1], d[2])
	}
	sizes := []int{}
	for _, g := range groups {
		sizes = append(sizes, len(g))
	}
	sizes = utils.SortIntArray(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func partTwo(input [][]int) int {
	connections := map[int]map[int]bool{}
	groups := []map[int]bool{}
	for i := range input {
		connections[i] = map[int]bool{}
	}
	for c := 0; true; c++ {
		i, j := findClosest(input, connections)
		connections[i][j] = true
		groups = groupInsert(groups, i, j)
		if len(groups[0]) == len(input) {
			return input[i][0] * input[j][0]
		}
	}
	return 0
}

func partTwoFast(input [][]int) int {
	distances := sortDistanceArray(findDstances(input))
	groups := []map[int]bool{}
	for _, d := range distances {
		groups = groupInsert(groups, d[1], d[2])
		if len(groups[0]) == len(input) {
			return input[d[1]][0] * input[d[2]][0]
		}
	}
	return 0
}

func groupInsert(groups []map[int]bool, source int, target int) []map[int]bool {
	hits := []int{}
	for g, group := range groups {
		if group[source] || group[target] {
			hits = append(hits, g)
			groups[g][source] = true
			groups[g][target] = true
		}
	}
	if len(hits) == 1 {
		return groups
	}
	if len(hits) == 2 {
		for i := range groups[hits[1]] {
			groups[hits[0]][i] = true
		}
		groups = append(groups[:hits[1]], groups[hits[1]+1:]...)
		return groups
	}
	if len(hits) > 2 {
		panic("oh no")
	}
	groups = append(groups, map[int]bool{source: true, target: true})
	return groups
}

func dist(p1 []int, p2 []int) int {
	total := 0
	for i := range p1 {
		total += (p1[i] - p2[i]) * (p1[i] - p2[i])
	}
	return total
}

func findDstances(options [][]int) [][]int {
	result := [][]int{}
	for i := range options[:len(options)-1] {
		for j := range options[i+1:] {
			distance := dist(options[i], options[j+i+1])
			row := []int{distance, i, j + i + 1}
			result = append(result, row)
		}
	}
	return result
}

func findClosest(options [][]int, exclusions map[int]map[int]bool) (int, int) {
	smallest := infinity
	left := -1
	right := -1
	for i := range options[:len(options)-1] {
		for j := range options[i+1:] {
			if !exclusions[i][j+i+1] {
				distance := dist(options[i], options[j+i+1])
				if distance < smallest {
					smallest = distance
					left = i
					right = j + i + 1
				}
			}
		}
	}
	return left, right
}

func getFiles(version string) ([][]int, error) {
	file, err := utils.LoadFile("day_8", version)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	parsed, err := utils.GetNumbers(file, ",")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return parsed, nil
}

// It seemed fun to implement my own merge sort since I wasn't sure what go's builtin was
func sortDistanceArray(input [][]int) [][]int {
	result := make([][]int, 0, len(input))
	if len(input) == 0 || len(input) == 1 {
		return input
	}
	split := len(input) / 2
	left := sortDistanceArray(input[:split])
	right := sortDistanceArray(input[split:])
	l, r := 0, 0
	for l < len(left) || r < len(right) {
		if l == len(left) {
			result = append(result, right[r])
			r++
		} else if r == len(right) {
			result = append(result, left[l])
			l++
		} else if left[l][0] < right[r][0] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	return result
}
