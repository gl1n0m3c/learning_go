package structures

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Worker struct {
	Person
	Name string
}

func (person Person) printName() {
	fmt.Println(person.Name)
}

func Structures2() {
	worker := Worker{Person{"Gay", 7}, "Not Gay"}
	fmt.Printf("Type: %T Value: %#v\n", worker, worker)

	fmt.Println(worker.Person.Age)
	fmt.Println(worker.Age)
	fmt.Println(worker.Name)

	worker.printName()
}
