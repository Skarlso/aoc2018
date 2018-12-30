package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			getNeighbours(x, y)
		}
	}
}

func getNeighbours(x, y int) {
	types := make(map[rune]int)

	if x+1 < maxX && y+1 < maxY {
		types[field[y+1][x+1]]++
	}
	if x-1 >= 0 && y+1 < maxY {
		types[field[y+1][x-1]]++
	}
	if y+1 < maxY {
		types[field[y+1][x]]++
	}
	if x+1 < maxX {
		types[field[y][x+1]]++
	}
	if x+1 < maxX && y-1 >= 0 {
		types[field[y-1][x+1]]++
	}
	if y-1 >= 0 {
		types[field[y-1][x]]++
	}
	if x-1 >= 0 && y-1 >= 0 {
		types[field[y-1][x-1]]++
	}
	if x-1 >= 0 {
		types[field[y][x-1]]++
	}
	lumberyard := types['#']
	open := types['.']
	trees := types['|']
	fmt.Printf("x:%d y:%d lumberyard:%d trees:%d open:%d\n", x, y, lumberyard, trees, open)
}

func display(r [][]rune) {
	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[y]); x++ {
			fmt.Print(string(r[y][x]))
		}
		fmt.Println()
	}
}
