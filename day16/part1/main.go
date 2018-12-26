package main

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
		2:  multr,
		3:  multi,
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

func multr(a, b, c int) {
	registers[c] = registers[a] * registers[b]
}

func multi(a, b, c int) {
	registers[c] = registers[a] + b
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

}
