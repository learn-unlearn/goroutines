package main

import (
	"fmt"
	// "time"
	// "sync"
)

// func main () {
	// This will keep going
	// count("fish")
	// count("goat")

	// This print both creating another rotine for fish because of the added go
	// go count("fish") // creates another routine
	// count("goat")

	// // Doesnt wait for both print since they work in another thread (sorry for using thread)
	// go count("fish") // creates another routine
	// go count("goat")
// }

// func main () {
// 	var wg sync.WaitGroup

// 	wg.Add(1) // create a routine
// 	go func () {
// 		count("sheep")
// 		wg.Done() // monitor routine till done
// 	}()

// 	wg.Wait() //waiting for routine to finish
// }

// func count (things string) {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i, things)
// 		time.Sleep(time.Millisecond * 500)
// 	}
	
// 	// for i := 1; true; i++ {
// 	// 	fmt.Println(i, things)
// 	// 	time.Sleep(time.Millisecond * 500)
// 	// }
// }

// func main () {
// 	c := make(chan string)
// 	go count("sheep", c)

// 	// for {
// 	for msg := range c {
// 		// can cause deadlock
// 		// msg := <- c // sending data received through the channel to msg variable
		
// 		// msg, open := <- c 
// 		// if !open {
// 		// 	break
// 		// }
// 		fmt.Println(msg) // printing the received data
// 	}
// }

// func count (things string, c chan string) { // c -> channel variable
// 	for i := 1; i <= 5; i++ {
// 		// fmt.Println(i, things)  // instead of printing we can send data to our main routine
// 		c <- things // sending the string in to the channel
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	close(c)
// }

// Playing with deadlock
// func main () {
// 	c := make(chan string, 2) // the number 2 be like creating a buffer 
// 	// c := make(chan string) // cause dealocks
// 	c <- "hello"
// 	c <- "world" // adding more channel than 2 cause a deadlock

// 	msg := <- c
// 	fmt.Println(msg)

// 	msg = <- c
// 	fmt.Println(msg)
// }

// Parralel
// func main () {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func () {
// 		for {
// 			c1 <- "Every 500ms"
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}()

// 	go func () {
// 		for {
// 			c2 <- "Every two seconds"
// 			time.Sleep(time.Second * 2)
// 		}
// 	}()

// 	for {
// 		// fmt.Println(<- c1)
// 		// fmt.Println(<- c2)

// 		// fmt.Println(<-c1) // thought it is syntax difference
// 		// fmt.Println(<-c2) // but no they work the same

// 		// using select not to delay the other worker from giving response
// 		// parrallelism in go [one disturbs not the other]
// 		select {
// 			case msg1 := <- c1:
// 				fmt.Println(msg1)
// 			case msg2 := <- c2:
// 				fmt.Println(msg2)
// 		}
// 	}
// }

// Playing with jobs

func main () {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)

	for i := 0;i < 100; i++ {
		jobs <- i
	}
	close(jobs) // can close here since it's decrementing func is here

	for j := 0; j <= 100; j++ {
		fmt.Println(<- results)
	}
}

// jobs <-chan int =====> receives data into jobs
// results chan<- int =======> send data into results
func worker (jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

// the return value on this fib func is necessary
func fib (n int) int {
	if n <= 1 {
		return n
	} 
	return fib(n - 1) + fib(n - 2)
}