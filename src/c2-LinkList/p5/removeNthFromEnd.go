package main

import (
	sl "c2-LinkList/LinkList/SingleList"
	_ "fmt"
)

func removeNthFromEnd(head *sl.Node, n int) {
	slow, fast := head, head

	// 先让快指针移动n步
	for n > 0 && fast != nil {
		fast = fast.Next
		n--
	}
	//让快指针再移动一步 为了让慢指针移动到删除节点的前一个位置
	fast = fast.Next

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//现在slow 的位置是删除节点的前一个位置

	slow.Next = slow.Next.Next //删除
}

func main() {
	list := sl.NewSingleLinkList()

	for i := 1; i <= 10; i++ {
		node := sl.NewNode(i)
		list.Append(node)
	}

	list.ShowLinkList()

	removeNthFromEnd(list.Head, 3)

	list.ShowLinkList()

}
