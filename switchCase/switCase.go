package switchcase

import (
	"fmt"
	"time"
)

func Switchcase() {

	i := 3

	fmt.Printf("Write %d", i)
	switch i {
	case 1:
		fmt.Println(" as One")
	case 2:
		fmt.Println(" as two")
	case 3:
		fmt.Println(" as three")
	}

	thisTime := time.Now()

	switch {
	case thisTime.Hour() < 12:
		fmt.Println("Ainda é de manhã")
	default:
		fmt.Println("Já é tarde!")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Final de semana!")
	default:
		fmt.Println("Hoje é ", time.Now().Weekday())
	}

	oQueEuSou := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Eu sou booleano")
		case int:
			fmt.Println("Eu sou um inteiro")
		case string:
			fmt.Println("Eu sou uma string")
		default:
			fmt.Printf("Não conheço o tipo %T \n", t)
		}
	}

	oQueEuSou("Oi")
	oQueEuSou(true)
	oQueEuSou(13)
	oQueEuSou(3.14)
}
