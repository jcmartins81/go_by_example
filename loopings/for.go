package loopings

import "fmt"

func Loopings() {

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for {

		if i > 5 {
			break //para o loop quando i = 6
		}

		fmt.Println("loop")
		i++

	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue //Vai para o próximo loop
		}
		fmt.Println(n)
	}
}
