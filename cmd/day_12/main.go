package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	shapes, targets, err := getFiles("real")
	if err != nil {
		fmt.Println(err)
		return
	}
	partOne(shapes, targets)
}

func partOne(shapes []shape, targets []target) {
	// I'm kinda annoyed that this turned out to be trivial, sigh,
	// I got tired of not having tried the problem so I figured I'd
	// start inspecting it and seeing if I might be able to be clever
	// by only looking at a few of them. Turns out they're all either
	// impossible or trivial, so here we are.
	big := 0
	veryBig := 0
	small := 0
	trivials := 0
	for _, target := range targets {
		space := checkSpace(shapes, target)
		if trivialCheckSpace(shapes, target) >= 0 {
			trivials++
		}
		fmt.Println(space)
		if space > 100 {
			veryBig++
			fmt.Println(target)
		} else if space >= 0 {
			big++
		} else {
			small++
		}
	}
	fmt.Println("veryBig", veryBig, "big", big, "small", small, "trivials", trivials, "trivials")
	return 1
}

func partTwo(input map[string][]string, bigSource string, destination string, avoid1 string, avoid2 string) int {
	fmt.Println(input)
	return 1
}

func checkSpace(shapes []shape, target target) int {
	total := target.height * target.width
	for k, v := range target.goals {
		total -= shapes[k].size * v
	}
	return total
}

func trivialCheckSpace(shapes []shape, target target) int {
	regions := target.height / 3 * target.width / 3
	for _, v := range target.goals {
		regions -= v
	}
	return regions
}

type shape struct {
	locs [][]bool
	size int
}

type target struct {
	width  int
	height int
	goals  map[int]int
}

func getFiles(version string) ([]shape, []target, error) {
	file, err := utils.LoadFile("day_12", version)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	fmt.Println(file)
	shapes := []shape{}
	targets := []target{}
	inShapes := true
	currentShape := []string{}
	for _, v := range file {
		if len(v) > 1 && v[1] == 'x' || len(v) > 2 && v[2] == 'x' || len(v) > 3 && v[3] == 'x' {
			inShapes = false
		}
		if v == "" {
			newShape := shape{[][]bool{}, 0}
			for _, r := range currentShape {
				row := []bool{}
				for _, c := range r {
					row = append(row, c == '#')
					if c == '#' {
						newShape.size++
					}
				}
				newShape.locs = append(newShape.locs, row)
			}
			toInsert := newShape
			shapes = append(shapes, toInsert)
			currentShape = []string{}

		} else if inShapes && v[1] != ':' {
			currentShape = append(currentShape, v)
		} else if !inShapes {
			sections := strings.Split(v, ": ")
			dims := strings.Split(sections[0], "x")
			newTarget := target{}
			newTarget.height, err = strconv.Atoi(dims[0])
			if err != nil {
				return nil, nil, err
			}
			newTarget.width, err = strconv.Atoi(dims[1])
			if err != nil {
				return nil, nil, err
			}
			newTarget.goals = map[int]int{}
			goalStrings := strings.Split(sections[1], " ")
			for i, goalStr := range goalStrings {
				goal, err := strconv.Atoi(goalStr)
				if err != nil {
					return nil, nil, err
				}
				newTarget.goals[i] = goal
			}
			targets = append(targets, newTarget)
		}
	}
	fmt.Println(shapes)
	fmt.Println(targets)
	return shapes, targets, nil
}
