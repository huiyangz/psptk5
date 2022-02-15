package service

import (
	"io"
	"os"

	"github.com/huiyangz/psptk5/pkg/model"
	"github.com/huiyangz/psptk5/pkg/reader"
)

func loadFileInfos(pathOfEboot string) ([]*model.TK5File, error) {
	const textAddr = 0x08804000             // Read from EBOOT Section Headers
	const textOff = 0x60                    // Read from EBOOT Section Headers
	const fileInfoAddrInMemory = 0x08BCF050 // Fixed address
	const fileInfoAddrInEBOOT = fileInfoAddrInMemory - textAddr + textOff
	const fileCount = 0x314 // Fixed value

	f, err := os.Open(pathOfEboot)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sr := io.NewSectionReader(f, fileInfoAddrInEBOOT, fileCount*4)
	files := make([]*model.TK5File, 0)
	cursor, _ := reader.ReadInt32(sr)
	for i := 1; i < fileCount; i++ {
		size, _ := reader.ReadInt32(sr)
		files = append(files, &model.TK5File{
			BaseAddr: int64(cursor << 0xB),
			Size:     int64(size),
		})
		cursor += (size + 0x7FF) >> 0xB
	}
	return files, nil
}
