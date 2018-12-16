package main

import (
	"fmt"
	"strconv"
)

const (
	limit = 293801
	// limit = 2018
)

var scoreboard = make([]int, 0)

func main() {
	scoreboard = append(scoreboard, 3, 7)
	elfOne := 0
	elfTwo := 1
	over := 0
	for len(scoreboard) < limit+10 {
		// create new recipe
		score := scoreboard[elfOne] + scoreboard[elfTwo]
		scoreDigits := strconv.Itoa(score)
		for _, s := range scoreDigits {
			i, _ := strconv.Atoi(string(s))
			scoreboard = append(scoreboard, i)
			if len(scoreboard) > limit {
				over++
			}
			if over == 10 {
				fmt.Println(scoreboard[len(scoreboard)-10:])
				return
			}
		}
		// pick new recipes
		elfOne = ((scoreboard[elfOne] + 1) + elfOne) % len(scoreboard)
		elfTwo = ((scoreboard[elfTwo] + 1) + elfTwo) % len(scoreboard)

	}
	// fmt.Println(scoreboard)
}
