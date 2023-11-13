package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	var res int
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	res = countContained(input)

	fmt.Println(res)
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
