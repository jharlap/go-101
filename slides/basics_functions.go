package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

func main() {
	fmt.Println("sum:", sum([]int{1, 2, 3}))
	// sum: 6

	fourtytwo, err := QueryNumber(42)
	fmt.Println("42:", fourtytwo, "err:", err)
	// 42: 42 is the number of US gallons in a barrel of oil. err: <nil>
}

// sum computes the sum of a list of numbers.
func sum(s []int) int {
	var r int
	for _, v := range s {
		r += v
	}
	return r
}

// QueryNumber gets an interesting fact about a number from NumbersAPI.
func QueryNumber(n int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := ctxhttp.Get(ctx, nil, fmt.Sprintf("http://numbersapi.com/%d", n))
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(r.Body)
	return string(b), err
}
