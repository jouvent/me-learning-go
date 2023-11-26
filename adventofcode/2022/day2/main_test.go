package day2

import "testing"

func ProcessTest(t *testing.T, input string, part int, expect int) {
	res := process(input, part)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func PerformTest(t *testing.T, input []string, expect int) {
	res := rps2(input)

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

func TestOfficialInputs1(t *testing.T) {
	input := "../input_day2.txt"
	expect := 9177

	ProcessTest(t, input, 1, expect)
}

func TestOfficialInputs2(t *testing.T) {
	input := "../input_day2.txt"
	expect := 12111

	ProcessTest(t, input, 2, expect)
}
