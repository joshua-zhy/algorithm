package Doublelinklist

import (
	"errors"
	"fmt"
	"log"
)

type node struct {
	val  int
	pre  *node
	next *node
}

type DoubleLinkList struct {
	Head *node
	size int
}

func NewDoubleLinkList() *DoubleLinkList {
	return &DoubleLinkList{
		Head: &node{}, //head virtual node
		size: 0,       //no include virtual node
	}
}

func NewNode(val int) *node {
	return &node{
		val:  val,
		next: nil,
	}
}

// append node to linkList end
func (dl *DoubleLinkList) Append(newNode *node) {
	tmp := dl.Head

	for tmp.next != nil { //find list end
		tmp = tmp.next //next node
	}

	//append list end
	// Two-way binding
	tmp.next = newNode
	newNode.pre = tmp
	dl.size++
}

// show linkList
func (dl *DoubleLinkList) ShowLinkList() {
	tmp := dl.Head

	if tmp.next == nil { //empty link list exit
		log.Fatalln("empty link list,please append node")
	}
	for tmp.next != nil {
		fmt.Printf("%v ==> ", tmp.next.val)
		tmp = tmp.next
	}
	fmt.Print("nil\n")
}

// get link list size
func (dl *DoubleLinkList) Size() int {
	return dl.size
}

func (dl *DoubleLinkList) Get(index int) (int, error) {
	if index <= 0 || index > dl.size {
		return -1, errors.New("index error")
	}

	tmp := dl.Head

	if tmp.next == nil { //empty link list exit
		log.Fatalln("empty link list,please append node")
	}

	for i := 1; i <= index; i++ {
		tmp = tmp.next
	}
	return tmp.val, nil
}

// remove node of index
func (dl *DoubleLinkList) RemoveIndexofNode(index int) error {
	if index <= 0 || index > dl.size {
		return errors.New("index error")
	}
	tmp := dl.Head

	for i := 1; i < index; i++ {
		tmp = tmp.next
	}
	//remove new node at previous position of tmp
	tmp.next = tmp.next.next //pointer to next node of rmeove node
	tmp.next.next.pre = tmp
	dl.size--
	return nil
}

// insert node at head
func (dl *DoubleLinkList) InsertAtHead(newNode *node) {
	tmp := dl.Head

	newNode.next = tmp.next //newnode -> linkLnk first node
	tmp.next.pre = newNode

	tmp.next = newNode //head ->  newNode
	newNode.pre = tmp
	dl.size++
}

// insert node at index position
func (dl *DoubleLinkList) InsertAtIndex(index int, newNode *node) error {
	if index <= 0 || index > dl.size {
		return errors.New("index error")
	}
	tmp := dl.Head

	for i := 1; i < index; i++ {
		tmp = tmp.next
	}
	//insert new node at previous position of tmp
	newNode.next = tmp.next

	tmp.next.pre = newNode

	tmp.next = newNode
	newNode.pre = tmp
	dl.size++
	return nil
}
