package main

import "fmt"

func main() {
	var (
		x int  = 42
		y *int = &x
	)
	fmt.Println(x, y) // 42 0xc82000a2e8

	byValue(x)
	fmt.Println(x, y) // 42 0xc82000a2e8

	byReference(y)
	fmt.Println(x, y) // 9 0xc82000a2e8
}

func byValue(v int) {
	v = 6
}

func byReference(r *int) {
	*r = 9  // dereference to modify the value pointed at
	r = nil // modifies the local variable
}
