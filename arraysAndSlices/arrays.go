package arraysandslices

import "fmt"

func Arrays() {
	var a [5]int
	fmt.Println("Array vazio: ", a)

	a[4] = 100
	fmt.Println(a)

	fmt.Println("get a[4]: ", a[4])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
