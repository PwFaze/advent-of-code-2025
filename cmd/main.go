package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/PwFaze/advent-of-code-2025/day01"
	"github.com/PwFaze/advent-of-code-2025/day02"
)

type handler func(io.Reader) any

type aocResult struct {
	Result      string
	TimeElapsed time.Duration
}

type aocRunner struct {
	TaskName string
	Result   aocResult
	Handler  handler
	FileName string
	Day      int
	Part     int
}

var DAYRUNNER = []aocRunner{
	{
		TaskName: "Day 01, Part 01",
		FileName: "day01/day01.txt",
		Handler:  day01.Part01,
		Day:      1,
		Part:     1,
	},
	{
		TaskName: "Day 01, Part 02",
		FileName: "day01/day01.txt",
		Handler:  day01.Part02,
		Day:      1,
		Part:     2,
	},
	{
		TaskName: "Day 02, Part 01",
		FileName: "day02/day02.txt",
		Handler:  day02.Part01,
		Day:      2,
		Part:     1,
	},
	{
		TaskName: "Day 02, Part 02",
		FileName: "day02/day02.txt",
		Handler:  day02.Part02,
		Day:      2,
		Part:     2,
	},
}

func runPart(aocRunner aocRunner) aocResult {
	fmt.Printf("Running Task: %s \n", aocRunner.TaskName)
	f, err := os.Open(aocRunner.FileName)
	if err != nil {
		panic(fmt.Errorf("ReadFile in %s error: %s", aocRunner.TaskName, err))
	}
	defer f.Close()

	start := time.Now()
	r := aocRunner.Handler(f)
	duration := time.Since(start)

	res := aocResult{TimeElapsed: duration}

	switch v := r.(type) {
	case int:
		res.Result = strconv.Itoa(v)
	case int64:
		res.Result = strconv.FormatInt(v, 10)
	case uint64:
		res.Result = strconv.FormatUint(v, 10)
	case string:
		res.Result = v
	case fmt.Stringer:
		res.Result = v.String()
	default:
		res.Result = "unknown return value"
	}

	return res
}

func runDay(day int) {
	fmt.Printf("Running day %d\n", day)
	found := false
	for _, d := range DAYRUNNER {
		if d.Day == day {
			fmt.Printf("Day %d part %d\n", day, d.Part)
			r := runPart(d)
			fmt.Println("Result:", r.Result)
			fmt.Printf("Time elapsed: %s\n", r.TimeElapsed)
			found = true
		}
	}
	if !found {
		fmt.Printf("Did not find a solution for day %d\n", day)
	}
}

func main() {
	runDay(2)
}
