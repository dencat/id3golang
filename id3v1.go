package id3golang

import (
	"errors"
	"io"
	"strconv"
)

// fix size - 128 bytes
type ID3v1 struct {
	Type     string // Always 'TAG'
	Title    string // length 30
	Artist   string // length 30
	Album    string // length 30
	Year     int
	Comment  string
	ZeroByte byte
	Track    byte
	Genre    byte
}

func (id3v1 *ID3v1) String() string {
	return "Type: " + id3v1.Type + "\n" +
		"Title: " + id3v1.Title + "\n" +
		"Artist: " + id3v1.Artist + "\n" +
		"Album: " + id3v1.Album + "\n" +
		"Year: " + strconv.Itoa(id3v1.Year) + "\n" +
		"Comment: " + id3v1.Comment + "\n"
}

func (id3 *ID3v1) GetTag(key string) ([]byte, bool) {
	switch key {
	case "Type":
		return []byte(id3.Type), true
	case "Title":
		return []byte(id3.Title), true
	case "Artist":
		return []byte(id3.Artist), true
	case "Album":
		return []byte(id3.Album), true
	case "Year":
		return []byte(strconv.Itoa(id3.Year)), true
	case "Comment":
		return []byte(id3.Comment), true
	case "Genre":
		return []byte{id3.Genre}, true
	}
	return []byte{}, false
}

func (id3 *ID3v1) SetTag(key string, data []byte) bool {
	switch key {
	case "Type":
		id3.Type = string(data)
		return true
	case "Title":
		id3.Title = string(data)
		return true
	case "Artist":
		id3.Artist = string(data)
		return true
	case "Album":
		id3.Album = string(data)
		return true
	case "Year":
		year, err := strconv.Atoi(string(data))
		if err != nil {
			return false
		}
		id3.Year = year
		return true
	case "Comment":
		id3.Comment = string(data)
		return true
	case "Genre":
		if len(data) != 1 {
			return false
		}
		id3.Genre = data[0]
		return true
	}
	return false
}

func (id3 *ID3v1) getTagData(key string) []byte {
	data, _ := id3.GetTag(key)
	return data
}

func (id3 *ID3v1) GetAll() []ID3Tag {
	return []ID3Tag{
		{"Type", id3.getTagData("Type")},
		{"Title", id3.getTagData("Title")},
		{"Artist", id3.getTagData("Artist")},
		{"Album", id3.getTagData("Album")},
		{"Comment", id3.getTagData("Comment")},
		{"Genre", id3.getTagData("Genre")},
	}
}

func checkId3v1(input io.ReadSeeker) (id3Version, error) {
	if input == nil {
		return TypeID3Undefined, errors.New("empty file")
	}

	// id3v1
	data, err := seekAndRead(input, -128, io.SeekEnd, 3)
	if err != nil {
		return TypeID3Undefined, err
	}
	marker := string(data)
	if marker == "TAG" {
		return TypeID3v1, nil
	}

	return TypeID3Undefined, errors.New("Unsupported format")
}

func readHeaderID3v1(input io.ReadSeeker) (*ID3v1, error) {
	header := ID3v1{}
	if input == nil {
		return nil, errors.New("empty file")
	}

	// Header size
	_, err := input.Seek(-128, io.SeekEnd)
	if err != nil {
		return nil, err
	}

	headerByte := make([]byte, 128)
	nReaded, err := input.Read(headerByte)
	if err != nil {
		return nil, err
	}
	if nReaded != 128 {
		return nil, errors.New("error header length")
	}

	// Type
	marker := string(headerByte[0:3])
	if marker != "TAG" {
		return nil, errors.New("error file marker")
	}
	header.Type = marker

	// Title
	header.Title = string(headerByte[3:33])

	// Artist
	header.Artist = string(headerByte[33:63])

	// Album
	header.Album = string(headerByte[63:93])

	// Year
	header.Year, err = strconv.Atoi(string(headerByte[93:97]))
	if err != nil {
		return nil, errors.New("error year")
	}

	// Comment
	header.Comment = string(headerByte[97:127])

	// Genre
	header.Genre = headerByte[127]

	return &header, nil
}
