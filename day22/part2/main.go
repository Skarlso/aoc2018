package main

import (
	"fmt"

	"github.com/fatih/color"
)

type coord struct {
	x     int
	y     int
	cost  int
	index int
}

type pathPrioQueue []*coord

func (pq pathPrioQueue) Len() int { return len(pq) }
func (pq pathPrioQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq pathPrioQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pathPrioQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*coord)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *pathPrioQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// // update modifies the priority and value of an Item in the queue.
// func (pq *pathPrioQueue) update(item *coord, x, y int, cost int) {
// 	item.x = x
// 	item.y = y
// 	item.cost = cost
// 	heap.Fix(pq, item.index)
// }

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
	maxX   = target.x + 10
	maxY   = target.y + 10
)

var allPath = make([][]coord, 0)

func neighbours(c coord, cave [][]region) (paths []coord) {
	// calculate and add movement cost if switching is needed.
	// fmt.Println(c)
	if c.x+1 < len(cave[c.y]) {
		paths = append(paths, coord{x: c.x + 1, y: c.y})
	}
	if c.x-1 > 0 {
		paths = append(paths, coord{x: c.x - 1, y: c.y})
	}
	if c.y+1 < len(cave) {
		paths = append(paths, coord{x: c.x, y: c.y + 1})
	}
	if c.y-1 >= 0 {
		paths = append(paths, coord{x: c.x, y: c.y - 1})
	}
	return
}

func main() {
	// Create the map
	cave := make([][]region, target.y+10)
	for y := 0; y <= target.y+9; y++ {
		cave[y] = make([]region, target.x+10)
		for x := 0; x <= target.x+9; x++ {
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

	s := santa{
		heldTool: torch,
		pos:      coord{y: 0, x: 0},
	}

	// paths := make([][]coord, 0)
	path := make([]coord, 0)
	visited := make(map[coord]coord, 0)
	goal := target
	start := s.pos
	path = append(path, start)
	visited[start] = coord{y: -1, x: -1}
	// visited[start] = true
	// sofar := make([]coord, 0)
	// cost so far
	for len(path) > 0 {
		var current coord
		current, path = path[0], path[1:]
		if current == goal {
			// paths = append(paths, sofar)
			// sofar = make([]coord, 0)
			break
		}
		moves := neighbours(current, cave)
		for _, m := range moves {
			if _, ok := visited[m]; !ok {
				visited[m] = current
				path = append(path, m)
				// sofar = append(sofar, m)
			}
		}
	}

	allPath := make([]coord, 0)
	current := goal

	for current != start {
		allPath = append(allPath, current)
		current = visited[current]
	}
	displayPath(allPath, cave)
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

func displayPath(path []coord, cave [][]region) {
	c := color.New(color.FgCyan).Add(color.Underline)
	for y := 0; y < len(cave); y++ {
		for x := 0; x < len(cave[y]); x++ {
			if contains(coord{y: y, x: x}, path) {
				switch cave[y][x].t {
				case rocky:
					c.Print(".")
				case wet:
					c.Print("=")
				case narrow:
					c.Print("|")
				}
			} else {
				switch cave[y][x].t {
				case rocky:
					fmt.Print(".")
				case wet:
					fmt.Print("=")
				case narrow:
					fmt.Print("|")
				}
			}
		}
		fmt.Println()
	}
}

func contains(c coord, path []coord) bool {
	for _, v := range path {
		if v == c {
			return true
		}
	}
	return false
}
