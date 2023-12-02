package day1

import "testing"

func ProcessTest(t *testing.T, input string, part, expect int) {
	res := process(input, part)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func Part1Test(t *testing.T, input []string, expect int) {
	res := part1(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func Part2Test(t *testing.T, input []string, expect int) {
	res := part2(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestExample1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expect := 142

	Part1Test(t, input, expect)
}

func TestOfficialInputs1(t *testing.T) {
	input := "../input_day1.txt"
	expect := 54390

	ProcessTest(t, input, 1, expect)
}

func TestExample2Live1(t *testing.T) {
	input := "two1nine"
	expect := 29

	res := evaluateLivePart2(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestExample2Live2(t *testing.T) {
	input := "eightwothree"
	expect := 83

	res := evaluateLivePart2(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestExample2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expect := 281

	Part2Test(t, input, expect)
}

func TestOfficialInputs2(t *testing.T) {
	input := "../input_day1.txt"
	expect := 54277

	ProcessTest(t, input, 2, expect)
}
