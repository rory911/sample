package sample

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	// Right => Scalene
	testTriangle(5, 4, 3, "Scalene", t)

	// Isosceles
	testTriangle(8, 10, 8, "Isosceles", t)

	// Equalateral
	testTriangle(8, 8, 8, "Equalateral", t)

	// Invalid 
	testTriangle(8, 8, 16, "Invalid", t)
	testTriangle(20, 8, 10, "Invalid", t)
	testTriangle(8, 18, 7, "Invalid", t)
}

func testTriangle(s1,s2,s3 uint, expected string, t *testing.T) {
	triangle := Triangle{Side1:s1, Side2:s2, Side3:s3}
	if got := triangle.Shape(); got == expected {
		fmt.Println("Triangle(", s1, ",", s2, ",", s3, ") is", got)
	} else {
		t.Error("For triangle(",s1,s2,s3,") got", got, "but expected", expected)
	}
}
