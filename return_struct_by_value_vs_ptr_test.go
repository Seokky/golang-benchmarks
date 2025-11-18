package golang_benchmarks

import "testing"

// go test -bench=. -benchmem ./return_struct_by_value_vs_ptr_test.go

type Child struct {
	a int64
	b [100]int
}
type Data struct {
	a int
	b string
	c bool
	d Child
}

func NewDataByValue() Data {
	return Data{}
}

func NewDataByPointer() *Data {
	return &Data{}
}

func BenchmarkReturnByValue(b *testing.B) {
	for b.Loop() {
		NewDataByValue()
	}
}

func BenchmarkReturnByPointer(b *testing.B) {
	for b.Loop() {
		NewDataByPointer()
	}
}

// BenchmarkReturnByValue-10       82020322                14.46 ns/op            0 B/op          0 allocs/op
// BenchmarkReturnByPointer-10      8881662               137.3 ns/op           896 B/op          1 allocs/op
