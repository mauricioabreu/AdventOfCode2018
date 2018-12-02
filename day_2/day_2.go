package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Checksum: %d\n", CheckSumBoxes())
}

// CheckSumBoxes : Calculate the checksum of boxes whose
// IDs contain a certain number of repeated letters
func CheckSumBoxes() int {
	goodBoxes := make(map[int]int, 0)
	lookUpNumbers := []int{2, 3}

	for _, boxID := range readLines("day_2_input.txt") {
		lettersOccurrence := countLetterOccurrence(boxID)
		// Try to find the boxes that have the exactly number
		// of repeated letters in the box
		// When we find it, we try to find the next number
		for _, n := range lookUpNumbers {
			for _, c := range lettersOccurrence {
				if c == n {
					goodBoxes[n]++
					break // Already find at least one occurrence, find the next
				}
			}
		}
	}
	return checkSum(goodBoxes)
}

// Read lines of a file into an array
func readLines(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func checkSum(lettersOccurrence map[int]int) int {
	total := 1
	for _, v := range lettersOccurrence {
		total *= v
	}
	return total
}

func countLetterOccurrence(letters string) map[string]int {
	lettersOccurrence := make(map[string]int)
	for _, letter := range letters {
		lettersOccurrence[string(letter)]++
	}
	return lettersOccurrence
}
