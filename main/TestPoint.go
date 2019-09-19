package main

import "fmt"

type info struct {
	x int
}

func main() {
	x := 10

	var p = &x

	*p += 20

	println(p, *p)

	a := info{}

	a.x = 100

	pa := &a
	pa.x += 100
	println(pa.x)

	fmt.Printf("a address is %p \n", &a)
	modifyInfo(a)
	println(a.x)

	modifyInfoByPoint(&a)
	println(a.x)

	i := 100
	fmt.Printf("i address is %d \n", &i)
	modifyInt(i)
	fmt.Printf("i address is %d \n", &i)
	modifyIntByPoint(&i)
	fmt.Printf("i address is %d \n", &i)
	fmt.Printf("i is %d", i)
}

func modifyInfo(i info) {
	fmt.Printf("a address is %p \n", &i)
	i.x = 1000
}

func modifyInfoByPoint(pi *info) {
	fmt.Printf("a address is %p \n", pi)
	pi.x = 1000
}

func modifyInt(i int) {
	fmt.Printf("i address is %d \n", &i)
	i += 100
}

func modifyIntByPoint(pi *int) {
	fmt.Printf("i address is %d \n", pi)
	*pi += 100
}
