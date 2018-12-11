package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

type star struct {
	x  int
	y  int
	vx int
	vy int
}

func (s *star) add() {
	s.x += s.vx
	s.y += s.vy
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")

	maxX := 0
	maxY := 0
	minX := 10000000
	minY := 10000000

	for _, l := range lines {
		var x, y, vx, vy int
		fmt.Sscanf(l, "position=<%d,  %d> velocity=<%d,  %d>", &x, &y, &vx, &vy)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if abs(x) < minX {
			minX = x
		}
		if abs(y) < minY {
			minY = y
		}
	}

	startX := maxX
	startY := maxY

	stars := make([]*star, 0)
	for _, l := range lines {
		var x, y, vx, vy int
		fmt.Sscanf(l, "position=<%d,  %d> velocity=<%d,  %d>", &x, &y, &vx, &vy)
		s := star{
			x:  startX + x,
			y:  startY + y,
			vx: vx,
			vy: vy,
		}
		stars = append(stars, &s)
	}

	alligned := false
	seconds := 0
	for !alligned {
		mx, my, max, may := boundries(stars)
		for _, s := range stars {
			if s.x+s.vx > max || s.y+s.vy > may || s.x+s.vx < mx || s.y+s.vy < my {
				alligned = true
				minX = mx
				minY = my
				maxX = max
				maxY = may
				break
			}
		}
		if !alligned {
			for _, s := range stars {
				s.add()
			}
		}
		seconds++
	}
	fmt.Println("seconds: ", seconds-1)
	writeSvg(minX, maxX, minY, maxY, stars)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func boundries(sky []*star) (minx, miny, maxx, maxy int) {
	maxX := 0
	maxY := 0
	minX := 10000000
	minY := 10000000

	for _, s := range sky {
		if s.x > maxX {
			maxX = s.x
		}
		if s.y > maxY {
			maxY = s.y
		}
		if abs(s.x) < minX {
			minX = s.x
		}
		if abs(s.y) < minY {
			minY = s.y
		}
	}
	return minX, minY, maxX, maxY
}

func writeSvg(minx, maxx, miny, maxy int, coords []*star) {
	fmt.Println(minx, maxx, miny, maxy)
	header := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, (maxx - minx), (maxy - miny))
	f, err := os.Create("sky.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(header)
	f.WriteString(fmt.Sprintf(`<rect x="0" y="0" width="%d" height="%d" fill="black" />`, maxx-minx, maxy-miny))
	for _, c := range coords {
		f.WriteString(fmt.Sprintf(`<circle cx="%d" cy="%d" r="0.5" fill="red"/>`, c.x-minx, c.y-miny))
	}
	f.WriteString("</svg>")
}
