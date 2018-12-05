package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	runes := []rune(string(content))
	for i := 0; i < len(runes)-1; i++ {
		curr := runes[i]
		next := runes[i+1]
		if curr != next && (unicode.ToUpper(curr) == next || unicode.ToLower(curr) == next) {
			runes = append(runes[:i], runes[i+2:]...)
			if i > 0 {
				i -= 2
			}
		}
	}
	fmt.Println(len(runes))
	fmt.Println(string(runes))
}
