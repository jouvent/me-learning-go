package day2

import (
	"fmt"
	"regexp"
	"strconv"

	"aoc2023/file"
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

	var res = 0
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, item := range items {
		r := regexp.MustCompile(`(\d+) (red|green|blue)`)
		matches := r.FindAllStringSubmatch(item, -1)
		var score = 0
		fmt.Sscanf(item, "Game %d:", &score)
		for i := range matches {
			c, _ := strconv.Atoi(matches[i][1])
			if bag[matches[i][2]] < c {
				score = 0
				break
			}
		}
		res += score
	}
	return res
}

func part2(items []string) int {
	var res = 0
	for _, item := range items {
		max := map[string]int{}
		r := regexp.MustCompile(`(\d+) (red|green|blue)`)
		matches := r.FindAllStringSubmatch(item, -1)
		for i := range matches {
			c, _ := strconv.Atoi(matches[i][1])
			if max[matches[i][2]] < c {
				max[matches[i][2]] = c
			}
		}
		res += max["red"] * max["green"] * max["blue"]
	}
	return res
}
