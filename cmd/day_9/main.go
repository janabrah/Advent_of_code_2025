package main

import (
	"fmt"
	"time"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	input, err := getFiles("real")
	if err != nil {
		fmt.Println(err)
		return
	}
	partTwo(input)
}

func partOne(input [][]int) {
	fmt.Println(input)
	fmt.Println(int(findBiggest(input)))
}

func partTwo(input [][]int) {
	startTime := time.Now()
	fmt.Println(input)
	fmt.Println(getVorticity(input))
	maxX, maxY := findMaxs(input)
	fmt.Println(maxX, maxY)
	fmt.Println(input)
	grid := makeGrid(input, maxX, maxY)
	elapsed := time.Since(startTime)
	fmt.Println("made grid,", "Time taken:", elapsed)
	//	utils.PrettyPrint2DRune(grid, 0)
	grid = colorGrid(grid)
	elapsed = time.Since(startTime)
	fmt.Println("colored grid", "Time taken:", elapsed)
	//	utils.PrettyPrint2DRune(grid, 0)
	fmt.Println(findBiggest2(input, grid))
}

func getVorticity(input [][]int) int {
	// Not used after all but it was fun
	total := 0
	vorticity := makeVorticityMap()
	direction := getDirection(input[0], input[1])
	for i := range input[1 : len(input)-1] {
		newDirection := getDirection(input[i+1], input[i+2])
		turn := vorticity[direction][newDirection]
		fmt.Println(turn)
		total += turn
		direction = newDirection
	}
	return total
}

func makeVorticityMap() map[rune]map[rune]int {
	vorticity := map[rune]map[rune]int{}
	vorticity['l'] = map[rune]int{}
	vorticity['r'] = map[rune]int{}
	vorticity['d'] = map[rune]int{}
	vorticity['u'] = map[rune]int{}
	vorticity['l']['d'] = 1
	vorticity['l']['u'] = -1
	vorticity['r']['u'] = 1
	vorticity['r']['d'] = -1
	vorticity['u']['l'] = 1
	vorticity['u']['r'] = -1
	vorticity['d']['r'] = 1
	vorticity['d']['l'] = -1
	return vorticity
}

func getDirection(source []int, target []int) rune {
	if source[0] > target[0] {
		return 'u'
	} else if source[0] < target[0] {
		return 'd'
	} else if source[1] > target[1] {
		return 'l'
	} else if source[1] < target[1] {
		return 'r'
	}
	panic("oh no")
}

func makeGrid(input [][]int, maxX int, maxY int) [][]rune {
	grid := [][]rune{}
	for range maxX + 2 {
		grid = append(grid, make([]rune, maxY+1))
	}
	input = append(input, input[0])
	for i, v := range input[:len(input)-1] {
		if input[i][0] > input[i+1][0] {
			for j := input[i+1][0]; j <= input[i][0]; j++ {
				grid[j][v[1]] = 'X'
			}
		}
		if input[i][0] < input[i+1][0] {
			for j := input[i+1][0]; j >= input[i][0]; j-- {
				grid[j][v[1]] = 'X'
			}
		}
		if input[i][1] > input[i+1][1] {
			for j := input[i+1][1]; j <= input[i][1]; j++ {
				grid[v[0]][j] = 'X'
			}
		}
		if input[i][1] < input[i+1][1] {
			for j := input[i+1][1]; j >= input[i][1]; j-- {
				grid[v[0]][j] = 'X'
			}
		}
		grid[v[0]][v[1]] = '#'
	}
	grid[input[0][0]][input[0][1]] = '#'
	return grid
}

func colorGrid(input [][]rune) [][]rune {
	for i := range input {
		crossings := 0
		for j := range input[i] {
			if input[i][j] == 'X' || input[i][j] == '#' {
				if input[i+1][j] == 'X' || input[i+1][j] == '#' {
					crossings++
				}
			} else {
				if crossings%2 == 1 {
					input[i][j] = 'X'
				}
			}
		}
	}
	return input
}

func findBiggest(input [][]int) int {
	biggest := 0
	for i, v := range input {
		if i%100 == 0 {
			fmt.Println(i, len(input))
		}
		for _, u := range input[i+1:] {
			left, right, up, down := 0, 0, 0, 0
			if v[0] > u[0] {
				left = u[0]
				right = v[0]
			} else {
				left = v[0]
				right = u[0]
			}
			if v[1] < u[1] {
				up = v[1]
				down = u[1]
			} else {
				up = u[1]
				down = v[1]
			}
			size := (right - left + 1) * (down - up + 1)
			if size > biggest {
				biggest = size
			}
		}
	}
	return biggest
}

func findBiggest2(input [][]int, grid [][]rune) int {
	startTime := time.Now()

	biggest := 0
	for i, v := range input {
		if i%100 == 0 {
			fmt.Println(i, len(input))
		}
		for _, u := range input[i+1:] {
			left, right, up, down := 0, 0, 0, 0
			if v[0] > u[0] {
				down = v[0]
				up = u[0]
			} else {
				down = u[0]
				up = v[0]
			}
			if v[1] < u[1] {
				left = v[1]
				right = u[1]
			} else {
				left = u[1]
				right = v[1]
			}
			size := (right - left + 1) * (down - up + 1)
			if size > biggest {
				if checkExterior(u, v, grid) && checkInterior(u, v, grid) {
					biggest = size
					elapsed := time.Since(startTime)
					fmt.Println(biggest, "Time taken:", elapsed)
				}
			}
		}
	}
	return biggest
}

func checkInterior(u []int, v []int, grid [][]rune) bool {
	left, right, up, down := 0, 0, 0, 0
	if v[0] > u[0] {
		down = v[0]
		up = u[0]
	} else {
		down = u[0]
		up = v[0]
	}
	if v[1] < u[1] {
		left = v[1]
		right = u[1]
	} else {
		left = u[1]
		right = v[1]
	}
	midH := (left + right) / 2
	midV := (up + down) / 2
	if grid[midV][midH] != 'X' && grid[midV][midH] != '#' {
		// Seems like checking the center is most likely to find a hole, so starting there
		return false
	}
	for i := up; i <= down; i++ {
		for j := left; j <= right; j++ {
			if grid[i][j] != 'X' && grid[i][j] != '#' {
				return false
			}
		}
	}
	return true
}

func checkExterior(u []int, v []int, grid [][]rune) bool {
	left, right, up, down := 0, 0, 0, 0
	if v[0] > u[0] {
		down = v[0]
		up = u[0]
	} else {
		down = u[0]
		up = v[0]
	}
	if v[1] < u[1] {
		left = v[1]
		right = u[1]
	} else {
		left = u[1]
		right = v[1]
	}
	midH := (left + right) / 2
	midV := (up + down) / 2
	if grid[midV][midH] != 'X' && grid[midV][midH] != '#' {
		// Seems like checking the center is most likely to find a hole, so starting there
		return false
	}
	for i := up; i <= down; i++ {
		if grid[i][left] != 'X' && grid[i][left] != '#' {
			return false
		}
		if grid[i][right] != 'X' && grid[i][right] != '#' {
			return false
		}
	}
	for j := left; j <= right; j++ {
		if grid[up][j] != 'X' && grid[up][j] != '#' {
			return false
		}
		if grid[down][j] != 'X' && grid[down][j] != '#' {
			return false
		}
	}
	return true
}

func findMaxs(input [][]int) (int, int) {
	maxX := 0
	maxY := 0
	for _, v := range input {
		if v[0] > maxX {
			maxX = v[0]
		}
		if v[1] > maxY {
			maxY = v[1]
		}
	}
	return maxX, maxY
}

func getFiles(version string) ([][]int, error) {
	file, err := utils.LoadFile("day_9", version)
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
