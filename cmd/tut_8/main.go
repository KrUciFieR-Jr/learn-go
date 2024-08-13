/*
Concurrency

    Definition: Concurrency refers to the ability of a system to handle multiple tasks at the same time by managing the execution of these tasks in an overlapping manner. It doesn't necessarily mean tasks are running simultaneously; they could be time-sliced or interleaved.
    Example: In a single-core processor, multiple threads can be managed by switching between them rapidly, giving the illusion of tasks running at the same time. For example, a web server handling multiple client requests by quickly switching between them.

Parallelism

    Definition: Parallelism involves actually performing multiple tasks at the same exact time by using multiple processing units (like CPU cores). It is a form of concurrency where tasks are literally executed in parallel on different cores or processors.
    Example: In a multi-core processor, parallelism occurs when different cores execute different threads or tasks at the same time, like processing multiple data streams in a parallel computing environment.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // When concurrent execution you need to wait
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var mux = sync.RWMutex{}
var results = []string{}

/*
What if you want to return from the function, if you written in Slice there would be memory
override by different function for that you use Mutex
*/

func main() {

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // Adds one to wait group,
		// then in function call we decrement by having the wg.Done method
		dbCall(i) // Non concurrent function call

	}
	wg.Wait() // Wait to grab the result

	fmt.Printf("\nTotal Execution for non concurrent time: %v", time.Since(t0))
	fmt.Printf("\nThe results are: %v", results)

	results = []string{}
	fmt.Println(results)
	t1 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i) // Concurrent function call

	}
	wg.Wait() // Wait to grab the result
	fmt.Printf("\nTotal Execution for concurrent time: %v", time.Since(t1))
	fmt.Printf("\nThe results are: %v", results)
}

func dbCall(i int) {
	// Simulating database call
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("\nThe result from the database is :", dbData[i])
	//mux.Lock()
	save(dbData[i])
	log()
	//mux.Unlock()
	wg.Done()
}

func save(result string) {
	mux.Lock()
	results = append(results, result)
	mux.Unlock()
}

func log() {
	/*
		RlockA
		lockA
		RlockB
		lockB

		RlockA will execute after lockA, RlockB, lockB
		Rlockb will execute after lockB

	*/
	mux.RLock()
	fmt.Println("\nThe results:", results)
	mux.RUnlock()
}

/*
Output :

The result from the database is : id1

The results: [id1]

The result from the database is : id2

The results: [id1 id2]

The result from the database is : id3

The results: [id1 id2 id3]

The result from the database is : id4

The results: [id1 id2 id3 id4]

The result from the database is : id5

The results: [id1 id2 id3 id4 id5]

Total Execution for non concurrent time: 4.749462083s
The results are: [id1 id2 id3 id4 id5][]

The result from the database is : id3

The results: [id3]

The result from the database is : id4

The results: [id3 id4]

The result from the database is : id2

The results: [id3 id4 id2]

The result from the database is : id5

The results: [id3 id4 id2 id5]

The result from the database is : id1

The results: [id3 id4 id2 id5 id1]

Total Execution for concurrent time: 1.426452334s
The results are: [id3 id4 id2 id5 id1]

*/
