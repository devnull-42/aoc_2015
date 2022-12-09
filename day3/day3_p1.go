package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	fs := bufio.NewScanner(file)

	fs.Split(bufio.ScanLines)

	houses := make(map[string]interface{})
	var (
		x int
		y int
	)

	// initial delivery
	houses["0,0"] = struct{}{}

	for fs.Scan() {
		instructions := fs.Text()

		for _, ins := range instructions {
			switch ins {
			case '^':
				y--
			case 'v':
				y++
			case '<':
				x--
			case '>':
				x++
			}
			houses[fmt.Sprintf("%d,%d", y, x)] = struct{}{}
		}
	}

	fmt.Printf("day3 part1: %d\n", len(houses))
}
