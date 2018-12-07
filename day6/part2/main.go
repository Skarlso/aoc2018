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

const (
	minC      = 0
	maxC      = 359
	threshold = 10000
)

var coords = make([]*coord, 0)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	for _, l := range lines {
		var x, y int
		fmt.Sscanf(l, "%d, %d", &x, &y)
		c := coord{
			x: x,
			y: y,
		}
		coords = append(coords, &c)
	}

	safeRegionSize := 0
	for x := 0; x <= maxC; x++ {
		for y := 0; y <= maxC; y++ {
			sum := 0
			for _, c := range coords {
				sum += abs(c.x-x) + abs(c.y-y)
			}
			if sum < threshold {
				safeRegionSize++
			}
		}
	}
	fmt.Println(safeRegionSize)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
