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
	fabric := make([][][]int, 1000)
	for i := range fabric {
		fabric[i] = make([][]int, 1000)
	}
	// matrix of IDs. if len(id) > 1 -- count it.
	for _, l := range lines {
		var id, leftEdge, topEdge, width, heigth int
		fmt.Sscanf(l, "#%d @ %d,%d: %dx%d", id, leftEdge, topEdge, width, heigth)
	}
}
