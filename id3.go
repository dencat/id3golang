package id3golang

import (
	"io"
	"os"
)

type ID3 struct {
	Version ID3Version
	Tags    map[string][]byte
}

type ID3Version byte

const (
	ID3Undefined ID3Version = 0
	ID3v1        ID3Version = 1
	ID3v22       ID3Version = 2
	ID3v23       ID3Version = 3
	ID3v24       ID3Version = 4
)

func ReadFile(path string) (*ID3, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Read(file)
}

func Read(input io.ReadSeeker) (*ID3, error) {
	var id3 ID3

	return &id3, nil
}
