package main

import (
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	//filename := os.Args[1]
	content, _ := ioutil.ReadFile("input.txt")
	run(content)
}
