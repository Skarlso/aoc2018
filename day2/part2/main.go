package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	for i, l := range lines {
		j := i + 1
		for j < len(lines) {
			l2 := lines[j]
			if ok, v := difference(l, l2); ok {
				fmt.Println(v)
				return
			}
			j++
		}
	}
}

func difference(a, b string) (bool, string) {
	counter := 0
	same := []rune{}
	for i, letter := range a {
		if letter != rune(b[i]) {
			counter++
		} else {
			same = append(same, letter)
		}
	}
	if counter == 1 {
		return true, string(same)
	}
	return false, ""
}
