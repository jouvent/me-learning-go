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

func myScore(match string) int {
	switch match {
	case
		"A Y", // draw to rock: rock
		"B X", // loose to paper: rock
		"C Z": // win to scisor: rock
		return 1
	case
		"A Z", // win to rock: paper
		"B Y", // draw to paper: paper
		"C X": // loose to scisor: paper
		return 2
	case
		"A X", // loose to rock: scisor
		"B Z", // win to paper: scisor
		"C Y": // draw to scisor: scisor
		return 3
	}
	return 0
}

func matchScore(match string) int {
	switch match[2] {
	case 'X':
		return 0
	case 'Y':
		return 3
	case 'Z':
		return 6
	}
	return 0
}
