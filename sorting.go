package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"cb", "ba", "ca"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.StringsAreSorted(strs)
	fmt.Println("Sorted:", s)

}