package sample

import (
	"fmt"
	"testing"
)

func TestPosition(t *testing.T) {
	//  As per request, the 5th element
	testPosition(3, 5, t)	// 5th element from the end of a list with 3 elements
	testPosition(5, 5, t)	// 5th element from the end of a list with 5 elements
	testPosition(15, 5, t)	// 5th element from the end of a list with 15 elements

	//  Test other positions as well
	testPosition(10, 0, t)	// 0th element from the end of a list with 10 elements
	testPosition(20, 3, t)	// 3rd element from the end of a list with 20 elements
	testPosition(19, 7, t)	// 7th element from the end of a list with 19 elements
	testPosition(7, 2, t)	// 2nd element from the end of a list with 7 elements
	testPosition(11, 1, t)	// 1st element from the end of a list with 11 elements
	testPosition(7, 21, t)	// 21st element from the end of a list with 7 elements
	testPosition(25, 23, t)	// 23rd element from the end of a list with 25 elements
}

func testPosition(num uint, pos uint, t *testing.T) {
	lnklst := LinkedList{}
	ordinal := map[uint]string {0:"th", 1:"st", 2:"nd", 3:"rd", 4:"th",
				    5:"th", 6:"th", 7:"th", 8:"th", 9:"th"}

	// Simple linked list build
	for i := 0; uint(i) < num; i++ {
		lnklst.AppendList(i)
	}

	fmt.Println("For the following linked list:")
	lnklst.PrintList()
	fmt.Printf("The %d%s element from the end ", pos, ordinal[pos%10])
	if node := lnklst.PosFromEnd(pos); node == nil {
		fmt.Print("does not exist\n\n")
	} else {
		fmt.Print("contains(", node.Data,")\n\n")
		if int(num - pos) != node.Data {
			t.Errorf("Expected %d, but got %d", num - pos, node.Data)
		}
	}
}
