package channels

import (
	"fmt"
	"time"
)

func Channels1() {
	// С ним абсолютно нечего делать, он бесполезен
	var nilChannel chan int
	fmt.Printf("Channel: %v Len: %d Cap: %d\n", nilChannel, len(nilChannel), cap(nilChannel))

	// Небуферезированный канал (нельзя ничего записать, ибо он без буфера)
	noBufferChannel := make(chan int)
	fmt.Printf("Len: %d Cap: %d\n", len(noBufferChannel), cap(noBufferChannel))

	// Нельзя записать, если никто не читает
	// noBufferChannel <- 1

	// И наоборот
	// <- noBufferChannel

	// Рабочие примеры
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		chanForWriting <- 1
	}(noBufferChannel)

	value := <-noBufferChannel
	fmt.Println(value)

	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-chanForReading
		fmt.Println(value)
	}(noBufferChannel)

	noBufferChannel <- 2
}

func Channels2() {
	bufferChannel := make(chan int, 2)
	fmt.Printf("Len: %d Cap: %d\n", len(bufferChannel), cap(bufferChannel))

	// Логично, что если cap = 2, то можно максимум 2 числа засунуть в буфер, то же и с чтением из буфера
	// Как с небуфферезированным каналом при len = cap / len = 0 можно создать корутину, которая ждет

	bufferChannel <- 1
	bufferChannel <- 2
	fmt.Printf("Len: %d Cap: %d\n", len(bufferChannel), cap(bufferChannel))

	fmt.Println(<-bufferChannel)
	fmt.Println(<-bufferChannel)
	fmt.Printf("Len: %d Cap: %d\n", len(bufferChannel), cap(bufferChannel))
}

func Channels3() {
	arr := []int{4, 5, 6, 7, 8}
	channel := make(chan int, 1)

	go func(ch chan<- int) {
		for _, val := range arr {
			ch <- val
		}
		close(ch)
	}(channel)

	for value := range channel {
		fmt.Println(value)
	}

	// ИЛИ

	// for {
	// 	v, ok := <-channel
	// 	fmt.Println(v, ok)

	// 	if !ok {
	// 		break
	// 	}
	// }
}
