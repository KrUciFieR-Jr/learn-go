package main

import (
	"fmt"
	"time"
)

/** channels

Hold Data
Thread Safe
Listen Data
*/

/*
*
c <- 1
var i = <- c
This will give the deadlock error, because you are still reading in the unbuffered channel
Will throw deadlock so use along with go routine
*/
func main() {
	var c = make(chan int)
	go process(c)

	for i := range c {
		fmt.Println(i)
	}

	var cbuff = make(chan int, 5)
	fmt.Println("\n==========\n")
	go process(cbuff)
	for i := range cbuff {
		// Process can buffer upto 5 so it will concurrently do its deed decoupled from main func
		fmt.Println(i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i <= 4; i++ {
		c <- i
	}
	fmt.Println("Exiting Process")
}

/*
0
1
2
3
4
Exiting Process

==========

Exiting Process
0
1
2
3
4
*/

/*
The process:
go process(c) -> starts the channel
for i := range c -> wait for the input

defer close(c) -> close the channel after the function is executed to avoid deadlock
add i to the channel and it will print
*/
