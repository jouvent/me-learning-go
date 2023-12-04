package day3

import "testing"

func ProcessTest(t *testing.T, input string, part, expect int) {
	res := process(input, part)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func Part1Test(t *testing.T, input string, expect int) {
	res := part1(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func Part2Test(t *testing.T, input string, expect int) {
	res := part2(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestExample1(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	expect := 4361

	Part1Test(t, input, expect)
}

func TestOfficialInputs1(t *testing.T) {
	input := "../input_day3.txt"
	expect := 527364

	ProcessTest(t, input, 1, expect)
}


func TestExample2(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	expect := 467835

	Part2Test(t, input, expect)
}

func TestOfficialInputs2(t *testing.T) {
	input := "../input_day3.txt"
	expect := 79026871

	ProcessTest(t, input, 2, expect)
}
