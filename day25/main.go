package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type point struct {
	x, y, z, s int
	chainID    *point
}

func (p point) dist(other point) int {
	return abs(p.x-other.x) + abs(p.y-other.y) + abs(p.z-other.z) + abs(p.s-other.s)
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	run(content)
}

func run(content []byte) {
	chains := make(map[*point][]*point)
	lines := strings.Split(string(content), "\n")
	points := make([]*point, 0)
	for _, l := range lines {
		var (
			x, y, z, s int
		)
		_, _ = fmt.Sscanf(l, "%d,%d,%d,%d", &x, &y, &z, &s)
		p := point{x: x, y: y, z: z, s: s, chainID: nil}
		points = append(points, &p)
	}
	for _, p1 := range points {
		for _, p2 := range points {
			if p2.dist(*p1) > 3 {
				continue
			}

			if p1.chainID != nil {
				//// Our point is currently in a constellation
				if p2.chainID == nil {
					// p2 is not in constellation
					chains[p1.chainID] = append(chains[p1.chainID], p2)
					p2.chainID = p1.chainID
				} else {
					// They are in the same constellation
					if p1.chainID == p2.chainID {
						continue
					}
					// We merge p2's constellation into p1's
					oldId := p2.chainID
					// Update all of p2's points to point to p1
					for _, p := range chains[p2.chainID] {
						p.chainID = p1.chainID
					}
					chains[p1.chainID] = append(chains[p1.chainID], chains[oldId]...)
					// Remove p2's old constellation
					delete(chains, oldId)
					p2.chainID = p1.chainID
				}
			} else {
				// p1 is not in a constellation
				if p2.chainID == nil {
					// p2 is not in a constellation either, we create a new one
					chains[p1] = make([]*point, 0)
					chains[p1] = append(chains[p1], p1, p2)
					p1.chainID = p1
					p2.chainID = p1
				} else {
					// p2 is already in a constellation we join it
					chains[p2.chainID] = append(chains[p2.chainID], p1)
					p1.chainID = p2
				}
			}
		}
	}
	fmt.Println(len(chains))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
