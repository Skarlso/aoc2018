package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	fabric := make([][][]int, 10)
	for i := range fabric {
		fabric[i] = make([][]int, 10)
		for j := range fabric[i] {
			fabric[i][j] = make([]int, 0)
		}
	}
	// matrix of IDs. if len(id) > 1 -- count it.
	for _, l := range lines {
		var id, leftEdge, topEdge, width, heigth int
		fmt.Sscanf(l, "#%d @ %d,%d: %dx%d", &id, &leftEdge, &topEdge, &width, &heigth)
		for i := leftEdge; i < width; i++ {
			for j := topEdge; j < heigth; j++ {
				fabric[i][j] = append(fabric[i][j], id)
			}
		}
	}

	for i := range fabric {
		for j := range fabric[i] {
			if len(fabric[i][j]) > 2 {
				fmt.Println("claimed by more than one.")
			}
		}
	}
}
