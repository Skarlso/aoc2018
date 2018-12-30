package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// read all numbers to a []int
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")

	ground := make([][]rune, 1)
	maxX, minX, maxY, minY := 0, 0, 0, 20
	xOffset, yOffset := 500, 0

	for _, l := range lines {
		switch l[0] {
		case 'x':
			var (
				x, y1, y2 int
				r         string
			)
			fmt.Sscanf(l, "x=%d, y=%s", &x, &r)
			fmt.Sscanf(r, "%d..%d", &y1, &y2)
			x -= xOffset
			y1 -= yOffset
			y2 -= yOffset
			for x >= maxX {
				maxX++
				for j := range ground {
					ground[j] = append(ground[j], '.')
				}
			}
			for x <= minX {
				minX--
				for j := range ground {
					ground[j] = append([]rune{'.'}, ground[j]...)
				}
			}
			for y2 > maxY {
				maxY++
				ground = append(ground, make([]rune, len(ground[0])))
				for j := range ground[len(ground)-1] {
					ground[len(ground)-1][j] = '.'
				}
			}
			if y1 < minY {
				minY = y1
			}
			for i := y1; i <= y2; i++ {
				ground[i][x-minX] = '#'
			}
		case 'y':
			var (
				x1, x2, y int
				// range
				r string
			)
			fmt.Sscanf(l, "y=%d, x=%s", &y, &r)
			fmt.Sscanf(r, "%d..%d", &x1, &x2)
			x1 -= xOffset
			x2 -= xOffset
			y -= yOffset
			for y > maxY {
				maxY++
				ground = append(ground, make([]rune, len(ground[0])))
				for j := range ground[len(ground)-1] {
					ground[len(ground)-1][j] = '.'
				}
			}
			for x2 >= maxX {
				maxX++
				for j := range ground {
					ground[j] = append(ground[j], '.')
				}
			}
			for x1 <= minX {
				minX--
				for j := range ground {
					ground[j] = append([]rune{'.'}, ground[j]...)
				}
			}
			for i := x1; i <= x2; i++ {
				ground[y][i-minX] = '#'
			}
			if y < minY {
				minY = y
			}
		}
	}

	waterCount := 0
	flowCount := 0
	roundLimit := 200000

	for ground[1][-minX] != '|' && waterCount < roundLimit {
		canMove := true
		x := -minX
		y := 1
		tryLeft := 0
		for canMove {
			if y+1 > maxY || ground[y+1][x] == '|' {
				ground[y][x] = '|'
				canMove = false
				if y >= minY {
					flowCount++
				}
			} else if ground[y+1][x] == '.' {
				y++
				tryLeft = 0
			} else if ground[y+1][x] == '#' || ground[y+1][x] == '~' {
				if (tryLeft == 1 && ground[y][x-1] == '|') ||
					(tryLeft == 2 && ground[y][x+1] == '|') ||
					(ground[y][x+1] == '|' && ground[y][x-1] != '.') ||
					(ground[y][x+1] != '.' && ground[y][x-1] == '|') {
					ground[y][x] = '|'
					flowCount++
					canMove = false
					for i := x + 1; ground[y][i] == '~'; i++ {
						ground[y][i] = '|'
						waterCount--
						flowCount++
					}
					for i := x - 1; ground[y][i] == '~'; i-- {
						ground[y][i] = '|'
						waterCount--
						flowCount++
					}
				} else if (tryLeft == 0 && ground[y][x-1] == '.') ||
					(tryLeft == 1 && ground[y][x-1] == '.') {
					x--
					tryLeft = 1
				} else if (tryLeft == 0 && ground[y][x+1] == '.') ||
					(tryLeft == 2 && ground[y][x+1] == '.') {
					x++
					tryLeft = 2
				} else {
					canMove = false
					ground[y][x] = '~'
					waterCount++
				}
			}

		}

	}

	// for j := range ground {
	// 	for _, v := range ground[j] {
	// 		fmt.Print(string(v))
	// 	}
	// 	fmt.Println()
	// }
	fmt.Println("Standing:", waterCount, "Flowing:", flowCount, "Total:", flowCount+waterCount)
}

func display(r [][]rune) {
	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[y]); x++ {
			fmt.Print(string(r[y][x]))
		}
		fmt.Println()
	}
}
