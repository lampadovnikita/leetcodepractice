// https://leetcode.com/problems/merge-two-sorted-lists/
// 21. Merge Two Sorted Lists
//
// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.
//
// Example 1:
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
//
// Example 2:
// Input: list1 = [], list2 = []
// Output: []
//
// Example 3:
// Input: list1 = [], list2 = [0]
// Output: [0]

package easy

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}

	resCurr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			resCurr.Next = list1

			list1 = list1.Next
		} else {
			resCurr.Next = list2

			list2 = list2.Next
		}

		resCurr = resCurr.Next
	}

	if list1 != nil {
		resCurr.Next = list1
	}

	if list2 != nil {
		resCurr.Next = list2
	}

	return dummy.Next
}
