package concurrency

import "fmt"

func slice_to_channel(nums []int) <-chan int {

	out := make(chan int)
	go func() {

		for _, i := range nums {
			out <- i
		}
		close(out)
	}()

	return out
}

func sqaure(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {

		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out

}

func Square_Channel() {
	nums := []int{1, 2, 3, 4, 5, 6}

	dataChannel := slice_to_channel(nums)

	sqaures := sqaure(dataChannel)

	for i := range sqaures {
		fmt.Println(i)
	}

}
