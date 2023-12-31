package main

import (
	"bufio"
	"os"
)

func part1() int {
	file, err := os.Open("input")
	if err != nil {
		return 0
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		for i := 0; i < len(line)-4; i++ {
			arr := []byte{line[i], line[i+1], line[i+2], line[i+3]}
			if hasDuplicate(arr) {
				continue
			}

			return i + 4
		}
	}
	return 0
}
