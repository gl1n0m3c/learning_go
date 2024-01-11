package structures

import "fmt"

type Square struct {
	Side int
}

func (square Square) show() {
	fmt.Println(square)
}

func (square *Square) increase(x int) {
	square.Side += x
}

func Structure() {
	square := Square{10}
	square.show()
	square.increase(10)
	square.show()
}

func Main() {
	fmt.Print()
}
