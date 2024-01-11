package contexts

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Contexts1() {
	ctx := context.Background()
	fmt.Println(ctx)

	toDo := context.TODO()
	fmt.Println(toDo)

	// withValue := context.WithValue(ctx, "name", "vasya")
	// fmt.Println(withValue)
	// fmt.Println(withValue.Value("name"))

	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println(withCancel.Err())
	cancel()
	fmt.Println(withCancel.Err())

	withTimout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	fmt.Println(<-withTimout.Done())
}

func Contexts2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	numsToProcess, numsProcessed := make(chan int, 5), make(chan int, 5)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numsToProcess, numsProcessed)
		}()
	}

	go func() {
		for i := 0; i <= 1000; i++ {
			numsToProcess <- i
			// Пример закрытия контекста
			// if i == 500 {
			// 	cancel()
			// }
		}
		close(numsToProcess)
	}()

	go func() {
		wg.Wait()
		close(numsProcessed)
	}()

	counter := 0
	for result := range numsProcessed {
		counter++
		fmt.Println(result)
	}
	fmt.Println(counter)
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}
