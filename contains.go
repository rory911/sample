//
//  "Given two lists, write a function that answers if all elements of
//  one list are contained in the other list."
//
//  See Contains() that operates on heterogeneous lists
//
package sample

//  Compares lists looking for subset containment
//
func Contains(a []interface{}, b []interface{}) bool {
	if len(b) < len(a) {
		return false
	}

	prev, count := 0, 0
	for _,v := range a {
		prev = count
		for _,w := range b {
			if (v == w) {
				count++
				break
			}
		}
		if prev == count {
			return false
		}
	}

	return (len(a) == count)
}
