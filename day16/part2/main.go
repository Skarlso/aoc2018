package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type op struct {
	name string
	v    func(a, b, c uint64)
}

var (
	registers = map[uint64]uint64{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}
	// Initial values, these will change after we determine which is which.
	opcodes = map[uint64]op{
		0:  op{v: addr, name: "addr"},
		1:  op{v: addi, name: "addi"},
		2:  op{v: mulr, name: "mulr"},
		3:  op{v: muli, name: "muli"},
		4:  op{v: banr, name: "banr"},
		5:  op{v: bani, name: "bani"},
		6:  op{v: borr, name: "borr"},
		7:  op{v: bori, name: "bori"},
		8:  op{v: setr, name: "setr"},
		9:  op{v: seti, name: "seti"},
		10: op{v: gtir, name: "gtir"},
		11: op{v: gtri, name: "gtri"},
		12: op{v: gtrr, name: "gtrr"},
		13: op{v: eqir, name: "eqir"},
		14: op{v: eqri, name: "eqri"},
		15: op{v: eqrr, name: "eqrr"},
	}
	possibleOpCodes = make(map[uint64]op)
)

func addr(a, b, c uint64) {
	registers[c] = registers[a] + registers[b]
}

func addi(a, b, c uint64) {
	registers[c] = registers[a] + b
}

func mulr(a, b, c uint64) {
	registers[c] = registers[a] * registers[b]
}

func muli(a, b, c uint64) {
	registers[c] = registers[a] * b
}

func banr(a, b, c uint64) {
	registers[c] = registers[a] & registers[b]
}

func bani(a, b, c uint64) {
	registers[c] = registers[a] & b
}

func borr(a, b, c uint64) {
	registers[c] = registers[a] | registers[b]
}

func bori(a, b, c uint64) {
	registers[c] = registers[a] | b
}

func setr(a, b, c uint64) {
	registers[c] = registers[a]
}

func seti(a, b, c uint64) {
	registers[c] = a
}

func gtir(a, b, c uint64) {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtri(a, b, c uint64) {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtrr(a, b, c uint64) {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqir(a, b, c uint64) {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqri(a, b, c uint64) {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqrr(a, b, c uint64) {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func main() {
	// Part 1:
	filename := "input1.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	var (
		before1, before2, before3, before4 uint64
		after1, after2, after3, after4     uint64
		opCode, a, b, c                    uint64
	)
	// sampleCount := 0
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
			matched := 0
			var code uint64
			var f func(a, b, c uint64)
			for k, v := range opcodes {
				registers[0] = before1
				registers[1] = before2
				registers[2] = before3
				registers[3] = before4
				v.v(a, b, c)
				// matched 1 and not set
				if registers[0] == after1 &&
					registers[1] == after2 &&
					registers[2] == after3 &&
					registers[3] == after4 {
					matched++
					code = k
					f = v.v
				}
			}
			if matched == 1 {
				if _, ok := possibleOpCodes[code]; !ok {
					possibleOpCodes[code] = op{name: "asdf", v: f}
				}
			}
			continue
		}
	}
	fmt.Println(possibleOpCodes)
	// part 2
	f := "input2.txt"
	con, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	ls := strings.Split(string(con), "\n")
	registers = map[uint64]uint64{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}
	for _, l := range ls {
		fmt.Println(registers)
		var (
			code, a2, b2, c2 uint64
		)
		fmt.Sscanf(l, "%d %d %d %d", &code, &a2, &b2, &c2)
		possibleOpCodes[code].v(a2, b2, c2)
	}
	fmt.Println("Register 0: ", registers[0])
}
