package ifelse

import "fmt"

func IfElse() {

	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é impar")
	}

	if num := -9; num < 0 {
		fmt.Println(num, " é negativo")
	} else if num < 10 {
		fmt.Println(num, "Tem apenas um digito")
	} else {
		fmt.Println(num, "tem multiplos digitos")
	}
}
