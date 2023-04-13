package main

import (
	sl "c2-LinkList/LinkList/SingleList"
	"fmt"
)

func reverseList(head *sl.Node) {
	cur := head.Next
	var pre *sl.Node

	for cur != nil {
		tmp := cur.Next //保存下一个node
		if tmp == nil { //最后一个node
			head.Next = cur //让虚拟头节点的下一个执行最后一个node

		}

		cur.Next = pre //让当前的node 指向前一个node
		pre = cur      //先移动pre
		cur = tmp      //在移动cur

	}
}

// func reverseListByRecursion(head *sl.Node)*sl.Node{
// 	return help(nil, head)
// }

// func help(pre, head *sl.Node)*sl.Node{
//     if head == nil {
//         return pre
//     }
//     next := head.Next
//     head.Next = pre
//     return help(head, next)
// }

func main() {
	list := sl.NewSingleLinkList()

	for i := 1; i <= 10; i++ {
		node := sl.NewNode(i)
		list.Append(node)
	}

	list.ShowLinkList()

	// reverseList(list.Head)

	list.Head = reverseListByRecursion(list.Head)

	fmt.Println("after reverse")

	list.ShowLinkList()

}
