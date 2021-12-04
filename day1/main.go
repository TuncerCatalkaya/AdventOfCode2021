package main

import (
	"AdventOfCode2021/filereader"
	"fmt"
	"strconv"
)

var sweepReport []int64

func main() {
	lines := filereader.ReadFile("day1/input.txt", "\r\n")

	for i := range lines {
		parseInt, err := strconv.ParseInt(lines[i], 10, 64)
		if err != nil {
			return
		}
		sweepReport = append(sweepReport, parseInt)
	}

	puzzle1()
	puzzle2()
}

func puzzle1() {
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

func puzzle2() {
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
