package main

import (
	"io/ioutil"
	"os"
)

// every step is a door
// #####
// #.|.#
// #-###
// #.|X#
// #####
// construct a map first
// after that we need to do a path search

type coord struct {
	x int
	y int
}

var (
	up    = coord{x: 0, y: -1}
	down  = coord{x: 0, y: +1}
	left  = coord{x: -1, y: 0}
	right = coord{x: 1, y: 0}
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	floorMap := make([][]rune, 1)
	floorMap[0] = make([]rune, 1)
	cx := 0
	cy := 0
	// once I encounter a ( I need the final closing ) and the string in between.
	for i := 0; i < len(content); i++ {
		switch content[i] {
		case 'N':
			cx += up.x
			cy += up.y
		case 'E':
			cx += right.x
			cy += right.y
		case 'S':
			cx += down.x
			cy += down.y
		case 'W':
			cx += left.x
			cy += left.y
		}
	}
}

// parseBranch returns the index at which to continue?
func parseBranch(branch string) (index int) {
	return
}
