package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		echo3()
	}
}
