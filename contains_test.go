package sample

import (
	"fmt"
	"testing"
)

// Just Some Random data structure
type JSR struct {
	alpha int	// Just
	beta float32	// Something
	gamma string	// Random
}

func TestContains(t *testing.T) {
	a := []interface{} {1, 3.14, "hey", JSR{10,2.1,"yes"}, 2}
	b := []interface{} {1, -1, 3.14, JSR{10,2.1,"yes"}, 2, "hey", 2, JSR{10,2.1,"no"}}
	c := []interface{} {"hey", JSR{10,2.1,"no"}, 1, 2, 3.14}

	testContains(a, b, true, t)
	testContains(a, c, false, t)
	testContains(b, a, false, t)
	testContains(b, c, false, t)
	testContains(c, a, false, t)
	testContains(c, b, true, t)
}

func testContains(a []interface{}, b []interface{}, expected bool, t *testing.T) {
	if got := Contains(a, b); got == expected {
		fmt.Println(b, "contains", a, "is", got)
	} else {
		t.Error(b, "contains", a, "wanted", expected, "but got", got)
	}
}
