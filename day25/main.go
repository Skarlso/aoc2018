package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type point struct {
	x, y, z, s int
	// key is the key of the chain that this point is a part off
	chainID int
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
	chains := make(map[int][]*point)
	//chains := make([]*[]point, 0)
	lines := strings.Split(string(content), "\n")
	points := make([]*point, 0)
	for i, l := range lines {
		var (
			x, y, z, s int
		)
		_, _ = fmt.Sscanf(l, "%d,%d,%d,%d", &x, &y, &z, &s)
		p := point{x: x, y: y, z: z, s: s, chainID: i}
		points = append(points, &p)
	}
	for _, p1 := range points {
		for _, p2 := range points {
			if p2.dist(*p1) > 3 {
				continue
			}

			if _, ok2 := chains[p2.chainID]; ok2 {
				if _, ok := chains[p1.chainID]; !ok {
					chains[p2.chainID] = append(chains[p2.chainID], p1)
					p1.chainID = p2.chainID
				} else {
					if p1.chainID == p2.chainID {
						continue
					}
					oldId := p1.chainID
					for _, e := range chains[p1.chainID] {
						e.chainID = p2.chainID
					}
					chains[p2.chainID] = append(chains[p2.chainID], chains[p1.chainID]...)
					delete(chains, oldId)
					p1.chainID = p2.chainID // although I think this already will be updated in the for
				}
			} else {
				chains[p1.chainID] = append(chains[p1.chainID], p2)
				p2.chainID = p1.chainID
			}
		}
	}
	fmt.Println(len(chains))
	//fmt.Println(chains)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
