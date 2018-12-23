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
	pos coord
	hp  int
	dmg int
	t   rune
}

func (e *enemy) String() string {
	return fmt.Sprintf("type: %s, x: %d, y: %d, hp: %d, dmg: %d;\n", string(e.t), e.pos.x, e.pos.y, e.hp, e.dmg)
}

func (e *enemy) scan() {
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
				if (e.pos.x+up.x) == g.pos.x && (e.pos.y+up.y) == g.pos.y ||
					(e.pos.x+down.x) == g.pos.x && (e.pos.y+down.y) == g.pos.y ||
					(e.pos.x+left.x) == g.pos.x && (e.pos.y+left.y) == g.pos.y ||
					(e.pos.x+right.x) == g.pos.x && (e.pos.y+right.y) == g.pos.y {
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
		fmt.Println("Attacks: ", min)
		e.attack(min)
		return
	}

}

func (e *enemy) move() {

}

func (e *enemy) attack(g *enemy) {
	g.hp -= e.dmg
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
					pos: coord{x: x, y: y},
					hp:  200,
					dmg: 3,
					t:   'G',
				}
				enemies = append(enemies, &g)
			} else if r == 'E' {
				e := enemy{
					pos: coord{x: x, y: y},
					hp:  200,
					dmg: 3,
					t:   'E',
				}
				enemies = append(enemies, &e)
			}
		}
	}
	display(playfield)
	fmt.Println(enemies)
	for {
		sort.Sort(enemySlice(enemies))
		for _, e := range enemies {
			e.scan()
		}
		break
	}
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
