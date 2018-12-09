package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var currentPosition = 0

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	c := string(content)
	split := strings.Split(c, " ")
	ints := convertAllToInt(split)
	fmt.Println(parseMetadata(ints))
}

func parseMetadata(ints []int) (sum int) {
	if ints[currentPosition] == 0 {
		currentPosition++
		m := ints[currentPosition]
		currentPosition++
		for i := 0; i < m; i++ {
			sum += ints[currentPosition+i]
		}
		currentPosition += m
		return sum
	}

	n := currentPosition
	currentPosition++
	m := ints[currentPosition]
	currentPosition++
	for i := 0; i < ints[n]; i++ {
		sum += parseMetadata(ints)
	}

	for i := 0; i < m; i++ {
		sum += ints[currentPosition+i]
	}
	currentPosition += m

	return sum
}

func convertAllToInt(s []string) []int {
	n := make([]int, 0)
	for _, v := range s {
		k, _ := strconv.Atoi(v)
		n = append(n, k)
	}
	return n
}
