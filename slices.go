package main 

import "fmt"

func main() {
	
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("cap:", cap(s))

	c := make([]string, 3)
	copy(c, s)
	fmt.Println("cpy:", c)

	d := make([]string, 10)
	copy(d, s)
	fmt.Println("cpyd:", d)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	two := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		two[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			two[i][j] = i + j
		}
	}

	fmt.Println("2d", two)
}