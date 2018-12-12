package main

import "fmt"

const (
	playerCount         = 1234
	lastMarbleWorth     = 71852
	testPlayerCount     = 10
	testLastMarbleWorth = 1618
)

func main() {
	players := make(map[int]int)
	for i := 0; i < testPlayerCount; i++ {
		players[i] = 0
	}
	marbles := make([]int, 0)
	marbles = append(marbles, 0)
	currentMarbleLocation := 1
	currentPlayer := 0
	currentMarbelValue := 1
	for currentMarbelValue < testLastMarbleWorth {
		currentPlayer = (currentPlayer + 1) % (testPlayerCount + 1)
		if currentPlayer == 0 {
			currentPlayer = 1
		}

		if currentMarbelValue%23 == 0 {
			players[currentPlayer] += currentMarbelValue
			seventh := abs((currentMarbleLocation - 6) % len(marbles))
			// fmt.Println("Score of the seventh: ", marbles[seventh])
			players[currentPlayer] += marbles[seventh]
			currentMarbleLocation = abs((seventh - 1) % len(marbles))
			currentMarbelValue++
			marbles = append(marbles[:seventh], marbles[seventh+1:]...)
			// fmt.Println(marbles, currentMarbleLocation, currentMarbelValue)
			// fmt.Println("Marble value after removing: ", currentMarbleLocation)
			// fmt.Println("Marble value after removing: ", currentMarbelValue)
			continue
		}

		newIndex := abs((currentMarbleLocation + 2) % len(marbles))
		currentMarbleLocation = newIndex
		marbles = append(marbles[:newIndex+1], append([]int{currentMarbelValue}, marbles[newIndex+1:]...)...)

		currentMarbelValue++
		// fmt.Println(currentPlayer, marbles, currentMarbleLocation, currentMarbelValue)
	}

	max := 0
	winner := 0
	for k, v := range players {
		if v > max {
			max = v
			winner = k
		}
	}
	fmt.Printf("Winner is with score: %d and number: %d\n", max, winner)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
