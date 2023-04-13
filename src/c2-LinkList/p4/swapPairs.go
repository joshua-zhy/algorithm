package main

import (
	sl "c2-LinkList/LinkList/SingleList"
	_ "fmt"
)

//两两交换
// 1 ==> 2 ==> 3 ==> 4 ==> 5 ==> 6 ==> 7 ==> 8 ==> 9 ==> 10 ==>
// 2 ==> 1 ==> 4 ==> 3 ==> 6 ==> 5 ==> 8 ==> 7 ==> 10 ==> 9 ==>

func swapPairs(head *sl.Node) {
	cur := head

	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next           //1
		tmp2 := cur.Next.Next.Next //3 操作2的时候 2->3 会断开
		cur.Next = cur.Next.Next   //虚拟头节点 指向 第二个元素 0->2 这时候0->1 已经断了 需要保存一下1
		cur.Next.Next = tmp1       //2->1
		tmp1.Next = tmp2           //1->3

		cur = cur.Next.Next //移动2个位置
	}

}

func main() {
	list := sl.NewSingleLinkList()

	for i := 1; i <= 10; i++ {
		node := sl.NewNode(i)
		list.Append(node)
	}

	list.ShowLinkList()

	swapPairs(list.Head)

	list.ShowLinkList()

}
