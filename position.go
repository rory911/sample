//
//  "For a single-linked list (forward only), write a function that returns 5th element
//  from the end of the list. The list can only be walked once (reverse, length, or size
//  of this list cannot be used)." 
//
//  See PosFromEnd() that operates on LinkedList objects
//
package sample

import "fmt"

//  Individual element, or node, for a single-linked list
type Node struct {
	Data int
	Next *Node
}

//  Single-linked list object 
type LinkedList struct {
	Head *Node
}

//  Method to build the single-linked list
func (L *LinkedList) AppendList(value int) {
	Node := &Node{ Next: nil, Data: value, }

	if L.Head == nil {
		L.Head = Node
		return
	}

	List := L.Head
	for List.Next != nil {
		List = List.Next
	}
	List.Next = Node
}

//  Method to print the single-linked list
func (L *LinkedList) PrintList() {
	List := L.Head
	for List != nil {
		fmt.Printf(" %+v->", List.Data)
		List = List.Next
	}
	fmt.Println(" nil")
}

//  Nth node from the end of a single-linked list. The 0th element is nil.
func (L *LinkedList) PosFromEnd(pos uint) *Node {
	if L.Head == nil {
		return nil
	}

	List := L.Head
	Delta := L.Head

	for i := uint(0); i < pos; i++ {
		if List == nil {
			return nil
		}
		List = List.Next
	}


	for List != nil {
		List = List.Next
		Delta = Delta.Next
	}

	return Delta
}
