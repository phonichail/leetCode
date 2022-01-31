package addTwoNumbers

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	val  int
	next *ListNode
}

func ListNodeFactory(nums []int) *ListNode {
	newListNode := ListNode{val: nums[len(nums)-1]}
	if 1 < len(nums) {
		newListNode.next = ListNodeFactory(nums[:len(nums)-1])
	}
	return &newListNode
}

func getNumsFromListOfNodes(node *ListNode) int {
	textNum := ""
	for x := node; x != nil; x = x.next {
		textNum = textNum + strconv.Itoa(x.val)
	}
	num, _ := strconv.Atoi(textNum)
	return num
}

func getSliceFromNum(num int) []int {
	textNum := strconv.Itoa(num)
	slice := []int{}

	for i := 0; i < len(textNum); {
		x, _ := strconv.Atoi(textNum[len(textNum)-1:])
		textNum = textNum[:len(textNum)-1]
		slice = append(slice, x)
	}
	fmt.Println(slice)
	return slice
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	x := getNumsFromListOfNodes(l1)
	y := getNumsFromListOfNodes(l2)
	return ListNodeFactory(getSliceFromNum(x + y))
}
