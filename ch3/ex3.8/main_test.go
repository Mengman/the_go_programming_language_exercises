package main

import "testing"

// run benchmark type "go test -bench=."
// 	goos: linux
//	goarch: amd64
//	pkg: github.com/Mengman/the_go_programming_language_exercises/ch3/ex3.8
//	BenchmarkMandelbrot-12                  30000000                40.2 ns/op
//	BenchmarkMandelbrot64-12                30000000                41.4 ns/op
//	BenchmarkMandelbrotBigFloat-12           3000000               440 ns/op
//	BenchmarkMandelbrotRat-12                2000000               909 ns/op
func BenchmarkMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrot64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot64(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotBigFloat(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotRat(complex(float64(i), float64(i)))
	}
}
