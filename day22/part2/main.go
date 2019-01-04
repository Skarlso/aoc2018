package main

import "fmt"

type coord struct {
	x int
	y int
}

type region struct {
	geoindex int
	erosion  int
	t        int
}

func (r *region) String() string {
	return fmt.Sprintf("index: %d erosion level: %d type: %d\n", r.geoindex, r.erosion, r.t)
}

const (
	rocky = iota
	wet
	narrow
)

const (
	neither = iota
	torch
	gear
)

type santa struct {
	pos      coord
	heldTool int
}

var (
	// depth  = 5355
	// target = coord{x: 14, y: 796}
	depth  = 510
	target = coord{x: 10, y: 10}
	up     = coord{x: 0, y: -1}
	down   = coord{x: 0, y: +1}
	left   = coord{x: -1, y: 0}
	right  = coord{x: 1, y: 0}
)

var allPath = make([][]coord, 0)

// func getAllPathToTarget(x, y int) {
// 	currentPath = make([]coord, 0)

// }

// also
func neighbours(x, y int, cave [][]region) (path []coord) {
	// give back all the valid path around the given coordinate
	// if cave[y+right.y][x+right.x] == '.' || cave[y+right.y][x+right.x] == e {
	// 	paths = append(paths, coord{x: x + right.x, y: y + right.y})
	// }
	// if cave[y+left.y][x+left.x] == '.' || cave[y+left.y][x+left.x] == e {
	// 	paths = append(paths, coord{x: x + left.x, y: y + left.y})
	// }
	// if cave[y+down.y][x+down.x] == '.' || cave[y+down.y][x+down.x] == e {
	// 	paths = append(paths, coord{x: x + down.x, y: y + down.y})
	// }
	// if cave[y+up.y][x+up.x] == '.' || cave[y+up.y][x+up.x] == e {
	// 	paths = append(paths, coord{x: x + up.x, y: y + up.y})
	// }
	if y+right.y < len(cave) {
		paths = append(paths, coord{x: x + right.x, y: y + right.y})
	}
	if cave[y+left.y][x+left.x] == '.' || cave[y+left.y][x+left.x] == e {
		paths = append(paths, coord{x: x + left.x, y: y + left.y})
	}
	if cave[y+down.y][x+down.x] == '.' || cave[y+down.y][x+down.x] == e {
		paths = append(paths, coord{x: x + down.x, y: y + down.y})
	}
	if cave[y+up.y][x+up.x] == '.' || cave[y+up.y][x+up.x] == e {
		paths = append(paths, coord{x: x + up.x, y: y + up.y})
	}
	return
}

func main() {
	// Create the map
	cave := make([][]region, target.y+1)
	for y := 0; y <= target.y; y++ {
		cave[y] = make([]region, target.x+1)
		for x := 0; x <= target.x; x++ {
			if y == 0 && x == 0 {
				cave[y][x] = region{geoindex: 0, t: 0, erosion: 0}
				continue
			}

			if y == target.y && x == target.x {
				cave[y][x] = region{geoindex: 0, t: 0, erosion: 0}
				continue
			}

			var r region
			if y == 0 && x > 0 {
				r = region{
					geoindex: (x * 16807),
					erosion:  0,
					t:        0,
				}
			} else if x == 0 && y > 0 {
				r = region{
					geoindex: (y * 48271),
					erosion:  0,
					t:        0,
				}
			} else if x > 0 && y > 0 {
				r = region{
					geoindex: cave[y][x-1].erosion * cave[y-1][x].erosion,
					erosion:  0,
					t:        0,
				}
			}
			r.erosion = (r.geoindex + depth) % 20183
			r.t = r.erosion % 3
			cave[y][x] = r
		}
	}
	display(cave)

	s := santa{
		heldTool: torch,
		pos:      coord{y: 0, x: 0},
	}

	paths := make([][]coord, 0)
	path := make([]coord, 0)
	visited := make(map[coord]bool, 0)
	goal := target
	start := s.pos
	path = append(path, start)
	visited[start] = true
	sofar := make([]coord, 0)
	for len(path) > 0 {
		var current coord
		current, path = path[0], path[1:]
		if current == goal {
			paths = append(paths, sofar)
			sofar = make([]coord, 0)
			continue
		}
		moves := neighbours()
		for _, m := range moves {
			if _, ok := visited[m]; !ok {
				visited[m] = true
				path = append(path, m)
				sofar = append(sofar, m)
			}
		}
	}
}

func display(r [][]region) {
	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[y]); x++ {
			switch r[y][x].t {
			case rocky:
				fmt.Print(".")
			case wet:
				fmt.Print("=")
			case narrow:
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}
