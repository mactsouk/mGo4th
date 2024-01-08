package allocate

import (
	"bytes"
	"testing"
)

func BenchmarkWBuf(b *testing.B) {
	msg := []byte("Mastering Go!")
	buffer := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		for k := 0; k < 50; k++ {
			writeMessageBuffer(msg, buffer)
		}
	}
}

func BenchmarkWBufPointerNoReset(b *testing.B) {
	msg := []byte("Mastering Go!")
	buffer := new(bytes.Buffer)

	for i := 0; i < b.N; i++ {
		for k := 0; k < 50; k++ {
			writeMessageBufferPointer(msg, buffer)
		}
	}
}

func BenchmarkWBufPointerReset(b *testing.B) {
	msg := []byte("Mastering Go!")
	buffer := new(bytes.Buffer)

	for i := 0; i < b.N; i++ {
		for k := 0; k < 50; k++ {
			writeMessageBufferPointer(msg, buffer)
			buffer.Reset()
		}
	}
}

func BenchmarkWBufWriterReset(b *testing.B) {
	msg := []byte("Mastering Go!")
	buffer := new(bytes.Buffer)

	for i := 0; i < b.N; i++ {
		for k := 0; k < 50; k++ {
			writeMessageBufferWriter(msg, buffer)
			buffer.Reset()
		}
	}
}
