package main

import (
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = InitSliceNew(i)
	}
}

func BenchmarkAppend(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = InitSliceAppend(i)
	}
}
