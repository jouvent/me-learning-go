package day3

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

func part1(input []string) int {
	res := 0
	for i := 0; i < len(input); i ++ {
		item := inspectSack(input[i])
		res += priority(item)
	}
	return res
}

func part2(input []string) int {
	res := 0
	for i := 0; i < len(input); i += 3 {
		item := findCommon(input[i], input[i+1], input[i+2])
		res += priority(item)
	}
	return res
}

func findCommon(first string, others ...string) byte {
	for _, v := range first {
		found := 0
		for j, other := range others {
			for _, z := range other {
				if v == z {
					found++
					break
				}
			}
			if found <= j {
				break
			}
		}
		if found == len(others) {
			return byte(v)
		}
	}
	return '*'
}

func inspectSack(inventory string) byte {
	half := len(inventory) / 2
	return findCommon(inventory[:half], inventory[half:])
}

func priority(item byte) int {
	if item < 'a' {
		return int(27 + item - 'A')
	} else {
		return int(1 + item - 'a')
	}
}
