package day1

import (
	"fmt"
	"strings"

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

func part1(items []string) (res int) {
	res = 0
	for _, item := range items {
		res += evaluateLivePart1(item)
	}
	return res
}

func evaluateLivePart1(item string) int {
	fi := strings.IndexAny(item, "0123456789")
	fv := rune(item[fi])
	li := strings.LastIndexAny(item, "0123456789")
	lv := rune(item[li])

	return 10*int(fv-'0') + int(lv-'0')
}

func part2(items []string) int {
	var res int
	for _, item := range items {
		res += evaluateLivePart2(item)
	}
	return res
}

func evaluateLivePart2(item string) int {
	fi := strings.IndexAny(item, "0123456789")
	var fv rune
	if fi != -1 {
		fv = rune(item[fi])
	} else {
		fi = len(item)
	}
	li := strings.LastIndexAny(item, "0123456789")
	var lv rune
	if li != -1 {
		lv = rune(item[li])
	}

	words := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	for w, v := range words {
		i := strings.Index(item, w)
		if i < fi && i != -1 {
			fi = i
			fv = v
		}
		i = strings.LastIndex(item, w)
		if i > li {
			li = i
			lv = v
		}
	}
	return 10*int(fv-'0') + int(lv-'0')
}
