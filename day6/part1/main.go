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

type mapCoord struct {
	x int
	y int
}

const (
	minC = 0
	maxC = 10
)

var coords = make([]coord, 0)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	areas := make([]int, 0)
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
		coords = append(coords, c)
	}
	for _, c := range coords {
		directions := []struct {
			x int
			y int
		}{
			{x: 0, y: 0},
			{x: -1, y: -1},
			{x: 1, y: 1},
			{x: 0, y: 1},
			{x: 0, y: -1},
			{x: -1, y: 1},
			{x: 1, y: 0},
			{x: 1, y: -1},
		}
		for _, d := range directions {
			closest, invalid := checkIfClosestInDirection(c.x, c.y, d.x, d.y)
			if invalid {
				c.invalid = true
				break
			}
			c.closesTo += closest
		}
		if !c.invalid {
			areas = append(areas, c.closesTo)
		}
		fmt.Printf("Checking: %d; with x: %d, y: %d; closestTo: %d, invalid: %v\n", c.n, c.x, c.y, c.closesTo, c.invalid)
	}
	sort.Ints(areas)
	fmt.Println(areas)
}

// Checks if this node is closest to that coordinate by looping through
// the rest of the coordinates and calculating the manhattan distance.
// If the new coordinate is > max or < min, it means it's an edge node,
// and extends into infinity therefore, invalid. It needs to includ itself too.
func checkIfClosestInDirection(startX, startY, dX, dY int) (closest int, invalid bool) {
	// keep adding the direction until either we hit the "edge" and become invalid
	// or find someone who is closer to that location.
	for {
		edge := false
		x := startX + dX
		y := startY + dY
		if x <= minC || y <= minC || x >= maxC || y >= maxC {
			fmt.Println("at the edge")
			edge = true
		}
		fmt.Println(x, y)
		for _, c := range coords {
			// if c.x == startX && c.y == startY {
			// 	closest++
			// 	continue
			// }
			myDistance := abs(startX-x) + abs(startY-y)
			neighbourDistance := abs(c.x-x) + abs(c.y-y)
			fmt.Printf("MyDistance: %d; NeighbourDistance: %d\n", myDistance, neighbourDistance)
			if myDistance < neighbourDistance {
				if edge {
					fmt.Println("marking invalid")
					return 0, true
				}
				closest++
			} else {
				return closest, false
			}
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
