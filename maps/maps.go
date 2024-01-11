package maps

import "fmt"

func Maps() {
	a := map[string]int{"f": 1, "s": 3, "g": 0}

	value1, ok1 := a["f"]
	value2, ok2 := a["s"]
	value3, ok3 := a["g"]
	value4, ok4 := a["r"]

	fmt.Println(value1, ok1)
	fmt.Println(value2, ok2)
	fmt.Println(value3, ok3)
	fmt.Println(value4, ok4)

	if !ok4 {
		a["r"] = 11
	}

	delete(a, "s")

	a["k"] = 7
	a["f"] = 20
	fmt.Println(a)
	for key, value := range a {
		fmt.Println(key, value)
	}

}
