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
	fmt.Println(partTwoFast(input))
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed)
}

func partOne(input [][]int) int {
	connections := map[int]map[int]bool{}
	cycles := map[int]int{}
	groups := []map[int]bool{}
	for i := range input {
		connections[i] = map[int]bool{}
	}
	for range 1000 {
		i, j := findClosest(input, connections)
		connections[i][j] = true
		insertCycle(cycles, i, j)
		groups = groupInsert(groups, i, j)
	}
	sizes := map[int]int{}
	for i := range input {
		ind := len(input) - i - 1
		sizes[ind] = sizes[cycles[ind]] + 1
		if cycles[ind] != 0 {
			sizes[cycles[ind]] = 0
		}
	}
	sizes2 := []int{}
	for _, g := range groups {
		sizes2 = append(sizes2, len(g))
	}
	sizes2 = utils.SortIntArray(sizes2)
	fmt.Println(sizes2)
	largests := []int{0, 0, 0}
	for i := range input {
		if sizes[i] > largests[0] {
			largests[2] = largests[1]
			largests[1] = largests[0]
			largests[0] = sizes[i]
		} else if sizes[i] > largests[1] {
			largests[2] = largests[1]
			largests[1] = sizes[i]
		} else if sizes[i] > largests[2] {
			largests[2] = sizes[i]
		}

	}
	fmt.Println(largests)
	fmt.Println(largests[2] * largests[1] * largests[0])
	fmt.Println(sizes2[len(sizes2)-1] * sizes2[len(sizes2)-2] * sizes2[len(sizes2)-3])
	return largests[2] * largests[1] * largests[0]
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
		if c%1000 == 0 {
			fmt.Println(c, len(groups[0]))
		}
		if len(groups[0]) == len(input) {
			return input[i][0] * input[j][0]
		}
	}
	return 0
}

func partTwoFast(input [][]int) int {
	distances := findDstances(input)
	distances = sortRangesArray(distances)
	connections := map[int]map[int]bool{}
	groups := []map[int]bool{}
	for i := range input {
		connections[i] = map[int]bool{}
	}
	for c := 0; true; c++ {
		i, j := findClosestFast(distances, connections)
		connections[i][j] = true
		groups = groupInsert(groups, i, j)
		if len(groups[0]) == len(input) {
			return input[i][0] * input[j][0]
		}
	}
	return 0
}

func groupInsert(groups []map[int]bool, source int, target int) []map[int]bool {
	hits := []int{}
	for g, group := range groups {
		for i := range group {
			if i == source || i == target {
				if len(hits) == 0 || (len(hits) == 1 && hits[0] != g) || (hits[0] != g && hits[1] != g) {
					hits = append(hits, g)
				}
				groups[g][source] = true
				groups[g][target] = true
			}
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
	newGroup := map[int]bool{}
	newGroup[source] = true
	newGroup[target] = true
	groups = append(groups, newGroup)
	return groups
}

func insertCycle(cycles map[int]int, source int, target int) map[int]int {
	if source == target {
		return cycles
	}
	if pointsAt(cycles, source, target) {
		return cycles
	}
	sourceSource := 0
	targetSource := 0
	for i := range 1000 {
		if cycles[i] == target {
			targetSource = i
		}
		if cycles[i] == source {
			sourceSource = i
		}
	}
	if targetSource > 0 {
		if sourceSource == 0 {
			if targetSource < source {
				cycles[targetSource] = source
				cycles = insertCycle(cycles, source, target)
			} else {
				cycles = insertCycle(cycles, source, targetSource)
			}
		} else {
			small := sourceSource
			midsmall := targetSource
			midbig := source
			big := target
			if targetSource > source {
				midsmall = source
				midbig = targetSource
			} else if sourceSource > targetSource {
				small = targetSource
				midsmall = sourceSource
			}
			cycles[small] = midsmall
			cycles[midsmall] = midbig
			cycles = insertCycle(cycles, midbig, big)
		}
		return cycles
	}
	if cycles[source] == 0 {
		cycles[source] = target
		return cycles
	}
	currentTarget := cycles[source]
	if currentTarget < target {
		cycles = insertCycle(cycles, currentTarget, target)
	} else {
		cycles[source] = target
		cycles = insertCycle(cycles, target, currentTarget)
	}
	return cycles
}

func pointsAt(cycles map[int]int, source int, target int) bool {
	if cycles[source] == target {
		return true
	}
	if cycles[source] == 0 {
		return false
	}
	return pointsAt(cycles, cycles[source], target)
}

func connect(left int, right int, connections map[int]map[int]bool) map[int]map[int]bool {
	connections[left][right] = true
	connections[right][left] = true
	for i := range connections[left] {
		connections[right][i] = true
		connections[i][right] = true
	}
	for i := range connections[right] {
		connections[left][i] = true
		connections[i][left] = true
	}
	return connections
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

func findClosestFast(options [][]int, exclusions map[int]map[int]bool) (int, int) {
	for _, option := range options {
		if !exclusions[option[1]][option[2]] {
			return option[1], option[2]
		}
	}
	return -1, -1
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
func sortRangesArray(input [][]int) [][]int {
	result := make([][]int, 0, len(input))
	if len(input) == 0 || len(input) == 1 {
		return input
	}
	split := len(input) / 2
	left := sortRangesArray(input[:split])
	right := sortRangesArray(input[split:])
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
