package id3golang

import (
	"errors"
	"io"
	"strconv"
)

type Id3v2Version byte

const (
	ID3v22 Id3v2Version = 1
	ID3v23 Id3v2Version = 2
	ID3v24 Id3v2Version = 3
)

func (id3 Id3v2Version) String() string {
	switch id3 {
	case ID3v22:
		return "ID3v2.2"
	case ID3v23:
		return "ID3v2.3"
	case ID3v24:
		return "ID3v2.4"
	}
	return ""
}

type Id3v2Flags byte

const (
	Id3FlagUnsynchronisation     Id3v2Flags = 1
	Id3FlagExtendedheader        Id3v2Flags = 2
	Id3FlagExperimantalindicator Id3v2Flags = 3
)

func (flags *Id3v2Flags) String() string {
	return strconv.Itoa(int(*flags))
}

type ID3v2 struct {
	Marker     string // Always 'ID3'
	Version    Id3v2Version
	SubVersion int
	Flags      Id3v2Flags
	Length     int
	Tags       []ID3Tag
}

func (id3v2 *ID3v2) String() string {
	result := "Marker: " + id3v2.Marker + "\n" +
		"Version: " + id3v2.Version.String() + "\n" +
		"Subversion: " + strconv.Itoa(id3v2.SubVersion) + "\n" +
		"Flags: " + id3v2.Flags.String() + "\n" +
		"Length: " + strconv.Itoa(id3v2.Length) + "\n"

	for _, frame := range id3v2.Tags {
		result += frame.Key + ": " + string(frame.Value) + "\n"
	}

	return result
}

func checkId3v2(input io.ReadSeeker) (Id3Version, error) {
	if input == nil {
		return TypeID3Undefined, errors.New("empty file")
	}

	// read marker (3 bytes) and version (1 byte) for ID3v2
	data, err := seekAndRead(input, 0, io.SeekStart, 4)
	if err != nil {
		return TypeID3Undefined, err
	}
	marker := string(data[0:3])

	// id3v2
	if marker == "ID3" {
		versionByte := data[3]
		switch versionByte {
		case 2:
			return TypeID3v22, nil
		case 3:
			return TypeID3v22, nil
		case 4:
			return TypeID3v24, nil
		}
	}

	return TypeID3Undefined, errors.New("Unsupported format")
}

func readHeaderID3v2(input io.ReadSeeker) (*ID3v2, error) {
	header := ID3v2{}
	if input == nil {
		return nil, errors.New("empty file")
	}

	// Seek to file start
	startIndex, err := input.Seek(0, io.SeekStart)
	if startIndex != 0 {
		return nil, errors.New("error seek file")
	}

	if err != nil {
		return nil, err
	}

	// Header size
	headerByte := make([]byte, 10)
	nReaded, err := input.Read(headerByte)
	if err != nil {
		return nil, err
	}
	if nReaded != 10 {
		return nil, errors.New("error header length")
	}

	// Marker
	marker := string(headerByte[0:3])
	if marker != "ID3" {
		return nil, errors.New("error file marker")
	}
	header.Marker = marker

	// Version
	versionByte := headerByte[3]
	switch versionByte {
	case 2:
		header.Version = ID3v22
	case 3:
		header.Version = ID3v23
	case 4:
		header.Version = ID3v24
	default:
		return nil, errors.New("error file version")
	}

	// Sub version
	subVersionByte := headerByte[4]
	header.SubVersion = int(subVersionByte)

	// Flags
	header.Flags = Id3v2Flags(headerByte[5])

	// Length
	length := ByteToIntSynchsafe(headerByte[6:10])
	header.Length = length

	// Extended headers
	header.Tags = []ID3Tag{}
	curRead := 0
	for curRead < length {
		bytesExtendedHeader := make([]byte, 10)
		nReaded, err = input.Read(bytesExtendedHeader)
		if err != nil {
			return nil, err
		}
		if nReaded != 10 {
			return nil, errors.New("error extended header length")
		}
		// Frame identifier
		key := string(bytesExtendedHeader[0:4])

		//if bytesExtendedHeader[0] == 0 && bytesExtendedHeader[1] == 0 && bytesExtendedHeader[2] == 0 {
		//	break
		//}

		// Frame data size
		size := ByteToInt(bytesExtendedHeader[4:8])

		bytesExtendedValue := make([]byte, size)
		nReaded, err = input.Read(bytesExtendedValue)
		if err != nil {
			return nil, err
		}
		if nReaded != size {
			return nil, errors.New("error extended value length")
		}

		header.Tags = append(header.Tags, ID3Tag{
			key,
			bytesExtendedValue,
		})

		curRead += 10 + size
	}

	// TODO
	if curRead != length {
		return nil, errors.New("error extended frames")
	}
	return &header, nil
}

func (id3 *ID3v2) GetTag(key string) ([]byte, bool) {
	i := id3.findElement(key)
	if i == -1 {
		return []byte{}, false
	}
	data := id3.Tags[i]
	return data.Value, true
}

func (id3 *ID3v2) SetTag(key string, data []byte) bool {
	i := id3.findElement(key)
	if i == -2 || i == -1 {
		id3.Tags = append(id3.Tags, ID3Tag{
			key,
			data,
		})
	} else {
		id3.Tags[i].Value = data
	}
	return true
}

func (id3 *ID3v2) GetAll() []ID3Tag {
	return id3.Tags
}

// Return index if find element
// Return -1 if not found
// Return -2 if key == ""
func (id3 *ID3v2) findElement(key string) int {
	if key == "" {
		return -2
	}

	for i, val := range id3.Tags {
		if val.Key == key {
			return i
		}
	}
	return -1
}
