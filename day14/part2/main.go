package main

import (
	"fmt"
	"strconv"
)

var find = "293801"

var scoreboard = make([]int, 0)

func main() {
	scoreboard = append(scoreboard, 3, 7)
	elfOne := 0
	elfTwo := 1
	count := 1
	for {
		count++
		// create new recipe
		score := scoreboard[elfOne] + scoreboard[elfTwo]
		scoreDigits := strconv.Itoa(score)
		for _, s := range scoreDigits {
			i, _ := strconv.Atoi(string(s))
			scoreboard = append(scoreboard, i)
			if contains(find) {
				fmt.Println(len(scoreboard[:len(scoreboard)-len(find)]))
				return
			}
		}

		// pick new recipes
		elfOne = ((scoreboard[elfOne] + 1) + elfOne) % len(scoreboard)
		elfTwo = ((scoreboard[elfTwo] + 1) + elfTwo) % len(scoreboard)

	}
}

func contains(l string) bool {
	if len(scoreboard) < len(l) {
		return false
	}
	s := ""
	for _, i := range scoreboard[len(scoreboard)-len(l):] {
		s += strconv.Itoa(i)
	}
	return s == l
}
