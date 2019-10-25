package main

import (
	"fmt"
	"os"
)

func main() {
	argswithprog := os.Args
	argswitoutprog := os.Args[1:]

	args := os.Args[3]

	fmt.Println(argswithprog)
	fmt.Println(argswitoutprog)
	fmt.Println(args)
	
	
}