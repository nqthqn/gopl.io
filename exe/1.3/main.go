package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	timeTo(echo1)
	timeTo(echo2)
	timeTo(echo3)
	timeTo(echo1a)
}

func timeTo(f func()) {
	t := time.Now()
	f()
	fmt.Println(time.Since(t))
}

// 9586 ns/op	112 B/op	3 allocs/op
func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		// Strings are immutable ∴ this is littering
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo1a() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s ", os.Args[i])
	}
	fmt.Println()
}

// 3555 ns/op	112 B/op	3 allocs/op
func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// 2895 ns/op	80 B/op		2 allocs/op
func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

/*
op:		single iteration
ns/op:		how long each op took in nano seconds
allocs/op:	how many distinct memory allocations occurred per op
B/op:		how many bytes were allocated per op
*/
