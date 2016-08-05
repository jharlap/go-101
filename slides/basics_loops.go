package main

import "fmt"

func main() {
	// for/range over a slice
	vs := []int{1, 2, 3}
	for i, v := range vs {
		fmt.Printf("Item %d has value %d\n", i, v)
	}

	// for/range over a map
	m := map[string]string{"question": "what is 6 times 9 🤔"}
	for k, v := range m {
		fmt.Printf("key: %s val: %s\n", k, v)
	}

	// for initialization; condition; post-statement
	for i := 5; i > 0; i-- {
		fmt.Printf("%d ", i)
	}

	// for condition
	i := 5
	for i > 0 {
		i--
	}
}
