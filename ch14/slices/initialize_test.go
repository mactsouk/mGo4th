package main

import (
	"testing"
)

var t []int

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = InitSliceNew(i)
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = InitSliceAppend(i)
	}
}
