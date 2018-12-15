package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

type coord struct {
	x int
	y int
}

var railroad = make([][]rune, 0)
var directions = map[rune]coord{
	'>': coord{x: 1, y: 0},
	'<': coord{x: -1, y: 0},
	'^': coord{x: 0, y: -1},
	'v': coord{x: 0, y: 1},
}

type cart struct {
	id           int
	loc          coord
	direction    rune
	intersection int
	dead         bool
}

type carts []*cart

// Move moves the given cart into the direction that it's currently
// facing. If the cart encounters an intersection it will change
// its current heading.
func (c *cart) move(cs carts) bool {
	c.loc.x += directions[c.direction].x
	c.loc.y += directions[c.direction].y

	// TODO: Use integers instead of rune to determine the directions
	// This is too sloppy.
	switch railroad[c.loc.y][c.loc.x] {
	case '/':
		if c.direction == '^' {
			c.direction = '>'
		} else if c.direction == '<' {
			c.direction = 'v'
		} else if c.direction == 'v' {
			c.direction = '<'
		} else if c.direction == '>' {
			c.direction = '^'
		}
	case '\\':
		if c.direction == '^' {
			c.direction = '<'
		} else if c.direction == '>' {
			c.direction = 'v'
		} else if c.direction == 'v' {
			c.direction = '>'
		} else if c.direction == '<' {
			c.direction = '^'
		}
	case '+':
		if c.intersection == 0 {
			switch c.direction {
			case 'v':
				c.direction = '>'
			case '>':
				c.direction = '^'
			case '^':
				c.direction = '<'
			case '<':
				c.direction = 'v'
			}
		} else if c.intersection == 2 {
			switch c.direction {
			case 'v':
				c.direction = '<'
			case '>':
				c.direction = 'v'
			case '^':
				c.direction = '>'
			case '<':
				c.direction = '^'
			}
		}
		c.intersection = (c.intersection + 1) % 3
	}

	for _, other := range cs {
		if c.compare(other) {
			other.dead = true
			c.dead = true
			return true
		}
	}
	return false
}

func (c *cart) compare(other *cart) bool {
	return c.loc.x == other.loc.x && c.loc.y == other.loc.y && c.id != other.id
}

func (c *cart) String() string {
	return fmt.Sprintf("x: %d, y: %d, heading: %s, currentTrack: %s\n",
		c.loc.x,
		c.loc.y,
		string(c.direction),
		string(railroad[c.loc.y][c.loc.x]))
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	cs := make(carts, 0)
	for _, l := range lines {
		railroad = append(railroad, []rune(l))
	}

	id := 0
	for y := 0; y < len(railroad); y++ {
		for x := 0; x < len(railroad[y]); x++ {
			if railroad[y][x] == '>' ||
				railroad[y][x] == '<' ||
				railroad[y][x] == 'v' ||
				railroad[y][x] == '^' {
				c := cart{
					id:           id,
					loc:          coord{x: x, y: y},
					direction:    railroad[y][x],
					intersection: 0,
					dead:         false,
				}
				switch railroad[y][x] {
				case '>', '<':
					railroad[y][x] = '-'
				case '^', 'v':
					railroad[y][x] = '|'
				}
				cs = append(cs, &c)
				id++
			}
		}
	}

	crash := false
	ticks := 0
	for len(cs) > 1 {
		ticks++
		cs = moveCarts(cs)
		if crash {
			break
		}
		// display(railroad, cs)
		// time.Sleep(500 * time.Millisecond)
		// fmt.Println(cs)
	}
	fmt.Println(ticks)
	fmt.Println("Last cart location: ", cs[0].loc)
}

func showCarts(cs carts) {
	for _, c := range cs {
		fmt.Println(c)
	}
}

func moveCarts(cs carts) carts {
	newCarts := make(carts, 0)
	// cleanup := false
	for _, c := range cs {
		c.move(cs)
	}

	for _, c := range cs {
		if !c.dead {
			newCarts = append(newCarts, c)
		}
	}
	return newCarts
}

func display(r [][]rune, cs carts) {
	cartCoords := make(map[coord]*cart)
	for _, cart := range cs {
		cartCoords[cart.loc] = cart
	}
	cartColor := color.New(color.FgCyan).Add(color.Underline)
	for y := 0; y < len(railroad); y++ {
		for x := 0; x < len(railroad[y]); x++ {
			if v, ok := cartCoords[coord{x: x, y: y}]; ok {
				cartColor.Print(string(v.direction))
			} else {
				fmt.Print(string(railroad[y][x]))
			}
		}
		fmt.Println()
	}
}
