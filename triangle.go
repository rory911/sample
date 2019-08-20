//
//  "Write a function that takes three sides of a triangle and answers
//  if it's equilateral, isosceles, or scalene."
//
//  See func Shape() that operates on Triangle objects
//
package sample

// Data structure for the triangle
type Triangle struct {
	Side1, Side2, Side3 uint
}

// Method to identify a triangle as Equalateral, Isosceles, Scalene, or Invalid
func (t Triangle) Shape() string {
	if t.Side1 == t.Side2 && t.Side2 == t.Side3 {
		return "Equalateral"
	}
	if t.Side1 >= t.Side2 && t.Side1 >= t.Side3 && t.Side3 + t.Side2 <= t.Side1 {
		return "Invalid"
	}
	if t.Side2 >= t.Side1 && t.Side2 >= t.Side3 && t.Side1 + t.Side3 <= t.Side2 {
		return "Invalid"
	}
	if t.Side3 >= t.Side1 && t.Side3 >= t.Side2 && t.Side1 + t.Side2 <= t.Side3 {
		return "Invalid"
	}
	if t.Side1 == t.Side2 || t.Side2 == t.Side3 || t.Side1 == t.Side3 {
		return "Isosceles"
	}
	return "Scalene"
}
