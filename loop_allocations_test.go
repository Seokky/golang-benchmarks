package golang_benchmarks

import "testing"

// go test -bench=. -benchmem ./loop_allocations_test.go

func Init(value *int) {
	*value = 1
}

func BenchmarkWithoutLoopAllocation(b *testing.B) {
	var value int
	for b.Loop() {
		Init(&value)
	}
}

func BenchmarkWithLoopAllocation(b *testing.B) {
	for b.Loop() {
		var value int
		Init(&value)
	}
}

// BenchmarkWithoutLoopAllocation-10       581843455                2.025 ns/op           0 B/op          0 allocs/op
// BenchmarkWithLoopAllocation-10          587393434                2.033 ns/op           0 B/op          0 allocs/op
