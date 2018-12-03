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
	alone := make(map[int]bool, 0)
	for _, l := range lines {
		var id, leftEdge, topEdge, width, heigth int
		fmt.Sscanf(l, "#%d @ %d,%d: %dx%d", &id, &leftEdge, &topEdge, &width, &heigth)
		for i := topEdge; i < topEdge+heigth; i++ {
			for j := leftEdge; j < leftEdge+width; j++ {
				fabric[i][j] = append(fabric[i][j], id)
				if len(fabric[i][j]) > 1 {
					for _, f := range fabric[i][j] {
						alone[f] = false
					}
				} else if len(fabric[i][j]) == 1 {
					if _, ok := alone[fabric[i][j][0]]; !ok {
						alone[fabric[i][j][0]] = true
					}
				}
			}
		}
	}

	for k, v := range alone {
		if v {
			fmt.Println(k)
			break
		}
	}
}
