package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type coord struct {
	n        int
	x        int
	y        int
	invalid  bool
	closesTo int
}

func (c *coord) String() string {
	return fmt.Sprintf("x: %d, y: %d, n: %d, invalid: %v, closesTo: %d", c.x, c.y, c.n, c.invalid, c.closesTo)
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
	for i, l := range lines {
		var x, y int
		fmt.Sscanf(l, "%d, %d", &x, &y)
		c := coord{
			n:        i,
			x:        x,
			y:        y,
			invalid:  false,
			closesTo: 0,
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
