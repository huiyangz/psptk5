package service

import (
	"fmt"
	"os"

	"github.com/huiyangz/psptk5/pkg/model"
)

func getFileExtension(buffer []byte) (ext string) {
	ext = "bin"
	if string(buffer[0:4]) == ".PNG" {
		ext = "png"
	} else if string(buffer[0:4]) == "TK5P" {
		ext = "tk5p"
	} else if string(buffer[0:4]) == "tai5" {
		ext = "tai5"
	} else if string(buffer[0:3]) == "TOD" {
		ext = "tod"
	} else if string(buffer[0:3]) == "OMG" {
		ext = "gmo"
	} else if string(buffer[0:4]) == "ILNK" {
		ext = "ilnk"
	} else if string(buffer[0:4]) == "KSEF" {
		ext = "ksef"
	} else if string(buffer[0:4]) == "PPHD" {
		ext = "pphd"
	} else if string(buffer[0:4]) == "PSMF" {
		ext = "psmf"
	} else if string(buffer[0:4]) == "RIFF" {
		ext = "riff"
	} else if string(buffer[0:4]) == "CMPS" {
		ext = "cmps"
	}
	return
}

func writeFile(dirToExtract string, tk5File *model.TK5File) error {
	ext := getFileExtension(tk5File.OriginBuffer)

	if err := os.MkdirAll(dirToExtract, os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s/%08X.%s", dirToExtract, tk5File.BaseAddr, ext))
	if err != nil {
		return err
	}
	defer file.Close()
	if tk5File.DecryptBuffer != nil {
		if _, err := file.Write(tk5File.DecryptBuffer); err != nil {
			return err
		}
	} else {
		if _, err := file.Write(tk5File.OriginBuffer); err != nil {
			return err
		}
	}
	return nil
}
