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

	wires := make(map[string]uint16)
	gates := make(map[string][]string)

	for fs.Scan() {
		gate := strings.Split(fs.Text(), " ")
		gates[gate[len(gate)-1]] = gate[:len(gate)-2]
	}

	for len(wires) < len(gates) {
		for wire, gate := range gates {
			if _, ok := wires[wire]; ok {
				continue
			}
			switch {
			case len(gate) == 1:
				// initial value
				if val, err := strconv.Atoi(gate[0]); err == nil {
					wires[wire] = uint16(val)
					// direct input from another wire
				} else if _, ok := wires[gate[0]]; ok {
					wires[wire] = wires[gate[0]]
				}

			case len(gate) == 2:
				if _, ok := wires[gate[1]]; ok {
					if gate[0] == "NOT" {
						wires[wire] = ^(wires[gate[1]])
					}
				}

			case len(gate) == 3:
				switch gate[1] {
				case "AND":
					if inWires(wires, gate[0]) && inWires(wires, gate[2]) {
						wires[wire] = wires[gate[0]] & wires[gate[2]]
					} else if gate[0] == "1" && inWires(wires, gate[2]) {
						wires[wire] = 1 & wires[gate[2]]
					}
				case "OR":
					if inWires(wires, gate[0]) && inWires(wires, gate[2]) {
						wires[wire] = wires[gate[0]] | wires[gate[2]]
					}
				case "LSHIFT":
					amount, err := strconv.Atoi(gate[2])
					if inWires(wires, gate[0]) && err == nil {
						wires[wire] = wires[gate[0]] << amount
					}
				case "RSHIFT":
					amount, err := strconv.Atoi(gate[2])
					if inWires(wires, gate[0]) && err == nil {
						wires[wire] = wires[gate[0]] >> amount
					}
				}
			}
		}
	}
	fmt.Printf("day7 part1: %v\n", wires["a"])
}

func andGate(x, y int) func() {
	return func() {}
}

func isNumber(val string) bool {
	if _, err := strconv.Atoi(val); err != nil {
		return false
	}
	return true
}

func inWires(wires map[string]uint16, val string) bool {
	_, ok := wires[val]
	return ok
}
