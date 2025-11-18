package golang_benchmarks

// go test -bench=. -benchmem ./unsafe_bytes_to_string_test.go

import (
	"testing"
	"unsafe"
)

func BenchmarkSafeToString(b *testing.B) {
	bytes := []byte("Hello World")
	for b.Loop() {
		result = string(bytes)
	}
}

func BenchmarkUnsafeToStringWithPtr(b *testing.B) {
	bytes := []byte("Hello World")
	for b.Loop() {
		result = ToStringWithUnsafePointer(bytes)
	}
}

func BenchmarkUnsafeToStringWithoutPtr(b *testing.B) {
	bytes := []byte("Hello World")
	for b.Loop() {
		result = ToStringWithUnsafeString(bytes)
	}
}

func ToStringWithUnsafePointer(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func ToStringWithUnsafeString(bytes []byte) string {
	if len(bytes) == 0 {
		return ""
	}

	return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}

var result string

// BenchmarkSafeToString-10                        98654934                12.25 ns/op           16 B/op          1 allocs/op
// BenchmarkUnsafeToStringWithPtr-10               585261528                2.056 ns/op           0 B/op          0 allocs/op
// BenchmarkUnsafeToStringWithoutPtr-10            592686736                2.029 ns/op           0 B/op          0 allocs/op
