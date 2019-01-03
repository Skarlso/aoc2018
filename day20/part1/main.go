package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type coord struct {
	x int
	y int
}

type pos struct {
	pos   coord
	doors int
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
	c := coord{x: 0, y: 0}
	doorCount := 0
	stack := make([]pos, 0)
	doors := make(map[coord]int, 0)
	for _, r := range content {
		switch r {
		case 'N', 'E', 'W', 'S':
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

			doorCount++
			if v, ok := doors[c]; ok {
				doors[c] = min(doorCount, v)
			} else {
				doors[c] = doorCount
			}
		case '(':
			stack = append(stack, pos{pos: c, doors: doorCount})
		case ')':
			var p pos
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			c, doorCount = p.pos, p.doors
		case '|':
			p := stack[len(stack)-1]
			c, doorCount = p.pos, p.doors
		}
	}
	fmt.Println(max(doors))
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
