package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	var floor int

	fs.Scan()
	for _, r := range fs.Text() {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}

	}
	fmt.Printf("day1 part1: %d\n", floor)
}
