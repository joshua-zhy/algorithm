package main

import (
	"c2-LinkList/LinkList"
)

func removeElements(head *LinkList.ListNode, ele any) *LinkList.ListNode {//返回head
	tmp := head
	
	for tmp.Next != nil {
		if tmp.Next.Ele == ele {
			tmp.Next = tmp.Next.Next //删除元素的节点的前一个节点指向/删除元素的节点的下一个节点  gc 会回收删除的节点
		}else {
			tmp = tmp.Next//tmp移动到下一个
		}
	}

	return head
}

func main() {
	head := &LinkList.ListNode{}

	node1 := &LinkList.ListNode{1, nil}
	LinkList.InsertNode(head, node1)

	node2 := &LinkList.ListNode{2, nil}
	LinkList.InsertNode(head, node2)

	node3 := &LinkList.ListNode{3, nil}
	LinkList.InsertNode(head, node3)

	node4 := &LinkList.ListNode{4, nil}
	LinkList.InsertNode(head, node4)

	node5 := &LinkList.ListNode{5, nil}
	LinkList.InsertNode(head, node5)

	
	node6 := &LinkList.ListNode{2, nil}
	LinkList.InsertNode(head, node6)

	LinkList.ShowLinkList(head)
	// 1 ==> 2 ==> 3 ==> 4 ==> 5 ==> 2 ==> nil	删除2

	removeElements(head,2)

	LinkList.ShowLinkList(head)
}
