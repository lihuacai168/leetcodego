package medium

//给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。
//
//你应当保留两个分区中每个节点的初始相对位置。
//
//
//
//示例：
//
//输入：head = 1->4->3->2->5->2, x = 3
//输出：1->2->2->4->3->5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/partition-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	// 哑节点
	smallHead := small
	bigger := &ListNode{}
	biggerHead := bigger
	for head != nil {
		if head.Val > x {
			// 大节点指向当前的节点
			bigger.Next = head
			// 大节点的最新节点更新
			bigger = bigger.Next
		} else {
			small.Next = head
			small = small.Next
		}
		// 原始链表的节点向前推进一个
		head = head.Next
	}
	// 大节点链表的最后节点必须是指向空
	bigger.Next = nil
	// 小节点拼接大节点
	small.Next = biggerHead.Next
	return smallHead.Next
}
