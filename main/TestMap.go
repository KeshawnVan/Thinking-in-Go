package main

func main() {
	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	println(x, ok)
	delete(m, "a")
}
