package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a, b, c, d, e := 1, 2, 3, 4, 5
	fmt.Printf("Address from the main function:\na = %d b = %d c = %d d = %d e = %d\n", &a, &b, &c, &d, &e)
	// Type casting based on int types: a, b, c, d, e := int32(1), int32(2), int32(3), int32(4), int32(5)

	// pa is an variable that is an array of 5 integer pointers
	var pa [5]*int = [5]*int{&a, &b, &c, &d, &e}
	fmt.Printf("Size of pa (as an array of 5 integer pointers) = %d\n", unsafe.Sizeof(pa))
	for i := 0; i < len(pa); i++ {
		fmt.Printf("element[%d] points to %d with default address = %d at present address = %d\n",
			i, *(pa[i]), pa[i], uintptr(unsafe.Pointer(&pa[i])))
	}
}
