package main

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

	modifyInfo(a)
	println(a.x)

	modifyInfoByPoint(&a)
	println(a.x)
}

func modifyInfo(i info) {
	i.x = 1000
}

func modifyInfoByPoint(pi *info) {
	pi.x = 1000
}
