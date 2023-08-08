package main

import "testing"

func PerformTest(t *testing.T, input []string, expect int) {
	res := calories(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestSimple(t *testing.T) {
	input := []string{"1", "", "1", "1"}
	expect := 2

	PerformTest(t, input, expect)
}

func TestEmpty(t *testing.T) {
	input := []string{}
	expect := 0

	PerformTest(t, input, expect)
}

func TestWithZeroValues(t *testing.T) {
	input := []string{"1", "", "1", "0", "1"}
	expect := 2

	PerformTest(t, input, expect)
}
