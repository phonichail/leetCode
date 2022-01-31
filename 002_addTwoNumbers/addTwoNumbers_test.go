package addTwoNumbers

import (
	"testing"
)

func TestListNodeFactory(t *testing.T) {
	input := []int{5, 4, 3, 7, 4, 4, 2}
	expected := []int{2, 4, 4, 7, 3, 4, 5}
	result := ListNodeFactory(input)
	output := []int{}

	for node := result; node != nil; node = node.next {
		output = append(output, node.val)
	}

	for i := 0; i < len(output); i++ {
		if expected[i] != output[i] {
			t.Fatalf("\nExpected: %v\nGot:      %v\n", expected, output)
		}
	}

}

func testHelperAddTwoNumbers(expected *ListNode, got *ListNode) bool {
	y := got
	for x := expected; x != nil; x = x.next {
		if x.val != y.val {
			return false
		}
		y = y.next
	}
	return true
}

func TestOneAddTwoNumbers(t *testing.T) {
	l1 := ListNodeFactory([]int{2, 4, 3})
	l2 := ListNodeFactory([]int{5, 6, 4})
	expected := ListNodeFactory([]int{7, 0, 8})
	got := AddTwoNumbers(l1, l2)
	if !testHelperAddTwoNumbers(expected, got) {
		t.Fatalf("Wrong result")
	}
}

func TestTwoAddTwoNumbers(t *testing.T) {
	l1 := ListNodeFactory([]int{0})
	l2 := ListNodeFactory([]int{0})
	expected := ListNodeFactory([]int{0})
	got := AddTwoNumbers(l1, l2)
	testHelperAddTwoNumbers(expected, got)
	if !testHelperAddTwoNumbers(expected, got) {
		t.Fatalf("Wrong result")
	}
}

func TestThreeAddTwoNumbers(t *testing.T) {
	l1 := ListNodeFactory([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := ListNodeFactory([]int{9, 9, 9, 9})
	expected := ListNodeFactory([]int{8, 9, 9, 9, 0, 0, 0, 1})
	got := AddTwoNumbers(l1, l2)
	testHelperAddTwoNumbers(expected, got)
	if !testHelperAddTwoNumbers(expected, got) {
		t.Fatalf("Wrong result")
	}
}

func TestHelper(t *testing.T) {
	l1 := ListNodeFactory([]int{2, 4, 3})
	l2 := ListNodeFactory([]int{2, 4, 3})
	if !testHelperAddTwoNumbers(l1, l2) {
		t.Fatalf("\nFuck\n")
	}
}
