package id3golang

import (
	"errors"
	"io"
	"os"
)

func SaveFile(id3 *ID3, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return Save(id3, file)
}

func Save(id3 *ID3, writer io.Writer) error {
	var err error

	switch id3.version {
	case TypeID3v1:

	case TypeID3v22, TypeID3v23, TypeID3v24:
		// write header
		err = writeHeaderId3v2(id3, writer)
		if err != nil {
			return err
		}

		// write tags
		for _, tag := range id3.tags.GetAll() {
			err = writeTagHeader(tag.Key, tag.Value, writer)
			if err != nil {
				return err
			}
		}

		// write data
		_, err = writer.Write(id3.data)
		if err != nil {
			return err
		}
	default:
		err = errors.New("Unsupported format")
	}

	return nil
}

func writeHeaderId3v2(id3 *ID3, writer io.Writer) error {
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
	lengthByte := IntToByteSynchsafe(length)
	headerByte[6] = lengthByte[0]
	headerByte[7] = lengthByte[1]
	headerByte[8] = lengthByte[2]
	headerByte[9] = lengthByte[3]

	nWriten, err := writer.Write(headerByte)
	if err != nil {
		return err
	}
	if nWriten != 10 {
		return errors.New("Writing error")
	}
	return nil
}

func getLength(id3 *ID3) int {
	result := 0
	for _, tag := range id3.tags.GetAll() {
		// 10 - size of tag header
		result += 10 + len(tag.Value)
	}
	return result
}

func writeTagHeader(key string, value []byte, writer io.Writer) error {
	header := make([]byte, 10)

	// Frame id
	for i, val := range key {
		header[i] = byte(val)
	}

	// Frame size
	length := len(value)
	header[4] = byte(length >> 24)
	header[5] = byte(length >> 16)
	header[6] = byte(length >> 8)
	header[7] = byte(length)

	// write header
	_, err := writer.Write(header)
	if err != nil {
		return err
	}

	// write data
	_, err = writer.Write(value)
	if err != nil {
		return err
	}

	return nil
}
