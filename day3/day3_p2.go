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
		santaX int
		santaY int
		roboX  int
		roboY  int
	)

	// initial delivery
	houses["0,0"] = struct{}{}

	for fs.Scan() {
		instructions := fs.Text()

		for i, ins := range instructions {
			if i%2 == 0 {
				switch ins {
				case '^':
					roboY--
				case 'v':
					roboY++
				case '<':
					roboX--
				case '>':
					roboX++
				}
				houses[fmt.Sprintf("%d,%d", roboY, roboX)] = struct{}{}
			} else {
				switch ins {
				case '^':
					santaY--
				case 'v':
					santaY++
				case '<':
					santaX--
				case '>':
					santaX++
				}
				houses[fmt.Sprintf("%d,%d", santaY, santaX)] = struct{}{}
			}
		}
	}

	fmt.Printf("day3 part1: %d\n", len(houses))
}
