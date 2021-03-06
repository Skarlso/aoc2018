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
}

type carts []*cart

// Move moves the given cart into the direction that it's currently
// facing. If the cart encounters an intersection it will change
// its current heading.
func (c *cart) move(cs carts) (int, int, bool) {
	c.loc.x += directions[c.direction].x
	c.loc.y += directions[c.direction].y

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
			fmt.Println("CRASH!")
			return c.loc.x, c.loc.y, true
		}
	}
	return c.loc.x, c.loc.y, false
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
	collisionX := 0
	collisionY := 0
	ticks := 0
	for {
		ticks++
		collisionX, collisionY, crash = moveCarts(cs)
		if crash {
			break
		}
		// display(railroad, cs)
		// time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(ticks)
	fmt.Println(collisionX, collisionY)
}

func showCarts(cs carts) {
	for _, c := range cs {
		fmt.Println(c)
	}
}

func moveCarts(cs carts) (x int, y int, collision bool) {
	for _, c := range cs {
		x, y, collision = c.move(cs)
		if collision {
			return x, y, collision
		}
	}
	return
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
