package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	registers = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}
	// Initial values, these will change after we determine which is which.
	opcodes = map[int]func(a, b, c int){
		0:  addr,
		1:  addi,
		2:  mulr,
		3:  muli,
		4:  banr,
		5:  bani,
		6:  borr,
		7:  bori,
		8:  setr,
		9:  seti,
		10: gtir,
		11: gtri,
		12: gtrr,
		13: eqir,
		14: eqri,
		15: eqrr,
	}
)

func addr(a, b, c int) {
	registers[c] = registers[a] + registers[b]
}

func addi(a, b, c int) {
	registers[c] = registers[a] + b
}

func mulr(a, b, c int) {
	registers[c] = registers[a] * registers[b]
}

func muli(a, b, c int) {
	registers[c] = registers[a] * b
}

func banr(a, b, c int) {
	registers[c] = registers[a] & registers[b]
}

func bani(a, b, c int) {
	registers[c] = registers[a] & b
}

func borr(a, b, c int) {
	registers[c] = registers[a] | registers[b]
}

func bori(a, b, c int) {
	registers[c] = registers[a] | b
}

func setr(a, b, c int) {
	registers[c] = registers[a]
}

func seti(a, b, c int) {
	registers[c] = a
}

func gtir(a, b, c int) {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtri(a, b, c int) {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtrr(a, b, c int) {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqir(a, b, c int) {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqri(a, b, c int) {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqrr(a, b, c int) {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	var (
		before1, before2, before3, before4 int
		after1, after2, after3, after4     int
		opCode, a, b, c                    int
	)
	sampleCount := 0
	for _, l := range lines {
		if strings.Contains(l, "Before:") {
			fmt.Sscanf(l, "Before: [%d, %d, %d, %d]", &before1, &before2, &before3, &before4)
			continue
		}
		if !strings.Contains(l, "Before:") && !strings.Contains(l, "After:") {
			fmt.Sscanf(l, "%d %d %d %d", &opCode, &a, &b, &c)
			continue
		}
		if strings.Contains(l, "After:") {
			fmt.Sscanf(l, "After: [%d, %d, %d, %d]", &after1, &after2, &after3, &after4)
			possibleOpCodes := make(map[int]int)
			for _, v := range opcodes {
				registers[0] = before1
				registers[1] = before2
				registers[2] = before3
				registers[3] = before4
				v(a, b, c)
				if registers[0] == after1 &&
					registers[1] == after2 &&
					registers[2] == after3 &&
					registers[3] == after4 {
					possibleOpCodes[opCode]++
					if possibleOpCodes[opCode] > 2 {
						sampleCount++
						break
					}
				}
			}
			continue
		}
	}
	fmt.Println(sampleCount)
}
