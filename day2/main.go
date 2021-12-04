package main

import (
	"AdventOfCode2021/filereader"
	"fmt"
	"strconv"
	"strings"
)

var cmds []Command

func main() {
	lines := filereader.ReadFile("day2/input.txt", "\r\n")

	for i := 0; i < len(lines); i++ {
		str := strings.Split(lines[i], " ")
		parseInt, err := strconv.ParseInt(str[1], 10, 64)
		if err != nil {
			return
		}
		cmds = append(cmds, Command{direction: str[0], number: parseInt})
	}

	puzzle1()
	puzzle2()
}

func puzzle1() {
	s := Submarine{depth: 0, horizontal: 0, aim: 0}
	for i := 0; i < len(cmds); i++ {
		n := cmds[i].number
		switch cmds[i].direction {
		case "forward":
			s.horizontal = s.horizontal + n
		case "down":
			s.depth = s.depth + n
		case "up":
			s.depth = s.depth - n
		}
	}
	fmt.Println(s.horizontal * s.depth)
}

func puzzle2() {
	s := Submarine{depth: 0, horizontal: 0, aim: 0}
	for i := 0; i < len(cmds); i++ {
		n := cmds[i].number
		switch cmds[i].direction {
		case "forward":
			s.horizontal = s.horizontal + n
			s.depth = s.depth + s.aim*n
		case "down":
			s.aim = s.aim + n
		case "up":
			s.aim = s.aim - n
		}
	}
	fmt.Println(s.horizontal * s.depth)
}
