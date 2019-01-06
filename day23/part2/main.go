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
	maxX, maxY, maxZ := 0, 0, 0
	minX, minY, minZ := 10000000000000, 10000000000000, 10000000000000

	for _, l := range lines {
		var (
			x, y, z, r int
		)
		fmt.Sscanf(l, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		if z < minZ {
			minZ = z
		}
		if z > maxZ {
			maxZ = z
		}
		p := pos{x: x, y: y, z: z, r: r}
		positions = append(positions, p)
	}
	// fmt.Println(positions)
	// fmt.Println("Largest Radius: ", largestRadius)
	// fmt.Println("Largest Radius position: ", largestRadiusPosition)
	var largestInRangePosition pos
	largestInRangeNumber := 0
	// increase the coordinates and get in range.
	// save the largest
	// fmt.Println(minX, maxX, minY, maxY, minZ, maxZ)
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			for k := minZ; k <= maxZ; k++ {
				p := pos{x: i, y: j, z: k}
				inRange := inRangePositions(p, positions)
				if inRange > largestInRangeNumber {
					largestInRangeNumber = inRange
					largestInRangePosition = p
				}
			}
		}
	}

	fmt.Println("Position that's in range of the largest number of positions: ", largestInRangePosition)
	fmt.Println("Distance from 0,0,0: ", distance(pos{0, 0, 0, 0}, largestInRangePosition))

}

func inRangePositions(p pos, positions []pos) (i int) {
	for _, po := range positions {
		if distance(p, po) <= po.r {
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
