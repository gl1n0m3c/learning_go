package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Goroutines2() {

	noConcurrent()
	fmt.Println("exit")

}

// func withWait() {
// 	var wg sync.WaitGroup
// 	wg.Add(6)

// 	for i := 0; i < 6; i++ {
// 		go func(i int) {
// 			fmt.Println(i + 1)
// 			wg.Done()
// 		}(i)
// 	}

// 	wg.Wait()
// }

// func concurrent() {
// 	start := time.Now()
// 	var counter int

// 	for i := 0; i < 1000; i++ {
// 		time.Sleep(time.Nanosecond)
// 		counter++
// 	}

// 	fmt.Println(counter)
// 	fmt.Println(time.Since(start))

// }

func noConcurrent() {
	start := time.Now()
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Since(start))
}
