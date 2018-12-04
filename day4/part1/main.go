package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type date struct {
	month  int
	day    int
	onDuty []int
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	guards := make(map[int][]int, 0)
	mostAsleep := make(map[int]int)
	var year, month, day, hour, minute, id int
	for _, l := range lines {
		if strings.Contains(l, "begins shift") {
			fmt.Sscanf(l, "[%d-%d-%d %d:%d] Guard #%d begins shift", &year, &month, &day, &hour, &minute, &id)
			if _, ok := guards[id]; !ok {
				guards[id] = make([]int, 59)
			}
		}
		if strings.Contains(l, "falls asleep") {
			fmt.Sscanf(l, "[%d-%d-%d %d:%d] falls asleep", &year, &month, &day, &hour, &minute)
		}
		if strings.Contains(l, "wakes up") {
			sleptFrom := minute
			fmt.Sscanf(l, "[%d-%d-%d %d:%d] wakes up", &year, &month, &day, &hour, &minute)
			diff := int(math.Abs(float64(sleptFrom) - float64(minute)))
			for i := 0; i < diff; i++ {
				index := (sleptFrom + i) % 60
				guards[id][index]++
			}
			mostAsleep[id] += diff
		}
	}

	max := 0
	maxID := 0
	for k, v := range mostAsleep {
		if v > max {
			max = v
			maxID = k
		}
	}

	mostDays := 0
	mostDaysIndex := 0
	for i, v := range guards[maxID] {
		if v > mostDays {
			mostDays = v
			mostDaysIndex = i
		}
	}

	fmt.Println(maxID * mostDaysIndex)
}
