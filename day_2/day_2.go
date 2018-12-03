package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines("day_2_input.txt")
	fmt.Printf("Checksum: %d\n", CheckSumBoxes(lines))
	fmt.Printf("Similar letters: %s\n", FindSimilarity(lines))
}

// CheckSumBoxes : Calculate the checksum of boxes whose
// IDs contain a certain number of repeated letters
func CheckSumBoxes(lines []string) int {
	goodBoxes := make(map[int]int, 0)
	lookUpNumbers := []int{2, 3}

	for _, boxID := range lines {
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

func commonLetters(lines []string) (string, string) {
	var resultCompareFrom, resultCompareTo string

	for x := 0; x < len(lines); x++ {
		compareFrom := lines[x]
		for y := x; y < len(lines); y++ {
			compareTo := lines[y]
			if x == y {
				continue
			}

			diff := 0
			for l := 0; l < len(compareTo); l++ {
				if compareFrom[l:l+1] != compareTo[l:l+1] {
					diff++
				}
			}
			if diff == 0 || diff == 1 {
				resultCompareFrom, resultCompareTo = lines[x], lines[y]
			}
		}
	}
	return resultCompareFrom, resultCompareTo
}

// FindSimilarity : find the most common letters in a large input
// of words
func FindSimilarity(lines []string) string {
	var similarities string
	compareFrom, compareTo := commonLetters(lines)

	for x := 0; x < len(compareFrom); x++ {
		if compareFrom[x] == compareTo[x] {
			similarities += string(compareFrom[x])
		}
	}
	return similarities
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
