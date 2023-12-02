package internal

import (
	"bytes"
	"io"
)

func ByteReaderFunc(data []byte) func() io.Reader {
	return func() io.Reader {
		return bytes.NewReader(data)
	}
}
