package concurrency

import "fmt"

func Testing_Channel() {
	my_channel := make(chan string)
	your_channel := make(chan string)

	go func() {
		my_channel <- "This data is coming from the channel"
	}()

	go func() {
		your_channel <- "This message is from another channel"
	}()

	select {
	case msgFromChannel := <-my_channel:
		fmt.Println("My Channel : ", msgFromChannel)
	case msgFromAnother := <-your_channel:
		fmt.Println("Your channel : ", msgFromAnother)
	}

	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		default:
			fmt.Println("default case")
		}

	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

}
