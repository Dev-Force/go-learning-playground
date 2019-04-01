package faninfanout

import (
	"fmt"
	"sync"
	"time"
)
/*
	The first stage, gen, is a function that converts a list of integers to a channel that emits the integers in the list.
	The gen function starts a goroutine that sends the integers on the channel
	and closes the channel when all the values have been sent
*/

func gen(numbers []int, delay time.Duration) chan int {
	ch := make(chan int)
	go func () {
		for _, n := range numbers {
			// time.Sleep(delay)
			ch <- n // blocks go routine for every insertion since its unbuffered
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
			square <-n*n
			// time.Sleep(1000 * time.Millisecond)
		}
		close(square)
	}()

	return square;
}

func merge(done chan struct{}, channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(channels))

	for _, c := range channels {
		go func(c <-chan int) {
			for n := range c {
				select {
				case out <- n:
				case <-done:
				}
			}
			wg.Done()
		}(c)
	}

	go func(wgc *sync.WaitGroup, c chan int) {
		wg.Wait()
		close(out)
	}(&wg, out)

	return out;
}

func Execute() {

	done := make(chan struct{}, 2)
	numbers := []int{2, 3}

	genNumbers := gen(numbers, 200 * time.Millisecond)


	sqdWorker1 := sq(genNumbers)

	sqdWorker2 := sq(genNumbers)

	out := merge(done, sqdWorker1, sqdWorker2)

	for n := range out {
		fmt.Println(n)
	}

	// here trying to fetch data from the channel will return 0 as value and false as isOpen

}
