package SingleLinkList

import (
	"errors"
	"fmt"
	"log"
)

//design singleList

type Node struct {
	Val  int
	Next *Node
}

type SingleLinkList struct {
	Head *Node
	size int
}

// constructor for list
func NewSingleLinkList() *SingleLinkList {
	return &SingleLinkList{
		Head: &Node{}, //head virtual Node
		size: 0,       //no include virtual Node
	}
}

// constructor for Node
func NewNode(val int) *Node {
	return &Node{
		Val:  val,
		Next: nil,
	}
}

// append Node to linkList end
func (sl *SingleLinkList) Append(newNode *Node) {
	tmp := sl.Head

	for tmp.Next != nil { //find list end
		tmp = tmp.Next //Next Node
	}

	//append list end
	tmp.Next = newNode
	sl.size++
}

// insert Node at head
func (sl *SingleLinkList) InsertAtHead(newNode *Node) {
	tmp := sl.Head

	newNode.Next = tmp.Next //newNode -> linkLnk first Node 

	tmp.Next = newNode //head ->  newNode
	sl.size++
}

//insert Node at index position
func (sl *SingleLinkList) InsertAtIndex(index int,newNode *Node) error {
	if index <= 0 || index > sl.size {
		return errors.New("index error")
	}
	tmp := sl.Head

	for i := 1; i < index ; i++ {
		tmp = tmp.Next
	}
	//insert new Node at previous position of tmp 
	newNode.Next = tmp.Next
	tmp.Next = newNode
	sl.size++
	return nil
} 

//remove Node of index
func (sl *SingleLinkList) RemoveIndexofNode(index int) error {
	if index <= 0 || index > sl.size {
		return errors.New("index error")
	}
	tmp := sl.Head

	for i := 1; i < index ; i++ {
		tmp = tmp.Next
	}
	//remove new Node at previous position of tmp 
	tmp.Next = tmp.Next.Next //pointer to Next Node of rmeove Node
	sl.size--
	return nil
}

// show linkList
func (sl *SingleLinkList) ShowLinkList() {
	tmp := sl.Head

	if tmp.Next == nil { //empty link list exit
		log.Fatalln("empty link list,please append Node")
	}
	for tmp.Next != nil {
		fmt.Printf("%v ==> ", tmp.Next.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

// get link list size
func (sl *SingleLinkList) Size() int {
	return sl.size
}

// get index Node val
func (sl *SingleLinkList) Get(index int) (int, error) {
	if index <= 0 || index > sl.size {
		return -1, errors.New("index error")
	}

	tmp := sl.Head

	if tmp.Next == nil { //empty link list exit
		log.Fatalln("empty link list,please append Node")
	}

	for i := 1; i <= index; i++ {
		tmp = tmp.Next
	}
	return tmp.Val, nil
}
