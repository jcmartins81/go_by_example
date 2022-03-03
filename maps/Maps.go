package maps

import "fmt"

func Maps() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)
	fmt.Println(m["k1"])

	delete(m, "k1")
	fmt.Println("map: ", m)

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m2) //sempre ordena os valores pela key

}
