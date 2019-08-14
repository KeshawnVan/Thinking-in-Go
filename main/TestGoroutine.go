package main

import (
	"fmt"
	"time"
)

func task(id int) {
	for i := 0; i < 100; i++ {
		fmt.Printf("id: %d number: %d \n", id, i)
	}
}

func main() {
	go task(1)
	go task(2)
	time.Sleep(10 * time.Second)
}
