package golang_benchmarks

import "testing"

// go test -bench=. -benchmem ./loop_allocations_test.go

//go:noinline
func Init(value *int) {
	*value = 1
}

// Different performance with b.Loop()

func BenchmarkWithoutLoopAllocation(b *testing.B) {
	var value int
	for n := 0; n < b.N; n++ {
		Init(&value)
	}
}

func BenchmarkWithLoopAllocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var value int
		Init(&value)
	}
}

//BenchmarkWithoutLoopAllocation-10       1000000000               0.9435 ns/op          0 B/op          0 allocs/op
//BenchmarkWithLoopAllocation-10          593888028                2.036 ns/op           0 B/op          0 allocs/op
