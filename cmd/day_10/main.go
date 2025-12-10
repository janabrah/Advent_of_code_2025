package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

var best int = 100000

func main() {
	// startTime := time.Now()
	input, err := getFiles("real")
	if err != nil {
		fmt.Println(err)
		return
	}
	partTwo(input)
	// elapsed := time.Since(startTime)
	// fmt.Println("Time taken:", elapsed)
}

func partOne(input []panel) {
	fmt.Println(input)
	total := 0
	for _, row := range input {
		presses := recursiveSolve(row.goal, row.buttonGroups, [][]int{}, make([][]int, 100000))
		fmt.Println(presses)
		total += len(presses)
	}
	fmt.Println(total)
}

func partTwo(input []panel) {
	startTime := time.Now()
	fmt.Println(input)
	total := 0
	for i, row := range input {
		best = 100000
		recursiveSolveTwo(row.joltages, row.buttonGroups, 0, 100000)
		//fmt.Println(presses)
		total += best
		elapsed := time.Since(startTime)
		fmt.Println(i, "Time taken:", elapsed)
	}
	fmt.Println(total)
}

func recursiveSolve(goal []bool, buttonGroups [][]int, presses [][]int, infinity [][]int) [][]int {
	best := infinity
	if isDone(goal) {
		return presses
	}
	if len(buttonGroups) == 0 {
		return infinity
	}
	newPresses := append(presses, buttonGroups[0])
	if len(buttonGroups) == 1 {
		return recursiveSolve(apply(goal, buttonGroups[0]), [][]int{}, newPresses, infinity)
	}
	test := recursiveSolve(apply(goal, buttonGroups[0]), buttonGroups[1:], newPresses, infinity)
	if len(test) < len(best) {
		best = test
	}
	test = recursiveSolve(goal, buttonGroups[1:], presses, infinity)
	if len(test) < len(best) {
		best = test
	}

	return best
}

func isDone(input []bool) bool {
	for _, v := range input {
		if v {
			return false
		}
	}
	return true
}

func apply(lights []bool, buttons []int) []bool {
	newLights := make([]bool, 0, len(lights))
	newLights = append(newLights, lights...)
	for _, v := range buttons {
		newLights[v] = !lights[v]
	}
	return newLights
}

func recursiveSolveTwo(goal []int, buttonGroups [][]int, presses int, infinity int) {
	if presses >= best {
		return
	}
	if isDoneTwo(goal) {
		best = presses
	}
	if len(buttonGroups) == 0 || isPastTwo(goal) {
		return
	}
	recursiveSolveTwo(applyTwo(goal, buttonGroups[0]), buttonGroups, presses+1, infinity)
	recursiveSolveTwo(goal, buttonGroups[1:], presses, infinity)
}

func isDoneTwo(input []int) bool {
	for _, v := range input {
		if v != 0 {
			return false
		}
	}
	return true
}

func isPastTwo(input []int) bool {
	for _, v := range input {
		if v < 0 {
			return true
		}
	}
	return false
}

func applyTwo(lights []int, buttons []int) []int {
	newLights := make([]int, 0, len(lights))
	newLights = append(newLights, lights...)
	for _, v := range buttons {
		newLights[v] = lights[v] - 1
	}
	return newLights
}

type panel struct {
	goal         []bool
	buttonGroups [][]int
	joltages     []int
}

func getFiles(version string) ([]panel, error) {
	file, err := utils.LoadFile("day_10", version)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res := []panel{}
	for _, row := range file {
		data := strings.Split(row, " ")
		target := data[0]
		goal := []bool{}
		fmt.Println(target)
		for _, v := range target[1 : len(target)-1] {
			if v == '#' {
				goal = append(goal, true)
			} else {
				goal = append(goal, false)
			}
		}
		buttons := [][]int{}
		for _, v := range data[1 : len(data)-1] {
			bs := strings.Split(v[1:len(v)-1], ",")
			but := []int{}
			for _, b := range bs {
				x, err := strconv.Atoi(b)
				if err != nil {
					panic(err)
				}
				but = append(but, x)
			}
			buttons = append(buttons, but)
		}
		jolt := data[len(data)-1]
		js := strings.Split(jolt[1:len(jolt)-1], ",")
		joltage := []int{}
		for _, j := range js {
			jval, err := strconv.Atoi(j)
			if err != nil {
				panic(err)
			}
			joltage = append(joltage, jval)
		}
		res = append(res, panel{goal, buttons, joltage})
	}

	// parsed, err := utils.GetNumbers(file, "")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, nil
	// }
	return res, nil
}
