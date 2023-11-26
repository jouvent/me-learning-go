package day3

import "testing"

func PerformTestInspectSack(t *testing.T, input string, expect byte) {
	res := inspectSack(input)

	if res != expect {
		t.Errorf("from %v, got %c, wanted %c", input, res, expect)
	}
}

func PerformTestSack(t *testing.T, input []string, expect int) {
	res := sack(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestEmpty(t *testing.T) {
	input := []string{}
	expect := 0

	PerformTestSack(t, input, expect)
}

func TestExampleLine1(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp"
	var expect byte = 'p'

	PerformTestInspectSack(t, input, expect)
}

func TestExampleLine2(t *testing.T) {
	input := "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
	var expect byte = 'L'

	PerformTestInspectSack(t, input, expect)
}

func TestExampleLine3(t *testing.T) {
	input := "PmmdzqPrVvPwwTWBwg"
	var expect byte = 'P'

	PerformTestInspectSack(t, input, expect)
}

func TestExampleLine4(t *testing.T) {
	input := "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"
	var expect byte = 'v'

	PerformTestInspectSack(t, input, expect)
}

func TestExampleLine5(t *testing.T) {
	input := "ttgJtRGJQctTZtZT"
	var expect byte = 't'

	PerformTestInspectSack(t, input, expect)
}

func TestExampleLine6(t *testing.T) {
	input := "CrZsJsPPZsGzwwsLwLmpwMDw"
	var expect byte = 's'

	PerformTestInspectSack(t, input, expect)
}

func TestExample(t *testing.T) {
	input := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	expect := 70

	PerformTestSack(t, input, expect)
}
