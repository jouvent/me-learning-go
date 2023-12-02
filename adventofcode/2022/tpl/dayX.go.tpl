package dayX

import (
	"fmt"

	"aoc2022/file"
)

func Process(input string, part int) {
	fmt.Println(process(input, part))
}

func process(input string, part int) (res int) {
	items := file.ReadLines(input)
	if part == 1 {
		res = part1(items)
	} else {
		res = part2(items)
	}
	return
}

func part1(items []string) int {
	for _, item := range items {
	    fmt.Println(item)
	}
	var res = 0
	return res
}

func part2(items []string) int {
	for _, item := range items {
	    fmt.Println(item)
	}
	var res = 0
	return res
}
