package main

const c string = "hello"
const (
	x = iota
	y
	z
)

func main() {
	println(c)
	println(x)
	println(y)
	println(z)
}
