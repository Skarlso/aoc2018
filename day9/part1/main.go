package main

import (
	"fmt"
)

const (
	lastMarbleValue = 7185200
	playerCount     = 404
)

type marbleCircle struct {
	next *marbleCircle
	prev *marbleCircle
	v    int
}

func (mc *marbleCircle) insert(v int) *marbleCircle {
	newMarbleCircle := marbleCircle{v: v}
	nextCricle := mc.next
	mc.next, nextCricle.prev = &newMarbleCircle, &newMarbleCircle
	newMarbleCircle.prev, newMarbleCircle.next = mc, nextCricle
	return &newMarbleCircle
}
func (mc *marbleCircle) delete() *marbleCircle {
	mc.prev.next = mc.next
	mc.next.prev = mc.next
	return mc.next
}

func main() {
	score := make([]int, playerCount)
	circle := &marbleCircle{v: 0}
	circle.next = circle
	circle.prev = circle

	player := 1
	for i := 1; i <= lastMarbleValue; i++ {
		if i%23 == 0 {
			for j := 0; j < 7; j++ {
				circle = circle.prev
			}
			score[player] += (i + circle.v)
			circle = circle.delete()
		} else {
			circle = circle.next.insert(i)
		}
		player = (player + 1) % playerCount
	}

	max := 0
	winner := 0
	for i := 0; i < len(score); i++ {
		if score[i] > max {
			max = score[i]
			winner = i
		}
	}
	fmt.Printf("Player %d won with score: %d\n", winner, max)
}
