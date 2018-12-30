package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	// OPEN .
	OPEN = iota
	// TREE |
	TREE
	// LUMBERYARD #
	LUMBERYARD
)

type coord struct {
	x int
	y int
}

var directions = []coord{
	coord{x: 0, y: -1},
	coord{x: -1, y: 0},
	coord{x: 1, y: 0},
	coord{x: 0, y: 1},
	coord{x: -1, y: -1},
	coord{x: 1, y: 1},
	coord{x: 1, y: -1},
	coord{x: -1, y: 1},
}

var field = make([][]rune, 0)
var maxX, maxY int

func main() {
	// read all numbers to a []int
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	for _, l := range lines {
		r := []rune(l)
		maxX = len(r)
		field = append(field, r)
	}

	maxY = len(field)

	display(field)
}

func gatherAcres(x, y int) {
	// x-2
	// y-2
	// x-2, y-2
	// x-2, y
	// x, y-2
	//
}

func getAcreType(x, y int) int {
	lumberyard := 0
	trees := 0
	open := 0
	for _, d := range directions {
		if x+d.x > 0 && x+d.x < maxX && y+d.y < maxY && y+d.y > 0 {
			if field[y][x] == '#' {
				lumberyard++
			} else if field[y][x] == '|' {
				trees++
			} else if field[y][x] == '.' {
				open++
			}
		}
	}
	if lumberyard == 8 {
		return LUMBERYARD
	}
	if trees == 8 {
		return TREE
	}
	if open == 8 {
		return OPEN
	}
	return -1
}

func display(r [][]rune) {
	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[y]); x++ {
			fmt.Print(string(r[y][x]))
		}
		fmt.Println()
	}
}
