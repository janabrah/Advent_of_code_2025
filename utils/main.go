package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadFile(filenames ...string) ([]string, error) {
	var file []byte
	var err error
	if len(filenames) == 1 {
		file, err = os.ReadFile(filenames[0])
		if err != nil {
			return nil, err
		}
	} else {
		file, err = os.ReadFile("./inputs/" + filenames[0] + "_" + filenames[1] + ".txt")
		if err != nil {
			return nil, err
		}
	}
	stringFile := string(file)
	result := strings.Split(stringFile, "\n")
	return result, nil
}

func GetNumbers(input []string, divider string) ([][]int, error) {
	result := make([][]int, 0)
	for _, v := range input {
		row := make([]int, 0, 2)
		split := strings.Split(v, divider)
		for _, x := range split {
			val, err := strconv.Atoi(x)
			if err != nil {
				return nil, err
			}
			row = append(row, val)
		}
		result = append(result, row)
	}
	return result, nil
}

// It seemed fun to implement my own merge sort since I wasn't sure what go's builtin was
func SortIntArray(input []int) []int {
	result := make([]int, 0, len(input))
	if len(input) == 0 || len(input) == 1 {
		return input
	}
	split := len(input) / 2
	left := SortIntArray(input[:split])
	right := SortIntArray(input[split:])
	l, r := 0, 0
	for l < len(left) || r < len(right) {
		if l == len(left) {
			result = append(result, right[r])
			r++
		} else if r == len(right) {
			result = append(result, left[l])
			l++
		} else if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	return result
}

func PrettyPrint1D(input []string) {
	for _, v := range input {
		fmt.Println(v)
	}
}

func PrettyPrint2DString(input [][]string, spacer string) {
	for _, v := range input {
		for i := range v {
			if v[i] == "" {
				v[i] = " "
			}
		}
		fmt.Println(strings.Join(v, spacer))
	}
}

func PrettyPrint2DRune(input [][]rune, spacer rune) {
	for _, v := range input {
		stringArray := []string{}
		for i := range v {
			if v[i] == 0 {
				v[i] = ' '
			}
			stringArray = append(stringArray, string(v[i]))
		}
		fmt.Println(strings.Join(stringArray, string(spacer)))
	}
}

func PrettyPrint2DInt(input [][]int, spacer string) {
	for _, v := range input {
		var stringArray []string
		for _, s := range v {
			stringArray = append(stringArray, strconv.Itoa(s))
		}
		fmt.Println(strings.Join(stringArray, spacer))
	}
}

func Substring(input string, beginning int, end int) (string, error) {
	if beginning < 0 {
		return "", errors.New("beginning must be nonnegative")
	}
	if beginning >= end {
		return "", errors.New("beginning must be less than end")
	}
	runeString := []rune(input)
	if end > len(runeString) {
		return "", errors.New("end cannot be more than the length of the string")
	}
	subRunes := runeString[beginning:end]
	result := string(subRunes)
	return result, nil
}
