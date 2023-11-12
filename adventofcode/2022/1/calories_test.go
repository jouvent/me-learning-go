package main

import "testing"

func PerformTest(t *testing.T, input []string, top, expect int) {
	res := calories(input, top)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestSimple(t *testing.T) {
	input := []string{"1", "", "1", "1"}
	expect := 2

	PerformTest(t, input, 1, expect)
}

func TestEmpty(t *testing.T) {
	input := []string{}
	expect := 0

	PerformTest(t, input, 1, expect)
}

func TestWithZeroValues(t *testing.T) {
	input := []string{"1", "", "1", "0", "1"}
	expect := 2

	PerformTest(t, input, 1, expect)
}

func TestTop3(t *testing.T) {
	input := []string{"3", "", "1", "", "1", "", "5", "", "4"}
	expect := 12

	PerformTest(t, input, 3, expect)
}
