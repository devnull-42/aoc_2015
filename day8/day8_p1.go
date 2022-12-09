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
	var unRaw int

	for fs.Scan() {
		raw += len(fs.Text())
		s, _ := strconv.Unquote(fs.Text())
		unRaw += len(s)
	}

	fmt.Printf("day8 part1: %v\n", raw-unRaw)
}
