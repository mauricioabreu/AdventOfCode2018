// First day puzzle: https://adventofcode.com/2018/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Resulting frequency: %d\n", GetResultingFrequency("day_1_input.txt"))
	fmt.Printf("First resulting frequency seen twice: %d\n", GetFirstDuplicate("day_1_input.txt"))
}

// GetFirstDuplicate : retrieve the first result frequency seen twice
func GetFirstDuplicate(file string) int {
	var resultingFrequency int
	seen := make([]int, 0)

	for {
		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		s := bufio.NewScanner(f)
		for s.Scan() {
			currentFrequency, err := strconv.Atoi(s.Text())
			if err != nil {
				fmt.Printf("Could not convert the given frequency: %s\n", s.Text())
			}
			resultingFrequency += currentFrequency
			if hasDuplicate(resultingFrequency, seen) {
				return resultingFrequency
			}
			seen = append(seen, resultingFrequency)
		}
	}
}

// GetResultingFrequency : get the resulting frequency of the list
// It is the first one, different from the duplicated frequency that
// can be read more than once
func GetResultingFrequency(file string) int {
	var resultingFrequency int
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		currentFrequency, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Printf("Could not convert the given frequency: %s\n", s.Text())
		}
		resultingFrequency += currentFrequency
	}
	return resultingFrequency
}

func hasDuplicate(value int, values []int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
