package helpers

import (
	"fmt"
	"sync"
	"time"
)

func PrintNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func PrintLetters(wg *sync.WaitGroup) {
	defer wg.Done()

	for c := 'a'; c < 'g'; c++ {
		fmt.Printf("Letter : %c\n", c)
		time.Sleep(150 * time.Millisecond)
	}

	fmt.Println()
}
