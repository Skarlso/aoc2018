package main

import "fmt"

type coord struct {
	x int
	y int
}

type region struct {
	geoindex int
	erosion  int
	t        int
}

const (
	rocky = iota
	wet
	narrow
)

var (
	depth  = 5355
	target = coord{x: 14, y: 796}
	// depth  = 510
	// target = coord{x: 10, y: 10}
)

func main() {
	cave := make([][]region, target.y+1)
	risk := 0
	for y := 0; y <= target.y; y++ {
		cave[y] = make([]region, target.x+1)
		for x := 0; x <= target.x; x++ {
			if y == 0 && x == 0 {
				cave[y][x] = region{geoindex: 0, t: 0, erosion: 0}
				continue
			}

			if y == target.y && x == target.x {
				cave[y][x] = region{geoindex: 0, t: 0, erosion: 0}
				continue
			}

			var r region
			if y == 0 {
				r = region{
					geoindex: (x * 16807),
					erosion:  0,
					t:        0,
				}
			} else if x == 0 {
				r = region{
					geoindex: (y * 48271),
					erosion:  0,
					t:        0,
				}
			} else {
				r = region{
					geoindex: cave[y][x-1].erosion * cave[y-1][x].erosion,
					erosion:  0,
					t:        0,
				}
			}
			r.erosion = (r.geoindex + depth) % 20183
			r.t = r.erosion % 3
			cave[y][x] = r
			risk += r.t
		}
	}

	fmt.Println("risk: ", risk)
}
