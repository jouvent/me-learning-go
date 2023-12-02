package day4

import (
	"fmt"
	"regexp"
	"strconv"

	"aoc2022/file"
)

func Process(input string, part int) {
	fmt.Println(process(input, part))
}

func process(input string, part int) (res int) {
	items := file.ReadLines(input)
	if part == 1 {
		res = countContained(items)
	} else {
		res = countOverlaps(items)
	}
	return
}

func countContained(input []string) int {
	res := 0
	for _, pair := range input {
		if isContained(pair) {
			res++
		}
	}
	return res
}

func isContained(pair string) bool {
	re := regexp.MustCompile("[0-9]+")
	areas := re.FindAllString(pair, -1)
	a, _ := strconv.Atoi(areas[0])
	b, _ := strconv.Atoi(areas[1])
	c, _ := strconv.Atoi(areas[2])
	d, _ := strconv.Atoi(areas[3])

	return (a <= c && b >= d) || (a >= c && b <= d)
}

func countOverlaps(input []string) int {
	res := 0
	for _, pair := range input {
		if overlaps(pair) {
			res++
		}
	}
	return res
}

func overlaps(pair string) bool {
	re := regexp.MustCompile("[0-9]+")
	areas := re.FindAllString(pair, -1)
	a, _ := strconv.Atoi(areas[0])
	b, _ := strconv.Atoi(areas[1])
	c, _ := strconv.Atoi(areas[2])
	d, _ := strconv.Atoi(areas[3])

	return (c <= a && a <= d) || (c <= b && b <= d) || (a <= c && c <= b) || (a <= d && d <= b)
}
