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

	var nice int

	for fs.Scan() {
		word := fs.Text()
		pair := false
		dupe := false

		for i := range word {

			if !pair && i+2 < len(word) {
				part1 := word[:i]
				part2 := word[i+2:]
				if strings.Contains(part1, word[i:i+2]) || strings.Contains(part2, word[i:i+2]) {
					pair = true
				}
			}

			if !dupe && i+2 < len(word) {
				if word[i] == word[i+2] {
					dupe = true
				}
			}
		}
		if pair && dupe {
			nice++
		}
	}

	fmt.Printf("day5 part2: %d\n", nice)
}
