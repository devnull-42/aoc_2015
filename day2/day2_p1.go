package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	fs := bufio.NewScanner(file)

	fs.Split(bufio.ScanLines)

	var total int

	for fs.Scan() {
		dim := strings.Split(fs.Text(), "x")
		length, _ := strconv.Atoi(dim[0])
		width, _ := strconv.Atoi(dim[1])
		height, _ := strconv.Atoi(dim[2])

		sides := make([]int, 3)
		sides[0] = length * width
		sides[1] = length * height
		sides[2] = width * height

		sort.Sort(sort.IntSlice(sides))

		total += sides[0]*2 + sides[1]*2 + sides[2]*2 + sides[0]
	}

	fmt.Printf("day2 part1: %d\n", total)
}
