package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	for i, l := range lines {
		j := i + 1
		for j < len(lines) {
			l2 := lines[j]
			split1 := strings.Split(l, "")
			split2 := strings.Split(l2, "")
			// fmt.Println(diff)
			if ok, v := difference(split1, split2); ok {
				fmt.Println(strings.Join(v, ""))
				return
			}
			j++
		}
	}

}

func difference(a, b []string) (bool, []string) {
	counter := 0
	same := []string{}
	for i, letter := range a {
		if letter != b[i] {
			counter++
		} else {
			same = append(same, letter)
		}
	}
	if counter == 1 {
		return true, same
	}
	return false, []string{}
}
