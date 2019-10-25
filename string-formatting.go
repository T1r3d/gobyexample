package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var p = fmt.Printf

func main() {
	v := point{x: 1, y: 2}
	p("%v\n", v)
	p("%+v\n", v)
	p("%#v\n", v)

	p("%p\n", &v)

	p("%x\n", 456)
	p("%x\n", "hex this 1234567890-=[];'.,")

	p("%s\n", "\"string\"")
	p("%q\n", "\"string\"")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}