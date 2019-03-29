package pipeline

import (
	"fmt"
	"time"
)
/*
	The first stage, gen, is a function that converts a list of integers to a channel that emits the integers in the list.
	The gen function starts a goroutine that sends the integers on the channel
	and closes the channel when all the values have been sent
*/

func gen(numbers []int) chan int {
	ch := make(chan int, 5)
	go func () {
		for _, n := range numbers {
			ch <- n // blocks go routine for every insertion since its unbuffered
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	return ch;
}

/*
	The second stage, sq, receives integers from a channel and returns a channel that emits the square of each received integer.
	After the inbound channel is closed and this stage has sent all the values downstream, it closes the outbound channel
 */

func sq(numbers <-chan int) chan int {
	square := make(chan int, 5)
	go func () {

		for n := range numbers {
			//fmt.Println(n*n)
			square <-n*n
			time.Sleep(1 * time.Second)
		}
		close(square)
	}()

	return square;
}

func Execute() {

	numbers := []int{1, 2, 3, 4}

	w1 := sq(gen(numbers))

	for n := range w1 {
		fmt.Println(n)
	}

	// here trying to fetch data from the channel will return 0 as value and false as ok
	// value, ok := <-w1

	fmt.Println("Hello, playground")
}
