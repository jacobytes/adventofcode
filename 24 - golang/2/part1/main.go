package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	lines, err := readInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	validReports := 0

	for _, line := range lines {

		characters := strings.Split(line, " ")

		ints := convertToIntSlice(characters)

		first := ints[0]
		second := ints[1]

		hasIncreasingStart := first < second
		hasDecreasingStart := first > second

		isValid := true

		for i := 1; i < len(ints); i++ {

			if ints[i-1] == ints[i] {
				isValid = false
				break
			}

			if hasDecreasingStart {

				unexpectedIncrease := ints[i-1] < ints[i]
				if unexpectedIncrease || ints[i-1]-ints[i] > 3 {
					isValid = false
					break
				}
			}

			if hasIncreasingStart {
				unexpectedDecrease := ints[i-1] > ints[i]
				if unexpectedDecrease || ints[i]-ints[i-1] > 3 {
					isValid = false
					break
				}
			}
		}

		if isValid {
			validReports++
		}

	}

	fmt.Println(validReports)
}

func convertToIntSlice(strings []string) []int {
	ints := make([]int, len(strings))

	for i, str := range strings {
		intVal, _ := strconv.Atoi(str)

		ints[i] = intVal
	}

	return ints
}

func readInput() ([]string, error) {
	content, err := os.ReadFile("../input.txt")

	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(string(content), "\n")

	return lines, nil
}
