package dayX

import "testing"

func ProcessTest(t *testing.T, input string, part, expect int) {
	res := process(input, part)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func Part1Test(t *testing.T, input []string, top, expect int) {
	res := part1(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func Part2Test(t *testing.T, input []string, top, expect int) {
	res := part2(input)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func TestSimple(t *testing.T) {
	input := []string{"1", "", "1", "1"}
	expect := 2

	Part1Test(t, input, 1, expect)
}

func TestOfficialInputs1(t *testing.T) {
	input := "../input_dayX.txt"
	expect := 0

	ProcessTest(t, input, 1, expect)
}

func TestOfficialInputs2(t *testing.T) {
	input := "../input_dayX.txt"
	expect := 0

	ProcessTest(t, input, 2, expect)
}
