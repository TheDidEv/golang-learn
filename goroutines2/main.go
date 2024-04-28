package main

import (
	"fmt"
	"sync"
	"time"
)

// CHAPTER - 7
// Mutex/RWMutex

var (
	lock   sync.Mutex
	rwLock sync.RWMutex
	count  int
)

func main() {
	readAndWrite()
}

func basics() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Result: ", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}

func readAndWrite() {
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("done")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Printf("Read locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Read Unlock")
}

func write() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Printf("write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("write Unlock")
}

// CHAPTER - 6
// WAiting group

// func main() {
// 	var wg sync.WaitGroup
// 	names := []string{"FirstN", "asdwa", "jkjhs"}

// 	wg.Add(len(names))

// 	for _, name := range names {
// 		go somthingWithName(name, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("Done main!")
// }

// func somthingWithName(name string, wg *sync.WaitGroup) {
// 	fmt.Println(name)
// 	wg.Done()
// }

// CHAPTER - 5
// Channel Select Statement
// func main() {
// 	chan1, chan2 := make(chan string), make(chan string)

// 	go someFunk(chan1, "chan1")
// 	go someFunk(chan2, "chan2")

// 	select {
// 	case msg := <-chan1:
// 		fmt.Println(msg)
// 	case msg := <-chan2:
// 		fmt.Println(msg)
// 	}

// 	rouglyFair()
// }

// func someFunk(myChan chan string, message string) {
// 	myChan <- message
// }

// func rouglyFair() {
// 	first := make(chan interface{})
// 	close(first)
// 	second := make(chan interface{})
// 	close(second)

// 	var firstCounter, seocndCounter int
// 	for i := 0; i < 1000; i++ {
// 		select {
// 		case <-first:
// 			firstCounter++
// 		case <-second:
// 			seocndCounter++
// 		}
// 	}
// 	fmt.Printf("First: %d, second: %d", firstCounter, seocndCounter)
// }

// CHAPTER - 4
// Golang Channel Iteration & Channel Closing

// func main() {
// 	channel := make(chan string)
// 	// numRounds := 3

// 	// go throwingStar(channel, numRounds)
// 	// // fmt.Println(<-channel)
// 	// for i := 0; i < numRounds; i++ {
// 	// 	fmt.Println(<-channel)
// 	// }

// 	go throwingStar(channel)
// 	for {
// 		message, open := <-channel
// 		if !open {
// 			break
// 		} else {
// 			fmt.Println(message)
// 		}
// 	}
// }

// func throwingStar(channel chan string) {
// 	// func throwingStar(channel chan string, numRounds int) {
// 	rand.Seed(time.Now().UnixNano())
// 	numRounds := 3
// 	// // score := rand.Intn(10)
// 	// // channel <- fmt.Sprint("Your score: ", score)
// 	// for i := 0; i < numRounds; i++ {
// 	// 	score := rand.Intn(10)
// 	// 	channel <- fmt.Sprint("Your score: ", score)
// 	// }
// 	for i := 0; i < numRounds; i++ {
// 		score := rand.Intn(10)
// 		channel <- fmt.Sprint("Youre score: ", score)
// 	}
// 	close(channel)
// }

// CHAPTER - 3
// chan

// chan works from aprinciples FIFO
// func main() {
// 	channel := make(chan string, 2) // 2 - length or stack
// 	channel <- "First msg"
// 	channel <- "Second msg"

// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)
// 	// fmt.Println(<-channel) // will be error
// }

// CHAPTER - 2
// Try chan and go

// func main() {
// 	now := time.Now()
// 	defer func() {
// 		fmt.Println(time.Since(now))
// 	}()

// 	someSignal := make(chan bool)

// 	name := "SomeName"
// 	go attack(name, someSignal)
// 	someSignal <- false // when we change chan we had error - deadlock, fix on chapter - 3
// 	fmt.Println(<-someSignal)
// }

// func attack(name string, signal chan bool) {
// 	time.Sleep(time.Second)
// 	signal <- true
// 	fmt.Println("name: ", name)
// //	// fmt.Println(<-signal)
// }

// CHAPTER - 1
// Try defer

// func main() {
// 	start := time.Now()

// 	defer func() {
// 		fmt.Println(time.Since(start))
// 	}()

// 	Names := []string{"As", "asd", "ws"}

// 	for _, targ := range Names {
// 		go attack(targ)
// 	}
// 	time.Sleep(time.Second * 2)
// }

// func attack(target string) {
// 	fmt.Println(target)
// 	// fmt.Println(time.Second)
// }
