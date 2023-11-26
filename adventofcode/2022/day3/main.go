package day3

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
