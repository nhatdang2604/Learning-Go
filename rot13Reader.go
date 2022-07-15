package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b = 'A' + (b-'A'+13)%26
	} else if b >= 'a' && b <= 'z' {
		b = 'a' + (b-'a'+13)%26
	}
	return b
}

func (reader rot13Reader) Read(buffer []byte) (byte_count int, err error) {

	byte_count, err = reader.r.Read(buffer)

	for i, val := range buffer {
		buffer[i] = rot13(val)
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
