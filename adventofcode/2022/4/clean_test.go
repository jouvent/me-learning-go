package main

import "testing"

func PerformTestIsContained(t *testing.T, input string, expect bool) {
	res := isContained(input)

	if res != expect {
		t.Errorf("from %v, got %t, wanted %t", input, res, expect)
	}
}

func PerformTestCountContained(t *testing.T, input []string, expect int) {
	res := countContained(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestExampleLine1(t *testing.T) {
	input := "2-4,6-8"
	expect := false

	PerformTestIsContained(t, input, expect)
}

func TestExampleLine2(t *testing.T) {
	input := "2-3,4-5"
	expect := false

	PerformTestIsContained(t, input, expect)
}

func TestExampleLine3(t *testing.T) {
	input := "5-7,7-9"
	expect := false

	PerformTestIsContained(t, input, expect)
}

func TestExampleLine4(t *testing.T) {
	input := "2-8,3-7"
	expect := true

	PerformTestIsContained(t, input, expect)
}

func TestExampleLine5(t *testing.T) {
	input := "6-6,4-6"
	expect := true

	PerformTestIsContained(t, input, expect)
}

func TestExampleLine6(t *testing.T) {
	input := "2-6,4-8"
	expect := false

	PerformTestIsContained(t, input, expect)
}

func TestDebug(t *testing.T) {
	input := "4-11,12-12"
	expect := false

	PerformTestIsContained(t, input, expect)
}

func TestExample(t *testing.T) {
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	expect := 2

	PerformTestCountContained(t, input, expect)
}
