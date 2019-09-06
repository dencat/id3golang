package id3golang

import (
	"errors"
	"os"
)

func Save(id3 *ID3, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	switch id3.version {
	case TypeID3v1:

	case TypeID3v22, TypeID3v23, TypeID3v24:
		// write header
		err = writeHeaderId3v2(id3, file)
		if err != nil {
			return err
		}

		// write tags
		for key, value := range id3.tags.GetAll() {
			err = writeTagHeader(key, value, file)
			if err != nil {
				return err
			}
		}
	default:
		err = errors.New("Unsupported format")
	}

	return err
}

func writeHeaderId3v2(id3 *ID3, file *os.File) error {
	headerByte := make([]byte, 10)

	// ID3
	headerByte[0] = 'I'
	headerByte[1] = 'D'
	headerByte[2] = '3'

	// Version
	switch id3.version {
	case TypeID3v22:
		headerByte[3] = 2
	case TypeID3v23:
		headerByte[3] = 3
	case TypeID3v24:
		headerByte[3] = 4
	default:
		headerByte[3] = 0
	}

	// Subversion
	headerByte[4] = 0

	// Flags
	headerByte[5] = 0

	// Length
	length := getLength(id3)
	headerByte[6] = byte(length >> 24)
	headerByte[7] = byte(length >> 16)
	headerByte[8] = byte(length >> 8)
	headerByte[9] = byte(length)

	nWriten, err := file.Write(headerByte)
	if err != nil {
		return err
	}
	if nWriten != 10 {
		return errors.New("Writing error")
	}
	return nil
}

func getLength(id3 *ID3) int {
	result := 10
	for _, value := range id3.tags.GetAll() {
		// 10 - size of tag header
		result += 10 + len(value)
	}
	return result
}

func writeTagHeader(key string, value []byte, file *os.File) error {
	return nil
}