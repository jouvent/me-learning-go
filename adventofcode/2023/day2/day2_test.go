package day2

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
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expect := 8

	Part1Test(t, input, expect)
}

func TestOfficialInputs1(t *testing.T) {
	input := "../input_day2.txt"
	expect := 2239

	ProcessTest(t, input, 1, expect)
}

func TestExample2(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expect := 2286

	Part2Test(t, input, expect)
}

func TestOfficialInputs2(t *testing.T) {
	input := "../input_day2.txt"
	expect := 83435

	ProcessTest(t, input, 2, expect)
}
