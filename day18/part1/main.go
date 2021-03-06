package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

var field = make([][]rune, 0)
var maxX, maxY int

func main() {
	// read all numbers to a []int
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	for _, l := range lines {
		r := []rune(l)
		maxX = len(r)
		field = append(field, r)
	}

	maxY = len(field)

	minutes := 0
	for minutes < 131 {
		newField := make([][]rune, maxY)
		for i := 0; i < maxX; i++ {
			newField[i] = make([]rune, maxX)
		}
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				neighbours := getNeighbours(x, y)
				currentAcre := field[y][x]
				if currentAcre == '.' {
					if neighbours['|'] >= 3 {
						currentAcre = '|'
					}
				} else if currentAcre == '|' {
					if neighbours['#'] >= 3 {
						currentAcre = '#'
					}
				} else if currentAcre == '#' {
					if neighbours['#'] == 0 || neighbours['|'] == 0 {
						currentAcre = '.'
					}
				}
				newField[y][x] = currentAcre
			}
		}
		minutes++
		field = newField
	}

	lumber := 0
	tree := 0
	for _, i := range field {
		for _, j := range i {
			if j == '|' {
				tree++
			}
			if j == '#' {
				lumber++
			}
		}
	}

	fmt.Println("tree * lumber: ", tree*lumber)
}

func getNeighbours(x, y int) map[rune]int {
	types := make(map[rune]int)

	if x+1 < maxX && y+1 < maxY {
		types[field[y+1][x+1]]++
	}
	if x-1 >= 0 && y+1 < maxY {
		types[field[y+1][x-1]]++
	}
	if y+1 < maxY {
		types[field[y+1][x]]++
	}
	if x+1 < maxX {
		types[field[y][x+1]]++
	}
	if x+1 < maxX && y-1 >= 0 {
		types[field[y-1][x+1]]++
	}
	if y-1 >= 0 {
		types[field[y-1][x]]++
	}
	if x-1 >= 0 && y-1 >= 0 {
		types[field[y-1][x-1]]++
	}
	if x-1 >= 0 {
		types[field[y][x-1]]++
	}
	// lumberyard := types['#']
	// open := types['.']
	// trees := types['|']
	// fmt.Printf("x:%d y:%d lumberyard:%d trees:%d open:%d\n", x, y, lumberyard, trees, open)
	return types
}

func display(r [][]rune) {
	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[y]); x++ {
			fmt.Print(string(r[y][x]))
		}
		fmt.Println()
	}
}
