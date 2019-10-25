package main 

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	var strs string = "test"
	p("Contains:", s.Contains(strs, "es"))
	p("Count: 	", s.Count(strs, "t"))
	p("HasPrefix:", s.HasPrefix(strs, "te"))
	p("HasSuffix:", s.HasSuffix(strs, "es"))
	p("Index:	", s.Index(strs, "e"))
	p("Join:	", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:	", s.Repeat("a", 5))
	p("Replace:	", s.Replace("foo", "o", "0", -1))
	p("Replace:	", s.Replace("foo", "o", "0", 1))
	p("Split:	", s.Split("a-b-c-d-e", "-"))
	p("ToLower:	", s.ToLower("TEST"))
	p("ToUpper:	", s.ToUpper("test"))
	p()

	p("Len:		", len("hello"))
	p("Char:	", "hello"[1])
}