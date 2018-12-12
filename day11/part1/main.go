package main

import "fmt"

const (
	gridSerialNumber = 9445
	testSerialNumber = 18
	minx             = 1
	miny             = 1
	maxx             = 300
	maxy             = 300
)

func main() {
	largestTotalPower := 0
	largestX := 0
	largestY := 0
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if x+3 > maxx || y+3 > maxy {
				continue
			}
			totalPower := getTotalPower(x, y)
			if totalPower > largestTotalPower {
				largestTotalPower = totalPower
				largestX = x
				largestY = y
			}
		}
	}
	fmt.Printf("Cell at x: %d and y: %d win with largest power level of: %d", largestX, largestY, largestTotalPower)
}

func getTotalPower(x, y int) (sum int) {
	for i := y; i < y+3; i++ {
		for j := x; j < x+3; j++ {
			rackID := j + 10
			powerLever := (rackID * i) + gridSerialNumber
			powerLever *= rackID
			hundredth := (powerLever % 1000) / 100
			hundredth -= 5
			sum += hundredth
		}
	}
	return sum
}
