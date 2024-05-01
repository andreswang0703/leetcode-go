package _141_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
		if s == f {
			return true
		}
	}
	return false
}
