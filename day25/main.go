package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type point struct {
	x, y, z, s int
	// key is the key of the chain that this point is a part off
	key *point
}

func (p point) dist(other point) int {
	d := math.Sqrt(float64((p.x-other.x)*(p.x-other.x) +
		(p.y-other.y)*(p.y-other.y) +
		(p.z-other.z)*(p.z-other.z) +
		(p.s-other.s)*(p.s-other.s)))
	return int(d)
}

func (p point) equal(other point) bool {
	return p.x == other.x && p.y == other.y && p.z == other.z && p.s == other.s
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	run(content)
}

func run(content []byte) {
	chains := make(map[*point][]point)
	//chains := make([]*[]point, 0)
	lines := strings.Split(string(content), "\n")
	points := make([]*point, 0)
	for _, l := range lines {
		var (
			x, y, z, s int
		)
		_, _ = fmt.Sscanf(l, "%d,%d,%d,%d", &x, &y, &z, &s)
		p := point{x: x, y: y, z: z, s: s}
		points = append(points, &p)
	}
	for _, p1 := range points {
		for _, p2 := range points {
			if p1.equal(*p2) {
				continue
			}
			// The encountered point is in a chain... We join that chain.
			if p1.dist(*p2) < 4 {
				if _, ok2 := chains[p2]; ok2 {
					if _, ok := chains[p1]; !ok {
						chains[p2] = append(chains[p2], *p1)
						p1.key = p2
					} else if p1.key == nil && ok {
						chains[p2] = append(chains[p2], chains[p1]...)
						delete(chains, p1)
						p1.key = p2
					} else if p1.key != nil && !ok {
						chains[p2] = append(chains[p2], chains[p1.key]...)
						delete(chains, p1.key)
						p1.key = p2
					}
				} else if p2.key == nil && ok2 {
					if _, ok1 := chains[p1]; !ok1 {
						chains[p1] = make([]point, 0)
					}
					chains[p1] = append(chains[p1], *p2)
					p2.key = p1
				} else if p2.key != nil && !ok2 {
					chains[p1] = append(chains[p1], chains[p2.key]...)
					delete(chains, p2.key)
					p2.key = p1
				}
			}
		}
	}
	fmt.Println(len(chains))
	display(points)
}

func display(points []*point) {
	for _, p := range points {
		fmt.Println(p)
	}
}
