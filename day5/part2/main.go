package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	ls := make([]int, 0)
	for a := 'a'; a <= 'z'; a++ {
		c := string(content)
		if !strings.ContainsRune(c, a) {
			continue
		}
		c = strings.Replace(c, string(a), "", -1)
		c = strings.Replace(c, string(unicode.ToUpper(a)), "", -1)
		runes := []rune(c)
		reacted := react(runes)
		ls = append(ls, len(reacted))
	}
	sort.Ints(ls)
	fmt.Println(ls)
}

func react(runes []rune) []rune {
	for i := 0; i < len(runes)-1; i++ {
		curr := runes[i]
		next := runes[i+1]
		if curr != next && (unicode.ToUpper(curr) == next || unicode.ToLower(curr) == next) {
			runes = append(runes[:i], runes[i+2:]...)
			if i > 0 {
				i -= 2
			} else {
				i--
			}
		}
	}
	return runes
}
