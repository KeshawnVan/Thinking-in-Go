package main

import "fmt"

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}
func main() {
	var m manager
	m.name = "Tom"
	m.age = 11
	m.title = "CTO"
	fmt.Println(m)
	println(m.ToString())
}
