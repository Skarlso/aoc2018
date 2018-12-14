package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	generations = 20
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	var plants string
	rules := make(map[string]string)
	for _, l := range lines {
		if strings.Contains(l, "initial state:") {
			fmt.Sscanf(l, "initial state: %s", &plants)
			continue
		}
		var rule, outcome string
		fmt.Sscanf(l, "%s => %s", &rule, &outcome)
		if len(rule) > 0 {
			rules[rule] = outcome
		}
	}

	g := 0
	offset := 0
	fmt.Println(plants)
	for g < generations {
		plants = ".." + plants + ".."
		offset += 2
		newGeneration := plants
		for i := 2; i < len(plants)-2; i++ {
			match := plants[i-2 : i+3]
			runes := []rune(newGeneration)
			if v, ok := rules[match]; ok {
				runes[i] = []rune(v)[0]
			} else {
				runes[i] = '.'
			}
			newGeneration = string(runes)
		}
		plants = newGeneration
		g++
	}

	sum := 0
	for i, v := range plants {
		if v == '#' {
			sum += (i - offset)
		}
	}
	fmt.Println(sum)
}
