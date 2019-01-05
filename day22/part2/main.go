package main

import (
	"container/heap"
	"fmt"

	"github.com/fatih/color"
)

type coord struct {
	x        int
	y        int
	gear     int
	index    int
	priority int
}

type pathPrioQueue []*coord

func (pq pathPrioQueue) Len() int { return len(pq) }
func (pq pathPrioQueue) Less(i, j int) bool {
	// implement the cost logic here?
	return pq[i].priority < pq[j].priority
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

// update modifies the priority and value of an Item in the queue.
func (pq *pathPrioQueue) update(item *coord, x, y int, p int) {
	item.x = x
	item.y = y
	item.priority = p
	heap.Fix(pq, item.index)
}

func cost(from coord, to coord, gear int, cave [][]region) (moveCost int, switchedGear int) {
	if cave[from.y][from.x].t != cave[to.y][to.x].t {
		toType := cave[to.y][to.x].t

		if (toType == rocky || toType == narrow) && from.gear == torch {
			return basicMoveCost, from.gear
		}

		if toType == wet && from.gear == torch {
			return gearSwitchCost, neither
		}

		if toType == rocky && from.gear == neither {
			return gearSwitchCost, torch
		}
	}
	return basicMoveCost, from.gear
}

type region struct {
	geoindex int
	erosion  int
	t        int
	gear     int
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
	torch = iota
	neither
	climbingGear // never use this
)

var (
	// depth  = 5355
	// target = coord{x: 14, y: 796}
	depth          = 510
	target         = coord{x: 10, y: 10, priority: 0}
	maxX           = target.x + 10
	maxY           = target.y + 10
	basicMoveCost  = 1
	gearSwitchCost = 7
)

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

	path := make(pathPrioQueue, 1)
	from := make(map[coord]*coord, 0)
	goal := target
	start := coord{x: 0, y: 0, priority: 0, index: 0, gear: torch}
	start.priority = 0
	path[0] = &start
	heap.Init(&path)
	from[start] = &coord{y: -1, x: -1}
	costSoFar := make(map[coord]int)
	costSoFar[start] = 0
	for len(path) > 0 {
		current := *path.Pop().(*coord)
		if current == goal {
			if current.gear != torch {
				costSoFar[current] += gearSwitchCost
				current.gear = torch
			}
			break
		}
		moves := neighbours(current, cave)
		for _, next := range moves {
			moveCost, newGear := cost(current, next, current.gear, cave)
			newCost := costSoFar[current] + moveCost

			if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
				next.gear = newGear
				costSoFar[next] = newCost
				priority := newCost + distance(goal, next)
				next.priority = priority
				// path.Push(next)
				heap.Push(&path, &next)
				path.update(&next, next.x, next.y, priority)
				from[next] = &current
			}
		}
	}
	fmt.Println(path)
	// allPath := make([]coord, 0)
	// current := goal

	// for current != start {
	// 	allPath = append(allPath, current)
	// 	current = *from[current]
	// }
	// displayPath(allPath, cave)
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distance(a, b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}
