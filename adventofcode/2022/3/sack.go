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
		input = append(input, scanner.Text())
	}

	res = sack(input)

	fmt.Println(res)
}

func sack(input []string) int {
	res := 0
	for _, sack := range input {
		item := inspectSack(sack)
		res += priority(item)
	}
	return res
}

func inspectSack(inventory string) byte {
	right := map[byte]bool{}
	left := map[byte]bool{}
	half := len(inventory) / 2
	for i := 0; i < half; i++ {
		left[inventory[i]] = true
		right[inventory[i+half]] = true
	}
	for k, _ := range left {
		if right[k] {
			return k
		}
	}
	return '*'
}

func priority(item byte) int {
	if item < 'a' {
		return int(27 + item - 'A')
	} else {
		return int(1 + item - 'a')
	}
}
