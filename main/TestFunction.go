package main

import (
	"errors"
	"fmt"
)

func main() {
	a, b := 10, 2
	c, err := div(a, b)
	fmt.Println(c, err)

	x := 100
	f := test(x)
	f()

	testDefer(1, 0)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func test(x int) func() {
	return func() {
		println(x)
	}
}

/**
defer 等同于finally
先执行end，再执行dispose
*/
func testDefer(a, b int) {
	defer println("dispose...")
	defer println("end...")
	println(a / b)
}
