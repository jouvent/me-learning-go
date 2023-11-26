package day1

import (
	"fmt"
	"sort"
	"strconv"

	"aoc2022/file"
)

func Process(input string, part int) {
	fmt.Println(process(input, part))
}

func process(input string, part int) (res int) {
	items := file.ReadLines(input)
	if part == 1 {
		res = calories(items, 1)
	} else {
		res = calories(items, 3)
	}
	return
}

func calories(items []string, top int) int {
	sums := []int{}
	var sum = 0

	for _, item := range items {
		cal, err := strconv.Atoi(item)
		if err != nil {
			sums = append(sums, sum)
			sum = 0
		} else {
			sum += cal
		}
	}
	sums = append(sums, sum)

	sort.Ints(sums)

	var res = 0
	for _, item := range sums[len(sums)-top:] {
		res += item
	}

	return res
}
