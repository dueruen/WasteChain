package compress

import (
	"bytes"
	"testing"
)

func TestCompressDecompress(t *testing.T) {
	data := "This is a test"
	compressed, err := Compress(data)
	if err != nil {
		t.Error(err)
	}
	decompressed, err := Decompress(compressed)
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare([]byte(data), decompressed) != 0 {
		t.Error("Not equal")
	}
}
