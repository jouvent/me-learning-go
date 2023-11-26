package day4

import "testing"

func ProcessTest(t *testing.T, input string, part int, expect int) {
	res := process(input, part)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func PerformTestOverlaps(t *testing.T, input string, expect bool) {
	res := overlaps(input)

	if res != expect {
		t.Errorf("from %v, got %t, wanted %t", input, res, expect)
	}
}

func PerformTestCountOverlaps(t *testing.T, input []string, expect int) {
	res := countOverlaps(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

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

func TestExampleLine1Part2(t *testing.T) {
	input := "2-4,6-8"
	expect := false

	PerformTestOverlaps(t, input, expect)
}

func TestExampleLine2Part2(t *testing.T) {
	input := "2-3,4-5"
	expect := false

	PerformTestOverlaps(t, input, expect)
}

func TestExampleLine3Part2(t *testing.T) {
	input := "5-7,7-9"
	expect := true

	PerformTestOverlaps(t, input, expect)
}

func TestExampleLine4Part2(t *testing.T) {
	input := "2-8,3-7"
	expect := true

	PerformTestOverlaps(t, input, expect)
}

func TestExampleLine5Part2(t *testing.T) {
	input := "6-6,4-6"
	expect := true

	PerformTestOverlaps(t, input, expect)
}

func TestExampleLine6Part2(t *testing.T) {
	input := "2-6,4-8"
	expect := true

	PerformTestOverlaps(t, input, expect)
}
func TestExamplePart2(t *testing.T) {
	input := []string{
		"2-4,6-8", // no
		"2-3,4-5", // no
		"5-7,7-9", // yes
		"2-8,3-7", // yes
		"6-6,4-6", // yes
		"2-6,4-8", // yes
	}
	expect := 4

	PerformTestCountOverlaps(t, input, expect)
}

func TestOfficialInputs1(t *testing.T) {
	input := "../input_day4.txt"
	expect := 511

	ProcessTest(t, input, 1, expect)
}

func TestOfficialInputs2(t *testing.T) {
	input := "../input_day4.txt"
	expect := 821

	ProcessTest(t, input, 2, expect)
}
