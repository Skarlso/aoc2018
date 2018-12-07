package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type coord struct {
	n        int
	x        int
	y        int
	invalid  bool
	closesTo int
}

const (
	minC = 0
	maxC = 359
)

var coords = make([]*coord, 0)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	areas := make([]int, 0)
	for _, l := range lines {
		var x, y int
		fmt.Sscanf(l, "%d, %d", &x, &y)
		c := coord{
			x:        x,
			y:        y,
			invalid:  false,
			closesTo: 0,
		}
		coords = append(coords, &c)
	}

	for _, c := range coords {
		// I could channel this out with go routines.
		closest, invalid := checkIfClosestInAllDirections(c.x, c.y)
		if invalid {
			c.invalid = true
		}
		c.closesTo += closest
		if !c.invalid {
			areas = append(areas, c.closesTo)
		}
	}
	sort.Ints(areas)
	fmt.Println(areas)
}

func checkIfClosestInAllDirections(startX, startY int) (closest int, invalid bool) {
	for x := 0; x <= maxC; x++ {
		for y := 0; y <= maxC; y++ {
			myDistance := abs(startX-x) + abs(startY-y)
			mostClose := false
			for _, c := range coords {
				if c.x == startX && c.y == startY {
					continue
				}
				neighbourDistance := abs(c.x-x) + abs(c.y-y)
				if myDistance < neighbourDistance {
					mostClose = true
				} else {
					mostClose = false
					break
				}
			}

			if mostClose {
				if x <= minC || y <= minC || x >= maxC || y >= maxC {
					return closest, true
				}
				closest++
			}
		}
	}
	return closest, invalid
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
