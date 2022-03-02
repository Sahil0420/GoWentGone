package datastructure

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) InsertAtEnd(data interface{}) {

	newNode := &Node{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (l *LinkedList) InsertAtHead(data interface{}) {
	newNode := &Node{Data: data, Next: l.Head}
	l.Head = newNode
}

func (l *LinkedList) Traverse() {
	temp := l.Head

	if temp == nil {
		fmt.Println("List is empty")
		return
	}

	for temp != nil {
		fmt.Print(temp.Data, "->")
		temp = temp.Next
	}
	fmt.Println("nil")
}

func (l *LinkedList) Delete(data interface{}) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == data {
		l.Head = l.Head.Next
		return
	}

	curr := l.Head
	for curr.Next != nil && curr.Next.Data != data {
		curr = curr.Next
	}

	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}

}

func (l *LinkedList) Reverse() {
	if l.Head == nil || l.Head.Next == nil {
		return
	}
	var prev *Node
	curr := l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

