package concurrency

import (
	"fmt"
	"time"
)

func DoWork(done <-chan string) {
	count := 0
	for {
		select {
		case msg := <-done:
			fmt.Println("Work is : ", msg)
			return
		default:
			fmt.Println("Default work : ", count)
			count++
		}
	}
}

func Testing_Done() {

	done := make(chan string)

	go DoWork(done)

	time.Sleep(time.Second * 1)
	close(done)

}
