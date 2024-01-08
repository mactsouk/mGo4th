package allocate

import (
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	msg := []byte("Mastering Go!")
	for i := 0; i < b.N; i++ {
		for k := 0; k < 50; k++ {
			writeMessage(msg)
		}
	}
}
