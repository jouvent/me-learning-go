package main

import "testing"

func PerformTest(t *testing.T, input []string, expect int) {
	res := rps(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestEmpty(t *testing.T) {
	input := []string{}
	expect := 0

	PerformTest(t, input, expect)
}

func TestExample1(t *testing.T) {
	input := []string{
		"A Y",
	}
	expect := 4

	PerformTest(t, input, expect)
}

func TestExample2(t *testing.T) {
	input := []string{
		"B X",
	}
	expect := 1

	PerformTest(t, input, expect)
}

func TestExample3(t *testing.T) {
	input := []string{
		"C Z",
	}
	expect := 7

	PerformTest(t, input, expect)
}

func TestExample(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}
	expect := 12

	PerformTest(t, input, expect)
}
