/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  func mergeKLists(lists []*ListNode) *ListNode {
// 	// filter nil node
// 	heap := make([]*ListNode, 0, len(lists))
// 	for _, list := range lists {
// 		if list != nil {
// 			heap = append(heap, list)
// 		}
// 	}
// 	// build min-heap
// 	n := len(heap)
// 	if n == 0 {
// 		return nil
// 	}
// 	for i := n / 2; i >= 0; i-- {
// 		adjust(heap, i, n-1)
// 	}
// 	head := &ListNode{0, nil}
// 	temp := head
// 	for heap[0].Val != math.MaxInt32 {
// 		temp.Next = heap[0]
// 		temp = temp.Next
// 		heap[0] = heap[0].Next
// 		if heap[0] == nil {
// 			heap[0] = &ListNode{math.MaxInt32, nil}
// 		}
// 		adjust(heap, 0, n-1)
// 	}
// 	return head.Next
// }

// func adjust(heap []*ListNode, lo, hi int) {
// 	i, j := lo, lo*2+1
// 	for j <= hi {
// 		if j+1 <= hi && heap[j+1].Val < heap[j].Val {
// 			j++
// 		}
// 		if heap[j].Val < heap[i].Val {
// 			heap[i], heap[j] = heap[j], heap[i]
// 			i = j
// 			j = i*2 + 1
// 		} else {
// 			break
// 		}
// 	}
// }
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	return mergeTwoLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head = &ListNode{0, nil}
	temp := head
	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			temp.Next = l2
			temp = temp.Next
			l2 = l2.Next
		} else {
			temp.Next = l1
			temp = temp.Next
			l1 = l1.Next
		}
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return head.Next
}

