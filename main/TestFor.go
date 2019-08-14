package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	for i := 10; i > 0; i-- {
		println(i)
	}

	x := 0
	for x < 5 {
		println(x)
		x++
	}

	for {
		println(x)
		x--
		if x < 0 {
			break
		}
	}
	array := []int{100, 101, 102}
	for i, n := range array {
		println(i, ":", n)
	}
}
