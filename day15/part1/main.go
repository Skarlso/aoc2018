package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type coord struct {
	x int
	y int
}

type coordSlice []coord

func (c coordSlice) Len() int      { return len(c) }
func (c coordSlice) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c coordSlice) Less(i, j int) bool {
	if c[i].y < c[j].y {
		return true
	} else if c[i].y == c[j].y {
		if c[i].x < c[j].x {
			return true
		}
	}
	return false
}

// TODO: Instead of this the map could contain IDs of entites
// the map then could be used to look up the entity in O(1)
// instead of this slice which is O(n). This is unuseable in
// case there are thousands of entities.
type enemySlice []*enemy

func (c enemySlice) Len() int      { return len(c) }
func (c enemySlice) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c enemySlice) Less(i, j int) bool {
	if c[i].pos.y < c[j].pos.y {
		return true
	} else if c[i].pos.y == c[j].pos.y {
		if c[i].pos.x < c[j].pos.x {
			return true
		}
	}
	return false
}

// Create an interface saying enemy. These functions will be duplicated.
type enemy struct {
	pos  coord
	hp   int
	dmg  int
	t    rune
	dead bool
}

func (e *enemy) String() string {
	return fmt.Sprintf("type: %s, x: %d, y: %d, hp: %d, dmg: %d;\n", string(e.t), e.pos.x, e.pos.y, e.hp, e.dmg)
}

func (e *enemy) scan() bool {
	removeDead()
	if e.dead {
		return false
	}
	// Look around if there are elfs in the vicinity
	// Actually, look around for enemies. Calculate their distances with manhattan
	// if there is one with distance of 1 or several, put them into `near` list.
	// Then select the one with the lowest hitpoint and attack.
	// -- if yes... attack the lowest hitpoint elf
	// -- if no... continue with moving
	nearby := make([]*enemy, 0)
	for _, g := range enemies {
		dis := abs(e.pos.x-g.pos.x) + abs(e.pos.y-g.pos.y)
		if e.t != g.t {
			if dis <= 1 {
				// don't attack diagonally
				if ((e.pos.x+up.x) == g.pos.x && (e.pos.y+up.y) == g.pos.y ||
					(e.pos.x+down.x) == g.pos.x && (e.pos.y+down.y) == g.pos.y ||
					(e.pos.x+left.x) == g.pos.x && (e.pos.y+left.y) == g.pos.y ||
					(e.pos.x+right.x) == g.pos.x && (e.pos.y+right.y) == g.pos.y) &&
					!g.dead {
					nearby = append(nearby, g)
				}
			}
		}
	}

	if len(nearby) > 0 {
		// sort based on location so the reading order is applied to the attack.
		sort.Sort(enemySlice(nearby))
		min := nearby[0]
		for _, g := range nearby {
			if min.hp > g.hp {
				min = g
			}
		}
		// fmt.Println("Attacks: ", min)
		e.attack(min)
		return true
	}
	// var nearest *enemy
	// var pathToNearest []coord
	var pathsToEnemies = make([][]coord, 0)
	// find something to attack
	// enemies are sorted so when I'm going through them to find the
	// nearest, I'm doing it in reading order anyways.
	// gather all path to all enemies if reachable and select the shortest.
	for _, g := range enemies {
		if g.t != e.t {
			if v, ok := e.canReach(g); ok {
				pathsToEnemies = append(pathsToEnemies, v)
			}
		}
	}

	if len(pathsToEnemies) < 1 {
		return false
	}

	min := pathsToEnemies[0]
	for _, p := range pathsToEnemies {
		if len(p) < len(min) {
			min = p
		}
	}
	// fmt.Println("Current: ", e)
	// fmt.Println("Nearest: ", nearest)
	// fmt.Println("path to nearest: ", pathToNearest)
	e.move(min[len(min)-1])
	return true
}

func (e *enemy) getPathTo(g *enemy) (path []coord) {
	from := make(map[coord]coord, 0)
	goal := g.pos
	start := e.pos
	path = append(path, start)
	from[start] = coord{y: -1, x: -1}
	for len(path) > 0 {
		var current coord
		current, path = path[0], path[1:]
		if current == goal {
			break
		}
		var eType rune
		if e.t == 'G' {
			eType = 'E'
		} else {
			eType = 'G'
		}
		movesForCurrent := neighbours(current, eType)
		for _, m := range movesForCurrent {
			if _, ok := from[m]; !ok {
				from[m] = current
				path = append(path, m)
			}
		}
	}

	// if the goal is not in the list that means we can't reach it.
	if _, ok := from[goal]; !ok {
		return []coord{}
	}

	// Construct a path
	allPath := make([]coord, 0)
	current := goal

	for current != start {
		allPath = append(allPath, current)
		// fmt.Println("current: ", current)
		// fmt.Println("start: ", start)
		// fmt.Println("goal: ", goal)
		// fmt.Println("goblin: ", g)
		// fmt.Println("attacker: ", e)
		// fmt.Println(from)
		// display(playfield)
		current = from[current]
	}
	return allPath
}

func neighbours(v coord, e rune) (paths []coord) {
	// give back all the valid path around the given coordinate
	if playfield[v.y+right.y][v.x+right.x] == '.' || playfield[v.y+right.y][v.x+right.x] == e {
		paths = append(paths, coord{x: v.x + right.x, y: v.y + right.y})
	}
	if playfield[v.y+left.y][v.x+left.x] == '.' || playfield[v.y+left.y][v.x+left.x] == e {
		paths = append(paths, coord{x: v.x + left.x, y: v.y + left.y})
	}
	if playfield[v.y+down.y][v.x+down.x] == '.' || playfield[v.y+down.y][v.x+down.x] == e {
		paths = append(paths, coord{x: v.x + down.x, y: v.y + down.y})
	}
	if playfield[v.y+up.y][v.x+up.x] == '.' || playfield[v.y+up.y][v.x+up.x] == e {
		paths = append(paths, coord{x: v.x + up.x, y: v.y + up.y})
	}
	return
}

func removeDead() {
	for i := 0; i < len(enemies); i++ {
		if enemies[i].dead {
			playfield[enemies[i].pos.y][enemies[i].pos.x] = '.'
			enemies = append(enemies[:i], enemies[i+1:]...)
		}
	}
}

func contains(v coord, r []coord) bool {
	for _, c := range r {
		if c.x == v.x && c.y == v.y {
			return true
		}
	}
	return false
}

func (e *enemy) move(newLocation coord) {
	// update previous location to `.`
	// and set new location
	// fmt.Printf("Moving from %v to %v\n", e.pos, newLocation)
	playfield[e.pos.y][e.pos.x], playfield[newLocation.y][newLocation.x] = '.', e.t
	e.pos = newLocation
}

func (e *enemy) canReach(g *enemy) ([]coord, bool) {
	path := e.getPathTo(g)
	return path, len(path) > 0
}

func (e *enemy) attack(g *enemy) {
	g.hp -= e.dmg
	if g.hp <= 0 {
		g.dead = true
	}
}

var (
	up        = coord{x: 0, y: -1}
	down      = coord{x: 0, y: +1}
	left      = coord{x: -1, y: 0}
	right     = coord{x: 1, y: 0}
	playfield = make([][]rune, 0)
	enemies   = make([]*enemy, 0)
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	for y, l := range lines {
		playfield = append(playfield, []rune(l))
		for x, r := range l {
			if r == 'G' {
				g := enemy{
					pos:  coord{x: x, y: y},
					hp:   200,
					dmg:  3,
					t:    'G',
					dead: false,
				}
				enemies = append(enemies, &g)
			} else if r == 'E' {
				e := enemy{
					pos:  coord{x: x, y: y},
					hp:   200,
					dmg:  3,
					t:    'E',
					dead: false,
				}
				enemies = append(enemies, &e)
			}
		}
	}
	count := 0
	// because one enemy might say it can't reach any more
	// but others might be attacking.
	canReach := true
	for canReach {
		canReach = false
		sort.Sort(enemySlice(enemies))
		for _, e := range enemies {
			reach := e.scan()
			if !canReach && reach {
				canReach = true
			} else if !reach && canReach {
				// do nothing
			}
		}

		// display(playfield)
		// time.Sleep(200 * time.Millisecond)
		count++
	}
	fmt.Println("battle ended after: ", count)

	sum := 0
	for _, e := range enemies {
		sum += e.hp
	}
	fmt.Println("health sum: ", sum)
	fmt.Println("outcome: ", sum*count)
}

func display(r [][]rune) {
	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[y]); x++ {
			fmt.Print(string(r[y][x]))
		}
		fmt.Println()
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
