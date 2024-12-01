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

	var listA, listB []int

	for _, line := range lines {
		varA, varB, isFound := strings.Cut(line, "   ")

		if !isFound {
			fmt.Println("Invalid separator")
			return
		}

		intA, err := strconv.Atoi(varA)

		if err != nil {
			fmt.Println("VarA is an invalid int")
			return
		}

		intB, err := strconv.Atoi(varB)

		if err != nil {
			fmt.Println("VarB is an invalid int")
			return
		}

		listA = append(listA, intA)
		listB = append(listB, intB)

	}

	var result int
	occurrenceIndex := make(map[string]int)

	for _, number := range listB {

		occurrenceIndex[strconv.Itoa(number)]++
	}

	for _, number := range listA {

		totalOccurrences, hasOccurred := occurrenceIndex[strconv.Itoa(number)]

		if hasOccurred {
			result += number * totalOccurrences
		}
	}

	fmt.Println(result)
}

func readInput() ([]string, error) {
	content, err := os.ReadFile("../input.txt")

	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(string(content), "\n")

	return lines, nil
}
