package variables

import "fmt"

func Variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Printf("O valor de b é %d e o valor de c é %d \n", b, c)

	var d int
	fmt.Println(d) //0

	e := "declarada e preenchida com uma string"
	fmt.Printf("A variável e foi %v \n", e)
}
