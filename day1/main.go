package main

import (
	"AdventOfCode2021/filereader"
	"fmt"
	"strconv"
)

func main() {
	challenge1()
	challenge2()
}

func challenge1() {
	lines := filereader.ReadFile("day1/input.txt", "\r\n")

	var sweepReport []int64

	for i := range lines {
		parseInt, err := strconv.ParseInt(lines[i], 10, 64)
		if err != nil {
			return
		}
		sweepReport = append(sweepReport, parseInt)
	}

	var counter int64 = 0
	for i := 0; i < len(sweepReport); i++ {
		if i+1 > len(sweepReport)-1 {
			break
		}

		prevDepth := sweepReport[i]
		nextDepth := sweepReport[i+1]

		if nextDepth > prevDepth {
			counter++
		}
	}
	fmt.Println(counter)
}

func challenge2() {
	lines := filereader.ReadFile("day1/input.txt", "\r\n")

	var sweepReport []int64

	for i := range lines {
		parseInt, err := strconv.ParseInt(lines[i], 10, 64)
		if err != nil {
			return
		}
		sweepReport = append(sweepReport, parseInt)
	}

	var advancedSweepReport []int64
	for i := 0; i < len(sweepReport); i++ {
		if i+2 > len(sweepReport)-1 {
			break
		}

		advancedSweepReport = append(advancedSweepReport, sweepReport[i]+sweepReport[i+1]+sweepReport[i+2])
	}

	var counter int64 = 0
	for i := 0; i < len(advancedSweepReport); i++ {
		if i+1 > len(advancedSweepReport)-1 {
			break
		}

		prevDepth := advancedSweepReport[i]
		nextDepth := advancedSweepReport[i+1]

		if nextDepth > prevDepth {
			counter++
		}
	}
	fmt.Println(counter)
}
