package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	var res int
	for  scanner.Scan() {
		item := scanner.Text()
		input = append(input, item)
	}

	res = calories(input)

	fmt.Println(res)
}

func calories(input []string) int {
	var max = 0
	var sum = 0

	for _, item := range input {
		cal, err := strconv.Atoi(item)
		if(err != nil) {
			sum = 0
		} else {
			sum += cal
			if(max < sum) {
				max = sum
			}
		}
	}

	return max
}
