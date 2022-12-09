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
		dims := make([]int, 3)
		dims[0], _ = strconv.Atoi(dim[0])
		dims[1], _ = strconv.Atoi(dim[1])
		dims[2], _ = strconv.Atoi(dim[2])

		sort.Sort(sort.IntSlice(dims))

		total += dims[0]*2 + dims[1]*2 + dims[0]*dims[1]*dims[2]
	}

	fmt.Printf("day2 part2: %d\n", total)
}
