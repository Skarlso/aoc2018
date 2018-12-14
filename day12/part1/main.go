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

type plant struct {
	next  *plant
	prev  *plant
	value rune
	index int
}

func (p *plant) add(value rune, index int) *plant {
	if index < 0 {
		p.prev = &plant{next: p, prev: nil, value: value, index: index}
	} else {
		p.next = &plant{next: nil, prev: p, value: value, index: index}
	}
	return p
}

func (p *plant) insert(v rune, i int) *plant {
	newPlant := plant{value: v, index: i, next: nil, prev: nil}
	nextPlant := p.next
	p.next, nextPlant.prev = &newPlant, &newPlant
	newPlant.prev, newPlant.next = p, nextPlant
	return &newPlant
}

func (p *plant) evolveLeft() *plant {
	return p
}

func (p *plant) evolveRight() *plant {
	return p
}

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

	firstPlantLocation := strings.Index(plants, "#")
	hallway := &plant{
		next:  nil,
		prev:  nil,
		value: '#',
		index: firstPlantLocation,
	}
	// constructing plants
	for i, r := range plants {
		if i < firstPlantLocation {
			p := plant{
				next: nil,
				prev:
			}
			continue
		}

	}

	// g := 0
	// if firstPlantLocation < 5 {
	// 	plants = "..." + plants
	// }
	// for g < generations {
	// 	for i, r := range plants {
	// 		if r == '#' {
	// 			if i+2 > len(plants) {
	// 				plants += ".."
	// 			}
	// 			if i-2 < 0 {
	// 				plants = ".." + plants
	// 			}

	// 		}
	// 	}
	// }
}
