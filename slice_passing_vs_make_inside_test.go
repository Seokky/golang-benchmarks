package golang_benchmarks

// go test -bench=. -benchmem ./slice_passing_vs_make_inside_test.go

import "testing"

const bufferSize = 30

type ReaderWithSliceArgument struct{}

func (r ReaderWithSliceArgument) Read(p []byte) (int, error) {
	for i := 0; i < bufferSize; i++ {
		p[i] = byte(i)
	}

	return bufferSize, nil
}

type ReaderWithSliceReturn struct{}

func (r ReaderWithSliceReturn) Read(n int) ([]byte, error) {
	p := make([]byte, n)
	for i := 0; i < n; i++ {
		p[i] = byte(i)
	}

	return p, nil
}

// Different performance with for-i

func BenchmarkSliceWithArgument(b *testing.B) {
	for b.Loop() {
		p := make([]byte, bufferSize)
		reader := ReaderWithSliceArgument{}
		_, _ = reader.Read(p)
	}
}

func BenchmarkSliceWithReturn(b *testing.B) {
	for b.Loop() {
		reader := ReaderWithSliceReturn{}
		_, _ = reader.Read(bufferSize)
	}
}

//BenchmarkSliceWithArgument-10           100000000               11.65 ns/op            0 B/op          0 allocs/op
//BenchmarkSliceWithReturn-10             48465021                24.02 ns/op           32 B/op          1 allocs/op
