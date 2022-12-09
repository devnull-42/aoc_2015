package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	var raw int
	var quoted int

	for fs.Scan() {
		raw += len(fs.Text())
		s := strconv.Quote(fs.Text())
		quoted += len(s)
	}

	fmt.Printf("day8 part2: %v\n", quoted-raw)
}
