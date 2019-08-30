package tests

import (
	"github.com/stretchr/testify/assert"
	"id3golang"
	"testing"
)

func TestId3v24(t *testing.T) {
	asrt := assert.New(t)
	id3, err := id3golang.ReadFile("meow_id2.4.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	title, ok := id3.GetTitle()
	asrt.Equal(true, ok)
	asrt.Equal("MEOW", title)

	album, ok := id3.GetAlbum()
	asrt.Equal(true, ok)
	asrt.Equal("CatAlbum", album)

	artist, ok := id3.GetArtist()
	asrt.Equal(true, ok)
	asrt.Equal("Cute Kitten", artist)

	year, ok := id3.GetYear()
	asrt.Equal(true, ok)
	asrt.Equal(2008, year)

	comment, ok := id3.GetComment()
	asrt.Equal(true, ok)
	asrt.Equal("catcomment", comment.Text)

	genre, ok := id3.GetGenre()
	asrt.Equal(true, ok)
	asrt.Equal("catmusic", genre)

	albumArtist, ok := id3.GetAlbumArtist()
	asrt.Equal(true, ok)
	asrt.Equal("CatAlbumArtist", albumArtist)
}
