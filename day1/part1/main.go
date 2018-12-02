package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	count := 0
	content, _ := ioutil.ReadFile("input.txt")
	lines := bytes.Split(content, []byte("\n"))
	var sum int
	seen := make(map[int]bool)
	for {
		for _, l := range lines {
			count++
			var c int
			c, _ = strconv.Atoi(string(l))
			sum += c
			if seen[sum] {
				fmt.Println(sum)
				fmt.Println(count)
				return
			}
			seen[sum] = true
		}
	}
}
