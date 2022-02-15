package service

import (
	"github.com/huiyangz/psptk5/pkg/decrypt"
)

func Extract(pathOfDataBin, pathOfEboot, dirToExtract string) error {
	// Read segmentation informations
	files, err := loadFileInfos(pathOfEboot)
	if err != nil {
		return err
	}

	// Read buffers from DATA.BIN
	for _, f := range files {
		if _, err := loadBufferFromDataBin(pathOfDataBin, f); err != nil {
			return err
		}
	}

	// Decrypt buffers of CMPS
	for _, f := range files {
		if string(f.OriginBuffer[0:4]) == "CMPS" {
			f.DecryptBuffer, err = decrypt.DecryptCMPS(f.OriginBuffer)
			if err != nil {
				return err
			}
		}
	}

	// Write files
	for _, f := range files {
		if err := writeFile(dirToExtract, f); err != nil {
			return err
		}
	}

	return nil
}
