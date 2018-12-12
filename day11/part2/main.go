package main

import "fmt"

const (
	gridSerialNumber = 9445
	testSerialNumber = 18
	maxx             = 300
	maxy             = 300
)

func main() {
	largestTotalPower := 0
	largestX := 0
	largestY := 0
	largestSize := 0
	for y := 1; y <= maxy; y++ {
		for x := 1; x <= maxx; x++ {
			for s := 1; s <= 300; s++ {
				totalPower := getTotalPower(s, x, y)
				if totalPower >= largestTotalPower {
					largestTotalPower = totalPower
					largestX = x
					largestY = y
					largestSize = s
				}

			}
		}
	}

	fmt.Printf(
		"Cell at x: %d and y: %d win with largest power level of: %d and size: %d\n",
		largestX,
		largestY,
		largestTotalPower,
		largestSize)
}

func getTotalPower(square, x, y int) (sum int) {
	for i := y; i <= y+square; i++ {
		for j := x; j <= x+square; j++ {
			rackID := j + 10
			powerLever := ((rackID * i) + testSerialNumber) * rackID
			hundredth := ((powerLever % 1000) / 100) - 5
			sum += hundredth
		}
	}
	return sum
}
