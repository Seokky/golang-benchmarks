package golang_benchmarks

// go test -bench=. -benchmem ./unsafe_bytes_to_string_test.go

import (
	"testing"
	"unsafe"
)

func ToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func BenchmarkSafeToString(b *testing.B) {
	bytes := []byte("Hello World")
	for i := 0; i < b.N; i++ {
		_ = string(bytes)
	}
}

func BenchmarkUnsafeToString(b *testing.B) {
	bytes := []byte("Hello World")
	for i := 0; i < b.N; i++ {
		_ = ToString(bytes)
	}
}

// for i := 0; i < b.N; i++
//BenchmarkSafeToString-10        453318574                2.501 ns/op           0 B/op          0 allocs/op
//BenchmarkUnsafeToString-10      1000000000               0.3118 ns/op          0 B/op          0 allocs/op

// for b.Loop()
//BenchmarkSafeToString-10        471675787                2.509 ns/op           0 B/op          0 allocs/op
//BenchmarkUnsafeToString-10      588611592                2.066 ns/op           0 B/op          0 allocs/op
