package main

import (
	"AdventOfCode2021/filereader"
	"fmt"
	"strconv"
	"strings"
)

var diagnosticReport []string

func main() {
	lines := filereader.ReadFile("day3/input.txt", "\r\n")
	for i := 0; i < len(lines); i++ {
		diagnosticReport = append(diagnosticReport, lines[i])
	}

	puzzle1()
	puzzle2()
}

func puzzle1() {
	var gamma string
	var epsilon string
	for n := 0; n < len(diagnosticReport[0]); n++ {
		var zero, one int
		for i := 0; i < len(diagnosticReport); i++ {
			if string(diagnosticReport[i][n]) == "0" {
				zero++
			} else {
				one++
			}
		}
		if zero > one {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
	}

	gammaConverted, _ := strconv.ParseInt(gamma, 2, 0)
	epsilonConverted, _ := strconv.ParseInt(epsilon, 2, 0)
	fmt.Println(gammaConverted * epsilonConverted)
}

func puzzle2() {
	oxygenGeneratorRating := lifeSupport("0", "1")
	co2ScrubberRating := lifeSupport("1", "0")
	fmt.Println(oxygenGeneratorRating * co2ScrubberRating)
}

func lifeSupport(zeroStr, oneStr string) int {
	lifesupport := diagnosticReport
	var prefix string
	for n := 0; n < len(lifesupport[0]); n++ {
		var zero, one int
		for i := 0; i < len(lifesupport); i++ {
			if string(lifesupport[i][n]) == "0" {
				zero++
			} else {
				one++
			}
		}

		if zero > one {
			prefix = prefix + zeroStr
		} else {
			prefix = prefix + oneStr
		}

		var lifesupportTmp []string
		for i := 0; i < len(lifesupport); i++ {
			if strings.HasPrefix(lifesupport[i], prefix) {
				lifesupportTmp = append(lifesupportTmp, lifesupport[i])
			}
		}

		if lifesupportTmp != nil {
			lifesupport = lifesupportTmp
		}

	}

	tmp, _ := strconv.ParseInt(lifesupport[0], 2, 0)

	return int(tmp)
}
