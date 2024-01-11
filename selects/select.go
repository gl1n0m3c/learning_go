package selects

import "fmt"

func Select() {
	bufferedChan := make(chan string, 2)
	bufferedChan <- "first"

	select {
	case str := <-bufferedChan:
		fmt.Println(str)
	case bufferedChan <- "second":
		fmt.Println(<-bufferedChan, <-bufferedChan)
	}
}
