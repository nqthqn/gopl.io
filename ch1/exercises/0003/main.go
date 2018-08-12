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
}

func timeTo(f func()) {
	t := time.Now()
	f()
	fmt.Println(time.Since(t))
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // Strings are immutable — this creates garbage that is cleaned up when s goes out of scope
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
