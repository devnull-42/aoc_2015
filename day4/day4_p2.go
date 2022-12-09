package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanWords)
	fs.Scan()

	secret := fs.Text()
	i := 1

	for {
		val := fmt.Sprintf("%s%d", secret, i)
		h := md5.New()
		h.Write([]byte(val))
		valHex := hex.EncodeToString(h.Sum(nil))

		if strings.HasPrefix(valHex, "000000") {
			break
		}
		i++
	}

	fmt.Printf("day4 part2: %d\n", i)
}
