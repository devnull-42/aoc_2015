package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	vowels := "aeiou"
	badStrings := []string{"ab", "cd", "pq", "xy"}

	var nice int

WordCheck:
	for fs.Scan() {
		vowelCount := 0
		duplicate := false
		word := fs.Text()

		for _, bs := range badStrings {
			if strings.Contains(word, bs) {
				continue WordCheck
			}
		}

		for i := range word {
			if strings.Contains(vowels, string(word[i])) {
				vowelCount++
			}
			if i > 0 {
				if word[i] == word[i-1] {
					duplicate = true
				}
			}
		}
		if vowelCount > 2 && duplicate {
			nice++
		}
	}

	fmt.Printf("day5 part1: %d\n", nice)
}
