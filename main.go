package main

import (
	datastructure "basics/data_structure"
	"fmt"
)

func main() {
	fmt.Println("Data Structure Main.go File")
	list := datastructure.LinkedList{}
	for i := 0; i < 10; i += 2 {
		list.InsertAtEnd(i)
	}

	for i := -5; i < 0; i += 1 {
		list.InsertAtHead(i)
	}

	list.Traverse()
	list.Delete(-3)
	list.Traverse()
	list.Reverse()
	list.Traverse()

}
