package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type pos struct {
	x int
	y int
	z int
	r int
}

func main() {
	// read all numbers to a []int
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	positions := make([]pos, 0)
	largestRadius := 0
	var largestRadiusPosition pos
	for _, l := range lines {
		var (
			x, y, z, r int
		)
		fmt.Sscanf(l, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		p := pos{x: x, y: y, z: z, r: r}
		positions = append(positions, p)
		if r > largestRadius {
			largestRadius = r
			largestRadiusPosition = p
		}
	}
	// fmt.Println(positions)
	fmt.Println("Largest Radius: ", largestRadius)
	fmt.Println("Largest Radius position: ", largestRadiusPosition)
	inRange := inRangePositions(largestRadiusPosition, positions)
	fmt.Println("In range positions: ", inRange)
}

func inRangePositions(p pos, positions []pos) (i int) {
	for _, po := range positions {
		if distance(p, po) < p.r {
			i++
		}
	}
	return i
}

func distance(p1, p2 pos) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
