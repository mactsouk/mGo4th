package allocate

import (
	"bytes"
	"io"
)

func writeMessageBuffer(msg []byte, b bytes.Buffer) {
	b.Write(msg)
}

func writeMessageBufferPointer(msg []byte, b *bytes.Buffer) {
	b.Write(msg)
}

func writeMessageBufferWriter(msg []byte, b io.Writer) {
	b.Write(msg)
}
