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

	// calculate when it starts to begin again.
	prev := make(map[int]int, 0)
	minutes := 0
	found := false
	prevMinuteMark := 0
	for minutes < 1000000000 {

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
		field = newField
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
		if _, ok := prev[lumber*tree]; ok {
			if found {
				fmt.Println("ocured again at: ", minutes)
				break
			}
			fmt.Println("started to recur at minute: ", minutes)
			found = true
			prevMinuteMark = minutes
		}

		prev[lumber*tree] = minutes
		minutes++
	}
	freq := minutes - prevMinuteMark
	loc := abs(1000000000-minutes) % freq
	fmt.Println(1000000000 - minutes)
	fmt.Println("resource: ", loc)
	fmt.Println(prev)
	// fmt.Println("Freq: ", minutes-prevMinuteMark)
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
