package day1

import "testing"

func ProcessTest(t *testing.T, input string, part, expect int) {
	res := process(input, part)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

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

func TestOfficialInputs1(t *testing.T) {
	input := "../input_day1.txt"
	expect := 69289

	ProcessTest(t, input, 1, expect)
}

func TestOfficialInputs2(t *testing.T) {
	input := "../input_day1.txt"
	expect := 205615

	ProcessTest(t, input, 2, expect)
}
