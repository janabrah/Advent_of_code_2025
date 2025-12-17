package main

import (
	"fmt"
	"strings"

	"github.com/janabrah/Advent_of_code_2025/utils"
)

func main() {
	input, err := getFiles("real")
	if err != nil {
		fmt.Println(err)
		return
	}
	total := partTwo(copyMap(input), "svr", "out", "xxxx", "xxxx")
	nofft := partTwo(copyMap(input), "svr", "out", "fft", "xxxx")
	nodac := partTwo(copyMap(input), "svr", "out", "dac", "xxxx")
	noboth := partTwo(copyMap(input), "svr", "out", "fft", "dac")
	fmt.Println(total-nofft-nodac+noboth, total, nofft, nodac, noboth)
}

func copyMap(input map[string][]string) map[string][]string {
	res := map[string][]string{}
	for i := range input {
		res[i] = input[i]
	}
	return res
}

func partOne(input map[string][]string, source string, depth int) int {
	if source == "out" {
		return 1
	}
	if depth > 100 {
		return 0
	}
	total := 0
	for _, dest := range input[source] {
		total += partOne(input, dest, depth+1)
	}
	return total
}

func partTwo(input map[string][]string, bigSource string, destination string, avoid1 string, avoid2 string) int {
	locs := map[string]loc{}
	for v := range input {
		locs[v] = loc{false, false, false, 0}
	}
	locs[destination] = loc{true, false, false, 1}
	locs[avoid1] = loc{true, false, false, 0}
	locs[avoid2] = loc{true, false, false, 0}
	input[avoid1] = []string{}
	input[avoid2] = []string{}
	counter := 0
	for !locs[bigSource].pathOut {
		counter++
		if counter > 10000 {
			panic("too high")
		}
		for source := range input {
			if len(input[source]) > 0 {
				total := 0
				done := true
				for _, dest := range input[source] {
					total += locs[dest].distance

					if !locs[dest].pathOut {
						done = false
					}
				}
				if done {
					newStruct := loc{true, false, false, total}
					locs[source] = newStruct
					input[source] = []string{}
				}
			}
		}
	}
	fmt.Println(counter)
	return locs[bigSource].distance
}

func partTwoBad(input map[string][]string, source string, depth int, dac bool, fft bool) int {
	if source == "out" && dac && fft {
		return 1
	}
	if source == "dac" {
		dac = true
	}
	if source == "fft" {
		fft = true
	}
	if depth > 100 {
		return 0
	}
	total := 0
	for _, dest := range input[source] {
		total += partTwoBad(input, dest, depth+1, dac, fft)
	}
	return total
}

func partOneBad(input map[string][]string) {
	locs := map[string]loc{}
	for v := range input {
		locs[v] = loc{false, false, false, 0}
	}
	locs["out"] = loc{true, false, false, 1}
	counter := 0
	for !locs["you"].pathOut && counter < 2000 {
		counter++
		for source := range input {
			if len(input[source]) > 0 {
				total := 0
				done := true
				for _, dest := range input[source] {
					total += locs[dest].distance
					if !locs[dest].pathOut {
						done = false
					}
				}
				if done {
					newStruct := loc{true, false, false, total}
					locs[source] = newStruct
					input[source] = []string{}
				}
			}
		}
	}
	fmt.Println(locs)
	fmt.Println(locs["you"].distance)
}

type loc struct {
	pathOut  bool
	seenDac  bool
	seenFft  bool
	distance int
}

func getFiles(version string) (map[string][]string, error) {
	file, err := utils.LoadFile("day_11", version)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	data := map[string][]string{}
	for _, v := range file {
		a := strings.Split(v, ": ")
		b := strings.Split(a[1], " ")
		outs := []string{}
		for _, out := range b {
			outs = append(outs, out)
		}
		data[a[0]] = outs
	}

	return data, nil
}
