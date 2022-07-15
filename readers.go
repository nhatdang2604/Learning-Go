package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(buffer []byte) (byte_count int, err error) {
	for i := range buffer {
		buffer[i] = 'A'
	}

	byte_count = len(buffer)
	err = nil
	return
}

func main() {
	reader.Validate(MyReader{})
}
