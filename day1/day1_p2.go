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
	var basement int

	fs.Scan()
	for i, r := range fs.Text() {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor < 0 {
			basement = i + 1
			break
		}
	}
	fmt.Printf("day1 part1: %d\n", basement)
}
