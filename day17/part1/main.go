package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// read all numbers to a []int
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	g := render(lines)
	display(g)
}

func display(r [][]rune) {
	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[y]); x++ {
			fmt.Print(string(r[y][x]))
		}
		fmt.Println()
	}
}

func render(lines []string) [][]rune {
	ground := make([][]rune, 0)
	minX := 100000000
	minY := 100000000
	maxX := 0
	maxY := 0
	return ground
}
