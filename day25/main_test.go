package main

import "testing"

func TestRun(t *testing.T) {
	content := []byte(` 0,0,0,0
 3,0,0,0
 0,3,0,0
 0,0,3,0
 0,0,0,3
 0,0,0,6
 9,0,0,0
12,0,0,0`)
	run(content)
}
