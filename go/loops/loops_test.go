package loops 

import "testing"

const N = 1000
var a [N]int

func f1(a *[N]int) {
	for i := range a {
		a[i] = i
	}
}

func f2(a *[N]int) {
	_ = *a 
	for i := range a {
		a[i] = i
	}
}

func Benchmark_f1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1(&a)
	}
}

func Benchmark_f2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f2(&a)
	}
}