package decrypt

import (
	"bufio"
	"bytes"

	"github.com/huiyangz/psptk5/pkg/reader"
)

// Reverse from 0x08857b90
func DecryptCMPS(in []byte) ([]byte, error) {
	br := bufio.NewReader(bytes.NewReader(in))
	reader.Seek(br, 4) // CMPS
	reader.Seek(br, 4) // FLAG 1
	size, _ := reader.ReadInt32(br)
	reader.Seek(br, 4)

	out := []byte{}
	s0 := 0xFEE
	s1 := 0
	mapCode := make(map[int]int)

	for size > 0 {
		s1 = s1 >> 1
		v1 := s1 & 0x100

		if v1 == 0 {
			v1, _ = reader.ReadByte(br)
			s1 = v1 | 0xFF00
		}

		v1 = s1 & 0x01
		if v1 != 0 {
			a0, _ := reader.ReadByte(br)
			mapCode[s0] = a0
			v1 = s0 + 0x1
			out = append(out, byte(a0))
			s0 = v1 & 0xFFF
			size--
		} else {
			a1, _ := reader.ReadByte(br)
			a0, _ := reader.ReadByte(br)
			a3 := 0
			v1 = a0 & 0xF0
			v1 = v1 << 0x4
			a2 := a1 | v1
			v1 = a0 & 0xF
			a1 = v1 + 0x2
			for a1 >= a3 && size > 0 {
				v1 = a2 + a3
				v1 = v1 & 0xFFF
				a0, ok := mapCode[v1]
				if !ok {
					a0 = 0
				}
				size--
				a3 = a3 + 0x01
				out = append(out, byte(a0))
				mapCode[s0] = a0

				v1 = s0 + 0x01
				s0 = v1 & 0xFFF
			}
		}
	}

	return out, nil
}
