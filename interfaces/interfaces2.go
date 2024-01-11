package interfaces

import (
	"fmt"
)

type Runner interface {
	Run()
	ShowStats()
}

type Sayer interface {
	Say()
}

type Human struct {
	Name, Surname string
	Age           int8
}

func (human Human) Run() {
	fmt.Println("Человек бежит.")
}

func (human Human) ShowStats() {
	fmt.Printf("Имя: %v\nФамилия: %v\nВозраст: %v\n", human.Name, human.Surname, human.Age)
}

func (human Human) Say() {
	fmt.Println("Человек говорит.")
}

func (human Human) Play() {
	fmt.Println("Человек играет.")
}

type Duck struct {
	Name string
	Age  int8
}

func (duck Duck) Run() {
	fmt.Println("Утка бежит.")
}

func (duck Duck) ShowStats() {
	fmt.Printf("Имя: %v\nВозраст: %v\n", duck.Name, duck.Age)
}

func (duck Duck) Say() {
	fmt.Println("Утка говорит.")
}

func Animals() {
	var runner Runner

	if runner == nil {
		fmt.Println("Runner is nil")
	}

	fmt.Printf("Type: %T Value: %v\n", runner, runner)

	var human *Human
	runner = human
	fmt.Printf("Type: %T Value: %v\n", runner, runner)

	human = &Human{"Ball", "Niggerson", 11}
	runner = *human
	fmt.Printf("Type: %T Value: %v\n", runner, runner)

	fmt.Printf("----------\n----------\n----------\n")

	typeChecking(runner)

	runner = Duck{"Niga", 5}

	typeChecking(runner)

}

func typeChecking(runner Runner) {
	fmt.Printf("Type: %T Value: %v\n", runner, runner)
	switch runner.(type) {
	case Human:
		fmt.Println("Человек.")
	case Duck:
		fmt.Println("Утка.")
	default:
		fmt.Println("Хз...")
	}

	// Или же можно сделать так

	if human, ok := runner.(Human); ok {
		fmt.Printf("Человек: %v\n", human)
	} else if duck, ok := runner.(Duck); ok {
		fmt.Printf("Утка: %v\n", duck)
	} else {
		fmt.Println("Хз...")
	}
}
