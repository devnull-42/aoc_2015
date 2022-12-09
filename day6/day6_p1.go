package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	l := newLights()

	for fs.Scan() {
		instruction := fs.Text()
		instructions := strings.Split(instruction, " ")

		// remove turn word
		if instructions[0] == "turn" {
			instructions = instructions[1:]
		}

		fromSlice := strings.Split(instructions[1], ",")
		toSlice := strings.Split(instructions[3], ",")

		fromX, _ := strconv.Atoi(fromSlice[0])
		fromY, _ := strconv.Atoi(fromSlice[1])
		toX, _ := strconv.Atoi(toSlice[0])
		toY, _ := strconv.Atoi(toSlice[1])

		l.toggle([]int{fromX, fromY}, []int{toX, toY}, instructions[0])
	}

	fmt.Printf("day6 part1: %d\n", l.numberOn())
}

type lights struct {
	grid [][]bool
}

func newLights() *lights {
	l := new(lights)
	l.grid = make([][]bool, 1000)
	for i := range l.grid {
		l.grid[i] = make([]bool, 1000)
	}
	return l
}

func (l *lights) toggle(from, to []int, setting string) {
	for y := from[1]; y <= to[1]; y++ {
		for x := from[0]; x <= to[0]; x++ {
			switch setting {
			case "off":
				l.grid[y][x] = false
			case "on":
				l.grid[y][x] = true
			case "toggle":
				l.grid[y][x] = !l.grid[y][x]
			}
		}
	}
}

func (l lights) numberOn() int {
	count := 0
	for y := 0; y < len(l.grid); y++ {
		for x := 0; x < len(l.grid[y]); x++ {
			if l.grid[y][x] {
				count++
			}
		}
	}
	return count
}
