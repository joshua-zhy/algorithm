package main

import (
	dl "c2-LinkList/LinkList/DoubleLinkList"
	"fmt"
)

func main() {
	list := dl.NewDoubleLinkList()

	for i := 1; i <= 10; i++ {
		node := dl.NewNode(i)		
		list.Append(node)
	}

	list.ShowLinkList()

	fmt.Println(list.Size())

	if val,err := list.Get(5); err != nil {
		panic(err)
	}else {
		fmt.Println(val)
	}

	
	// newNode := dl.NewNode(888)

	// list.InsertAtHead(newNode)

	// fmt.Println("after insert")
	// list.ShowLinkList()

	newNode2 := dl.NewNode(9999)
	err := list.InsertAtIndex(5,newNode2)

	if err != nil {
		fmt.Println(err)
		return 
	}

	list.ShowLinkList()
	
	list.RemoveIndexofNode(5)
	
	list.ShowLinkList()
}


	
