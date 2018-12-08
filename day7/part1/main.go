package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	jobs := make(map[string][]string, 0)
	jobsIDs := make([]string, 0)
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

	result := ""
	for i := 0; i < len(jobsIDs); i++ {
		v := jobsIDs[i]
		if len(jobs[v]) == 0 {
			result += v
			delete(jobs, v)
			for i2, j := range jobs {
				if at, ok := contains(v, j); ok {
					j = append(j[:at], j[at+1:]...)
					jobs[i2] = j
				}
			}
			jobsIDs = append(jobsIDs[:i], jobsIDs[i+1:]...)
			i = -1
		}
	}

	fmt.Println(result)
}

func contains(v string, a []string) (int, bool) {
	for i, arr := range a {
		if arr == v {
			return i, true
		}
	}
	return -1, false
}
