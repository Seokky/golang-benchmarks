package golang_benchmarks

import (
	"strings"
	"testing"
)

// go test -bench=. -benchmem ./reverse_string_test.go

func SimpleReverse(s string) string {
	var str string
	for _, v := range s {
		str = string(v) + str
	}
	return str
}

func BenchmarkSimpleReverse(b *testing.B) {
	for b.Loop() {
		SimpleReverse("Hello World!")
	}
}

func SwapReverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func BenchmarkSwapReverse(b *testing.B) {
	for b.Loop() {
		SwapReverse("Hello World!")
	}
}

func UsingStrBuilderToReverse(s string) string {
	var str strings.Builder
	str.Grow(len(s) - 1)
	for idx := range s {
		str.Write([]byte{s[(len(s)-1)-idx]})
	}
	return str.String()
}

func BenchmarkUsingStrBuilderToReverse(b *testing.B) {
	for b.Loop() {
		UsingStrBuilderToReverse("Hello World!")
	}
}

func RunesReverse(s string) string {
	rune_arr := []rune(s)
	var rev []rune
	for i := len(rune_arr) - 1; i >= 0; i-- {
		rev = append(rev, rune_arr[i])
	}
	return string(rev)
}

func BenchmarkRunesReverse(b *testing.B) {
	for b.Loop() {
		RunesReverse("Hello World!")
	}
}

func ReadFromRightReverse(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

func BenchmarkReadFromRightReverse(b *testing.B) {
	for b.Loop() {
		ReadFromRightReverse("Hello World!")
	}
}

func ReadFromRightBytesReverse(s string) string {
	var bytes strings.Builder
	bytes.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		bytes.WriteByte(s[i])
	}
	return bytes.String()
}

func BenchmarkReadFromRightBytesReverse(b *testing.B) {
	for b.Loop() {
		ReadFromRightBytesReverse("Hello World!")
	}
}
