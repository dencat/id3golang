package id3golang

import (
	"errors"
	"io"
	"strconv"
)

// fix size - 128 bytes
type ID3v1 struct {
	Type     string // Always 'TAG'
	Title    string // length 30. 30 characters of the title
	Artist   string // length 30. 30 characters of the artist name
	Album    string // length 30. 	30 characters of the album name
	Year     int    // length 4. A four-digit year.
	Comment  string // length 28 or 30. The comment.
	ZeroByte byte   // length 1. If a track number is stored, this byte contains a binary 0.
	Track    byte   // length 1. The number of the track on the album, or 0. Invalid, if previous byte is not a binary 0.
	Genre    byte   // length 1. Index in a list of genres, or 255
}

func (id3v1 *ID3v1) String() string {
	var trackNumber string
	if id3v1.ZeroByte == 0 {
		trackNumber = "TrackNumber: " + strconv.Itoa(int(id3v1.Track)) + "\n"
	}

	return "Type: " + id3v1.Type + "\n" +
		"Title: " + id3v1.Title + "\n" +
		"Artist: " + id3v1.Artist + "\n" +
		"Album: " + id3v1.Album + "\n" +
		"Year: " + strconv.Itoa(id3v1.Year) + "\n" +
		"Comment: " + id3v1.Comment + "\n" +
		trackNumber
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
	case "TrackNumber":
		if id3.ZeroByte == 0 {
			// like a year
			return []byte(strconv.Itoa(int(id3.Track))), true
		}
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
	case "TrackNumber":
		id3.ZeroByte = 0
		trackNumber, err := strconv.Atoi(string(data))
		if err != nil {
			return false
		}
		id3.Track = byte(trackNumber)
		return true
	}
	return false
}

func (id3 *ID3v1) DeleteTag(key string) {
	switch key {
	case "Type":
		id3.Type = ""
	case "Title":
		id3.Title = ""
	case "Artist":
		id3.Artist = ""
	case "Album":
		id3.Album = ""
	case "Year":
		id3.Year = 0
	case "Comment":
		id3.Comment = ""
	case "Genre":
		id3.Genre = 0
	case "TrackNumber":
		id3.ZeroByte = 1
		id3.Track = 0
	}
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

func checkId3v1(input io.ReadSeeker) (Id3Version, error) {
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
	// The track number is stored in the last two bytes of the comment field. If the comment is 29 or 30 characters long, no track number can be stored
	if headerByte[125] == 0 {
		header.Comment = string(headerByte[97:125])
		header.ZeroByte = 0
		header.Track = headerByte[126]
	} else {
		header.Comment = string(headerByte[97:127])
		header.ZeroByte = headerByte[125]
		header.Track = 0
	}

	// Genre
	// Index in a list of genres, or 255
	header.Genre = headerByte[127]

	return &header, nil
}
