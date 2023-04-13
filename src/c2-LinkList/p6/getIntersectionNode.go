package main

import (
	sl "c2-LinkList/LinkList/SingleList"
	"fmt"
)

// 4 ==> 1 ==> 8 ==> 4 ==> 5 ==>       
// 5 ==> 0 ==> 1 ==> 8 ==> 4 ==> 5 ==> 
// 求两个链表交点节点的指针  这里是 8

// c1 ，c2 表示链表的交集的长度
// a 表示l1除交集的长度
// b 表示l2除交集的长度

// 假设：ListA长度= a + c1，ListB长度= b + c2

// 核心逻辑：指针A遍历ListA + ListB，指针B遍历ListB + ListA

// 若ListA与ListB有相交=> c1与c2相等=>相遇点必为相交点=>（指针A）a + c1 + b =（指针B）b + c2 + a //遍历到相交点

// 若ListA与ListB无相交=> c1与c2不相等=>相遇点必为nil =>（指针A）a + c1 + b + c2 =（指针B）b + c2 + a + c1 相当于遍历两次L1.l2
func getIntersectionNode(headA, headB *sl.Node) *sl.Node {//传进来的是第一个元素
	l1,l2 := headA, headB

	for l1 != l2 {
        if l1 != nil {
            l1 = l1.Next
        } else {
            l1 = headB //l1结束了 到l2的第一个元素开始 直到有相交的点
        }

        if l2 != nil {
            l2 = l2.Next
        } else {
            l2 = headA //l2结束了 到l1的第一个元素开始 直到有相交的点
        }
    }
	return l1
}
func main() {
	list1 := sl.NewSingleLinkList()

	var point *sl.Node

	nodes1 := []int{4,1,8,4,5}
	for _,v := range nodes1 {
		if v == 8 {
			point = sl.NewNode(v)
			list1.Append(point)
			continue
		}
		node := sl.NewNode(v)
		list1.Append(node)

	}

	list1.ShowLinkList()
	// fmt.Println(point.Val)

	list2 := sl.NewSingleLinkList()

	nodes2 := []int{5,0,1}
	for _,v := range nodes2 {
		node := sl.NewNode(v)
		list2.Append(node)

	}
	list2.Append(point)
	list2.ShowLinkList()


	intersectionNode := getIntersectionNode(list1.Head.Next,list2.Head.Next)

	fmt.Println(intersectionNode.Val)



}
