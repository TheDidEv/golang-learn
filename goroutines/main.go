package main

import (
	"fmt"
	// "math/rand"
	"sync"
	// "time"
)

type SafeCounter struct {
	mu     sync.Mutex
	NumMap map[string]int
}

func (s *SafeCounter) Add(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.NumMap["key"] = num
}

func main() {
	var wg sync.WaitGroup

	s := SafeCounter{NumMap: make((map[string]int))}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(i)
	}

	// CHANELs

	// c := make(chan int)
	// wg.Add(1)
	// go func() {
	// 	var sum int = 0
	// 	defer wg.Done()
	// 	for i := 0; i < 100; i++ {
	// 		fmt.Println("First func: ", i)
	// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	// 		sum += 1
	// 	}
	// 	c <- sum
	// }()

	// res := <-c

	// fmt.Println("Result: ", res)

	// GOROUTINE

	// wg.Add(1)
	// worker := func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 100; i++ {
	// 		fmt.Println("Second func: ", i)
	// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	// 	}
	// }
	// go worker()

	wg.Wait()
	fmt.Println(s.NumMap["key"])
	fmt.Println("Done")
}
