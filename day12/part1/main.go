package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	generations = 5
)

// TODO: Consider re-writing the whole thing as linked list because shuffling
// a huge String is really crap.
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

	zeroLocation := strings.Index(plants, "#")
	g := 0
	negativOffset := 0
	plantRunes := []rune(plants)
	fmt.Println("PlantRunes length: ", len(plantRunes))
	for g < generations {
		begin := strings.Index(string(plantRunes), "#")
		end := strings.LastIndex(string(plantRunes), "#") + 1
		if end+2 > len(plantRunes) && end+1 <= len(plantRunes) {
			plantRunes = append(plantRunes, []rune("..")...)
			// end--
		} else if end+1 > len(plantRunes) {
			plantRunes = append(plantRunes, []rune("..")...)
			// end--
		}
		if begin-2 < 0 {
			plantRunes = append([]rune(".."), plantRunes...)
			negativOffset += 2
			zeroLocation += 2
			begin += 2
		} else if begin-1 < 0 {
			plantRunes = append([]rune(".."), plantRunes...)
			negativOffset++
			zeroLocation++
		}
		fmt.Println(string(plantRunes))
		fmt.Println(begin, end)
		newGeneration := make([]rune, len(plantRunes))
		copy(newGeneration, plantRunes)
		for i := begin; i <= end; i++ {
			match := string(plantRunes[i-2]) + string(plantRunes[i-1]) + string(plantRunes[i]) + string(plantRunes[i+1]) + string(plantRunes[i+2])
			if v, ok := rules[match]; ok {
				newGeneration[i] = []rune(v)[0]
			} else {
				newGeneration[i] = '.'
			}
		}
		plantRunes = newGeneration
		fmt.Println(string(plantRunes))
		g++
	}
}
