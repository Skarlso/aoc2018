package main

import (
	"fmt"
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

type pos struct {
	pos      coord
	distance int
}

var (
	up       = coord{x: 0, y: -1}
	down     = coord{x: 0, y: +1}
	left     = coord{x: -1, y: 0}
	right    = coord{x: 1, y: 0}
	floorMap = make([][]rune, 1)
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	floorMap[0] = make([]rune, 1)
	c := coord{x: 0, y: 0}
	distance := 0
	stack := make([]pos, 0)
	distances := make(map[coord]int, 0)
	// inf := int(math.Inf(0))
	for _, r := range content {
		if r == 'N' || r == 'E' || r == 'W' || r == 'S' {
			if r == 'N' {
				c.x = c.x + up.x
				c.y = c.y + up.y
			}
			if r == 'E' {
				c.x = c.x + right.x
				c.y = c.y + right.y
			}
			if r == 'S' {
				c.x = c.x + down.x
				c.y = c.y + down.y
			}
			if r == 'W' {
				c.x = c.x + left.x
				c.y = c.y + left.y
			}

			distance++
			if v, ok := distances[c]; ok {
				distances[c] = min(distance, v)
			} else {
				distances[c] = distance
			}
		} else if r == '(' {
			stack = append(stack, pos{pos: c, distance: distance})
		} else if r == ')' {
			var p pos
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			c, distance = p.pos, p.distance
		} else if r == '|' {
			var p pos
			p = stack[len(stack)-1]
			c, distance = p.pos, p.distance
		}
	}
	fmt.Println(max(distances))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a map[coord]int) int {
	m := 0
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}
