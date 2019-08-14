package main

func consumer(data chan int, done chan bool) {
	for x := range data {
		println("receive: ", x)
	}
	done <- true
}

func producer(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)
}

func main() {
	done := make(chan bool)
	date := make(chan int)

	go consumer(date, done)
	go producer(date)

	<-done
}
