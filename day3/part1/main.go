package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	size := os.Args[2]
	s, _ := strconv.Atoi(size)
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	fabric := make([][][]int, s)
	for i := range fabric {
		fabric[i] = make([][]int, s)
		for j := range fabric[i] {
			fabric[i][j] = make([]int, 0)
		}
	}
	// matrix of IDs. if len(id) > 1 -- count it.
	for _, l := range lines {
		var id, leftEdge, topEdge, width, heigth int
		fmt.Sscanf(l, "#%d @ %d,%d: %dx%d", &id, &leftEdge, &topEdge, &width, &heigth)
		for i := topEdge; i < topEdge+heigth; i++ {
			for j := leftEdge; j < leftEdge+width; j++ {
				fabric[i][j] = append(fabric[i][j], id)
			}
		}
	}

	square := 0
	// for i := range fabric {
	// 	for j := range fabric[i] {
	// 		if len(fabric[i][j]) > 1 {
	// 			fmt.Print("X")
	// 		} else if len(fabric[i][j]) == 1 {
	// 			fmt.Print(fabric[i][j][0])
	// 		} else if len(fabric[i][j]) == 0 {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	for i := range fabric {
		for j := range fabric[i] {
			if len(fabric[i][j]) == 1 {
				square++
			}
		}
	}
	fmt.Println(square)
}
