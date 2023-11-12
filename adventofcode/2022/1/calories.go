package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	items := []string{}
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	res := calories(items, 3)

	fmt.Println(res)
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
