package dayX

import "testing"

func ProcessTest(t *testing.T, input string, part, expect int) {
	res := process(input, part)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func Part1Test(t *testing.T, input []string, top, expect int) {
	res := part1(input, top)

	if res != expect {
		t.Errorf("from %v, got %d, wanted %d", input, res, expect)
	}
}

func Part2Test(t *testing.T, input []string, top, expect int) {
	res := part2(input, top)

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
	input := "../input_dayX.txt"
	expect := 0

	ProcessTest(t, input, 1, expect)
}

func TestOfficialInputs2(t *testing.T) {
	input := "../input_dayX.txt"
	expect := 0

	ProcessTest(t, input, 2, expect)
}
