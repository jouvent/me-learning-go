package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var max = 0
	var sum = 0
	for  scanner.Scan() {
		item, _ := strconv.Atoi(scanner.Text())
		if(item == 0) {
			sum = 0
		} else {
			sum += item
			if(max < sum) {
				max = sum
			}
		}

	}
	fmt.Println(max)
}

func calories() int {
	return 0
}
