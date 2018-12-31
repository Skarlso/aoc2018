package main

import "fmt"

func main() {
	s := 0
	a := 10551374 // gathered out of the output of part2/main.go
	for b := 1; b <= a; b++ {
		if a%b == 0 {
			s = s + b
		}
	}
	fmt.Printf("%d\n", s)
}
