package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	twos := 0
	threes := 0
	for _, l := range lines {
		hasTwos := false
		hasThrees := false

		letters := make(map[string]int, 0)
		for _, c := range strings.Split(l, "") {
			letters[c]++
		}

		for _, v := range letters {
			if v == 3 {
				hasThrees = true
			}
			if v == 2 {
				hasTwos = true
			}
		}
		if hasThrees {
			threes++
		}
		if hasTwos {
			twos++
		}
	}

	fmt.Println(twos * threes)
}
