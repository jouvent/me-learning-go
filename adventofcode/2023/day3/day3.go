package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Process(input string, part int) {
	fmt.Println(process(input, part))
}

func process(input string, part int) (res int) {
	bytes, _ := os.ReadFile(input)
	if part == 1 {
		res = part1(string(bytes))
	} else {
		res = part2(string(bytes))
	}
	return
}

func part1(input string) int {
	dec := regexp.MustCompile(`\d+`)
	spe := regexp.MustCompile(`[^\d.]`)
	items := strings.Split(input, "\n")
	m := len(items[0])-1
	var res = 0
	for i, item := range items {
		matches := dec.FindAllString(item, -1)
		indexes := dec.FindAllStringIndex(item, -1)
		for j := range matches {
			s := max(0, indexes[j][0]-1)
			e := min(m, indexes[j][1]+1)
			area1 := items[max(0, i-1)][s:e]
			area2 := items[i][s:e]
			area3 := items[min(m, i+1)][s:e]
			if matchAny(spe, area1, area2, area3) {
				val, _ := strconv.Atoi(matches[j])
				res += val
			}
		}
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func matchAny(search *regexp.Regexp, input ...string) bool {
	for _, in := range input {
		if search.MatchString(in) {
			return true
		}
	}
	return false
}

func part2(input string) int {
	items := strings.Split(input, "\n")
	h := len(items)
	var res = 0
	dec := regexp.MustCompile(`\d+`)
	spe := regexp.MustCompile(`\*`)
	for row, item := range items {
		for _, s := range spe.FindAllStringIndex(item, -1) {
			numsInRange := []int{}
			for j := max(0, row-1); j <= min(h,row+1); j++ {
				matches := dec.FindAllString(items[j], -1)
				indexes := dec.FindAllStringIndex(items[j], -1)
				for k, n := range indexes {
					if n[0]-1 <= s[0] && s[0] <= n[1] {
						num, _ := strconv.Atoi(matches[k])
						numsInRange = append(numsInRange, num)
					}
				}
			}
			if len(numsInRange) == 2 {
				res += numsInRange[0] * numsInRange[1]
			}
		}
	}
	return res
}
