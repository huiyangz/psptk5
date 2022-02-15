package service

import (
	"io"
	"os"

	"github.com/huiyangz/psptk5/pkg/model"
)

func loadBufferFromDataBin(pathOfDataBin string, tk5File *model.TK5File) (*model.TK5File, error) {
	f, err := os.Open(pathOfDataBin)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := io.NewSectionReader(f, tk5File.BaseAddr, tk5File.Size)
	tk5File.OriginBuffer = make([]byte, tk5File.Size)
	_, err = reader.Read(tk5File.OriginBuffer)
	if err != nil {
		return nil, err
	}

	return tk5File, nil
}
