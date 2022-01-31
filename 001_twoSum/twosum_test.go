package twosum

import "testing"

func testTwoSumHelper(expected []int, got []int) bool {
	for i := 0; expected[i] == got[i] && i < len(expected); i++ {
		if i == len(expected)-1 {
			return true
		}
	}
	return false
}
func TestTwoSumOne(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}
	got := TwoSum(input, target)
	if !testTwoSumHelper(expected, got) {
		t.Fatalf("Expected %v, Got %v\n", expected, got)
	}
}

func TestTwoSumTwo(t *testing.T) {
	input := []int{3, 2, 4}
	target := 6

	expected := []int{1, 2}
	got := TwoSum(input, target)
	if !testTwoSumHelper(expected, got) {
		t.Fatalf("Expected %v, Got %v", expected, got)
	}
}

func TestTwoSumThree(t *testing.T) {
	input := []int{3, 3}
	target := 6

	expected := []int{0, 1}
	got := TwoSum(input, target)
	if !testTwoSumHelper(expected, got) {
		t.Fatalf("Expected %v, Got %v", expected, got)
	}
}
