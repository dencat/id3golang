package tests

import (
	"github.com/dencat/id3golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadId3v1(t *testing.T) {
	asrt := assert.New(t)
	id3, err := id3golang.ReadFile("tests/id3v1.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	title, ok := id3.GetTitle()
	asrt.Equal(true, ok)
	asrt.Equal("TITLE1234567890123456789012345", title)

	album, ok := id3.GetAlbum()
	asrt.Equal(true, ok)
	asrt.Equal("ALBUM1234567890123456789012345", album)

	artist, ok := id3.GetArtist()
	asrt.Equal(true, ok)
	asrt.Equal("ARTIST123456789012345678901234", artist)

	year, ok := id3.GetYear()
	asrt.Equal(true, ok)
	asrt.Equal(2001, year)

	_, ok = id3.GetComment()
	asrt.Equal(true, ok)
	// TODO
	//asrt.Equal("COMMENT12345678901234567890123", comment.Text)

	genre, ok := id3.GetGenre()
	asrt.Equal(true, ok)
	asrt.Equal("Pop", genre)

}
