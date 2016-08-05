package main

var (
	a int         // 0
	b int64       // 0
	c float64     // 0.0
	d complex64   // complex(0.0, 0.0) == 0+0i
	e rune        // ''
	f bool        // false
	g string      // ""
	h *int        // nil
	i []int       // nil
	j func()      // nil
	k map[int]int // nil
	l [3]int      // {0, 0, 0}
	m chan int    // nil
	n struct {    // struct{a: 0}
		a int
	}
)
