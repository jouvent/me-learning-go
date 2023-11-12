package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	var res int
	for scanner.Scan() {
		item := scanner.Text()
		input = append(input, item)
	}

	res = rps(input)

	fmt.Println(res)
}

func rps(input []string) int {
	res := 0
	for _, match := range input {
		res += matchScore(match)
		res += myScore(match)
	}
	return res
}

func matchScore(match string) int {
	switch match {
	case "A Z", "B X", "C Y":
		return 0
	case "A X", "B Y", "C Z":
		return 3
	case "A Y", "B Z", "C X":
		return 6
	}
	return 0
}

func myScore(match string) int {
	switch match[2] {
	case 'X':
		return 1
	case 'Y':
		return 2
	case 'Z':
		return 3
	}
	return 0
}
