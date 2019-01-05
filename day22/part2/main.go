package main

import (
	"container/heap"
	"fmt"

	"github.com/fatih/color"
)

type pos struct {
	x int
	y int
}

type coord struct {
	p        pos
	gear     int
	index    int
	priority int
}

func (r coord) String() string {
	return fmt.Sprintf("p: %v gear: %d index: %d priority: %d\n", r.p, r.gear, r.index, r.priority)
}

type pathPrioQueue []*coord

func (pq pathPrioQueue) Len() int { return len(pq) }
func (pq pathPrioQueue) Less(i, j int) bool {
	// implement the cost logic here?
	return pq[i].priority > pq[j].priority
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
	item.p.x = x
	item.p.y = y
	item.priority = p
	heap.Fix(pq, item.index)
}

func cost(from coord, to pos, cave [][]region) (moveCost int, switchedGear int) {
	if cave[from.p.y][from.p.x].t != cave[to.y][to.x].t {
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
	target         = pos{x: 10, y: 10}
	maxX           = target.x + 10
	maxY           = target.y + 10
	basicMoveCost  = 1
	gearSwitchCost = 7
)

func neighbours(c coord, cave [][]region) (paths []*coord) {
	// calculate and add movement cost if switching is needed.
	// fmt.Println(c)

	//
	if c.p.x+1 < len(cave[c.p.y]) {
		to := pos{x: c.p.x + 1, y: c.p.y}
		_, g := cost(c, to, cave)
		paths = append(paths, &coord{p: to, gear: g, priority: 0, index: 0})
	}
	if c.p.x-1 >= 0 {
		to := pos{x: c.p.x - 1, y: c.p.y}
		_, g := cost(c, to, cave)
		paths = append(paths, &coord{p: to, gear: g, priority: 0, index: 0})
	}
	if c.p.y+1 < len(cave) {
		to := pos{x: c.p.x, y: c.p.y + 1}
		_, g := cost(c, to, cave)
		paths = append(paths, &coord{p: to, gear: g, priority: 0, index: 0})
		// paths = append(paths, pos{x: c.x, y: c.y + 1})
	}
	if c.p.y-1 >= 0 {
		to := pos{x: c.p.x, y: c.p.y - 1}
		_, g := cost(c, to, cave)
		paths = append(paths, &coord{p: to, gear: g, priority: 0, index: 0})
		// paths = append(paths, pos{x: c.x, y: c.y - 1})
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
	from := make(map[pos]*coord, 0)
	goal := target
	start := coord{p: pos{x: 0, y: 0}, priority: 0, index: 0, gear: torch}
	path[0] = &start
	heap.Init(&path)
	from[start.p] = &coord{p: pos{y: -1, x: -1}}
	costSoFar := make(map[pos]int)
	costSoFar[start.p] = 0
	for len(path) > 0 {
		current := *path.Pop().(*coord)
		// time.Sleep(time.Millisecond * 200)
		// fmt.Println(current)
		if current.p == goal {
			// fmt.Printf("current: %v; goal: %v\n", current, goal)
			if current.gear != torch {
				costSoFar[current.p] += gearSwitchCost
				current.gear = torch
			}
			break
		}
		moves := neighbours(current, cave)
		for _, next := range moves {
			moveCost, newGear := cost(current, next.p, cave)
			newCost := costSoFar[current.p] + moveCost

			if _, ok := costSoFar[next.p]; !ok || newCost < costSoFar[next.p] {
				next.gear = newGear
				costSoFar[next.p] = newCost
				priority := newCost + distance(goal, next.p)
				next.priority = priority
				path.Push(next)
				path.update(next, next.p.x, next.p.y, priority)
				from[next.p] = &current
			}
		}
	}
	// fmt.Println(from)
	allPath := make([]pos, 0)
	current := goal
	// allCost := 0
	for current != start.p {
		allPath = append(allPath, current)
		next := from[current]
		current = next.p
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

func displayPath(path []pos, cave [][]region) {
	c := color.New(color.FgCyan).Add(color.Underline)
	for y := 0; y < len(cave); y++ {
		for x := 0; x < len(cave[y]); x++ {
			if contains(pos{y: y, x: x}, path) {
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

func contains(c pos, path []pos) bool {
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

func distance(a, b pos) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}
