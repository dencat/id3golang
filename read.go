package id3golang

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
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
	var err error

	id3.version, err = readId3Version(input)
	switch id3.version {
	case TypeID3v1:
		id3.tags, err = readHeaderID3v1(input)
	case TypeID3v22, TypeID3v23, TypeID3v24:
		// read headers
		id3.tags, err = readHeaderID3v2(input)
		if err != nil {
			return &id3, err
		}
		// read another data
		id3.data, err = ioutil.ReadAll(input)
	default:
		err = errors.New("Unsupported format")
	}

	return &id3, err
}

func readId3Version(input io.ReadSeeker) (Id3Version, error) {
	// check id3v2
	versionId3v2, err := checkId3v2(input)
	if versionId3v2 != TypeID3Undefined && err == nil {
		return versionId3v2, nil
	}

	// check id3v1
	versionId3v1, err := checkId3v1(input)
	if versionId3v1 != TypeID3Undefined && err == nil {
		return versionId3v1, nil
	}

	return TypeID3Undefined, errors.New("Unsupported format")
}

func seekAndRead(input io.ReadSeeker, offset int64, whence int, read int) ([]byte, error) {
	startIndex, err := input.Seek(offset, whence)
	if startIndex != 0 || err != nil {
		return nil, errors.New("error seek file")
	}

	data := make([]byte, read)
	nReaded, err := input.Read(data)
	if err != nil {
		return nil, err
	}
	if nReaded != 4 {
		return nil, errors.New("error header length")
	}

	return data, nil
}
