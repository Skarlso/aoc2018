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

	// zeroLocation := strings.Index(plants, "#")
	g := 0
	negativOffset := 0
	plantRunes := []rune(plants)
	fmt.Println(plants)
	for g < generations {
		plantRunes = append(plantRunes, []rune("..")...)
		plantRunes = append([]rune(".."), plantRunes...)
		negativOffset -= 2

		newGeneration := make([]rune, len(plantRunes))
		copy(newGeneration, plantRunes)
		for i := 2; i < len(plantRunes)-2; i++ {
			match := string(plantRunes[i-2]) + string(plantRunes[i-1]) + string(plantRunes[i]) + string(plantRunes[i+1]) + string(plantRunes[i+2])
			if v, ok := rules[match]; ok {
				newGeneration[i] = []rune(v)[0]
			}
		}
		plantRunes = newGeneration
		g++
	}

	sum := 0
	for i, v := range plantRunes {
		if v == '#' {
			sum += i + negativOffset
		}
	}
	fmt.Println(sum)
}
