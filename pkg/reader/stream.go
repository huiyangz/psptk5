package reader

import (
	"encoding/binary"
	"io"
)

func Seek(r io.Reader, size int) error {
	buffer := make([]byte, size)
	n, err := r.Read(buffer)
	if err != nil {
		return err
	}
	if n != size {
		return io.EOF
	}
	return nil
}

func ReadByte(reader io.Reader) (int, error) {
	var n byte
	err := binary.Read(reader, binary.LittleEndian, &n)
	return int(n), err
}

func ReadInt32(reader io.Reader) (int, error) {
	var n int32
	err := binary.Read(reader, binary.LittleEndian, &n)
	return int(n), err
}
