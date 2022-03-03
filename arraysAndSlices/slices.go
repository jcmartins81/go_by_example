package arraysandslices

import "fmt"

func Slices() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	fmt.Println(s[2:4])
	fmt.Println(s[2:])
	fmt.Println(s[:4])

}
