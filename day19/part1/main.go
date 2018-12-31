package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	// Program Counter Register Index. Keep a separate track of the ip.
	pcri      = 0
	registers = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}
	// Initial values, these will change after we determine which is which.
	opcodes = map[string]func(a, b, c int){
		"addr": addr,
		"addi": addi,
		"mulr": mulr,
		"muli": muli,
		"banr": banr,
		"bani": bani,
		"borr": borr,
		"bori": bori,
		"setr": setr,
		"seti": seti,
		"gtir": gtir,
		"gtri": gtri,
		"gtrr": gtrr,
		"eqir": eqir,
		"eqri": eqri,
		"eqrr": eqrr,
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
	// pop the first line that's the ip assignment.
	ip, lines := lines[0], lines[1:]
	fmt.Sscanf(ip, "#ip %d", &pcri)
	fmt.Println("Instruction Pointer: ", pcri)
	for registers[pcri] < len(lines) {
		// not all instructions modify the pointer
		registers[pcri]++
		if registers[pcri] > len(lines) {
			fmt.Println("outside the instructions range.")
			break
		}

		var (
			op      string
			a, b, c int
		)
		fmt.Sscanf(lines[registers[pcri]], "%s %d %d %d", &op, &a, &b, &c)
		opcodes[op](a, b, c)
	}
	fmt.Println("Register 0: ", registers)
}
