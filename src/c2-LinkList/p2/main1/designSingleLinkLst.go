package main

import (
	sl "c2-LinkList/LinkList/SingleList"
	 "fmt"
)

func main() {
	list := sl.NewSingleLinkList()
	
	for i := 1; i <= 10; i++ {
		node := sl.NewNode(i)		
		list.Append(node)
	}

	list.ShowLinkList()

	size := list.Size()

	fmt.Println("linkList size =",size)

	index := 9
	if v,err := list.Get(index);err != nil {
		panic(err)
	}else {
		fmt.Printf("%v of node in linkList value is %v\n",index,v)
	}

	// newNode := sl.Newnode(888)

	// list.InsertAtHead(newNode)

	// fmt.Println("after insert")
	// list.ShowLinkList()
	
	newNode2 := sl.NewNode(9999)
	err := list.InsertAtIndex(5,newNode2)

	if err != nil {
		fmt.Println(err)
		return 
	}

	list.ShowLinkList()
	
	list.RemoveIndexofNode(5)
	
	list.ShowLinkList()


	

}