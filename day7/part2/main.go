package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unicode"
)

const (
	// TaskFrequency is the base length of a task in seconds.
	TaskFrequency = 60
	// WorkerCount is the maximum available workers.
	WorkerCount = 5
)

type worker struct {
	job      string
	busy     bool
	workTime int
}

func (w *worker) work(v string, currentTime int) {
	delete(jobs, v)
	offset := int(unicode.ToLower([]rune(v)[0]) - 'a')
	w.workTime = currentTime + TaskFrequency + offset
	w.job = v
}

func (w *worker) finish() {
	w.workTime = -1
	result = append(result, w.job)
	for i2, j := range jobs {
		if at, ok := contains(w.job, j); ok {
			j = append(j[:at], j[at+1:]...)
			jobs[i2] = j
		}
	}
	w.job = ""
}

func (w *worker) isBusy(t int) bool {
	return t <= w.workTime
}

func (w *worker) canFinish(t int) bool {
	return t >= w.workTime && w.job != ""
}

var result []string
var jobs = make(map[string][]string, 0)
var jobsIDs = make([]string, 0)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	for _, l := range lines {
		var prev, next string
		fmt.Sscanf(l, "Step %s must be finished before step %s can begin.", &prev, &next)
		if _, ok := jobs[prev]; !ok {
			jobs[prev] = make([]string, 0)
		}

		jobs[next] = append(jobs[next], prev)
		if _, ok := contains(prev, jobsIDs); !ok {
			jobsIDs = append(jobsIDs, prev)
		}
		if _, ok := contains(next, jobsIDs); !ok {
			jobsIDs = append(jobsIDs, next)
		}
	}

	sort.Strings(jobsIDs)
	workers := make([]*worker, 0)
	for i := 0; i < WorkerCount; i++ {
		workers = append(workers, &worker{job: "", busy: false, workTime: -1})
	}

	cTime := 0
	for len(result) != len(jobsIDs) {
		jobsWorkersCanDo := make([]string, 0)

		for _, v := range jobsIDs {
			if _, ok := jobs[v]; ok {
				if len(jobs[v]) == 0 {
					jobsWorkersCanDo = append(jobsWorkersCanDo, v)
				}
			}
		}

		for _, w := range workers {
			if w.canFinish(cTime) {
				w.finish()
			}
		}

		for _, j := range jobsWorkersCanDo {
		INNER_LOOP:
			for _, w := range workers {
				if !w.isBusy(cTime) {
					w.work(j, cTime)
					break INNER_LOOP
				}
			}
		}
		cTime++
	}

	fmt.Println("current time: ", cTime)
	fmt.Println(strings.Join(result, ""))
}

func contains(v string, a []string) (int, bool) {
	for i, arr := range a {
		if arr == v {
			return i, true
		}
	}
	return -1, false
}
