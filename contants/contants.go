package contants

import "fmt"

const s string = "constant"

func Constants() {
	fmt.Println(s)

	const n = 500000

	const d float64 = 3e20 / (n ^ 2)

	fmt.Printf("O valor de d é: %d \n", int64(d))

}
