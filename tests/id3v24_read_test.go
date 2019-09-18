package tests

import (
	"github.com/dencat/id3golang"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReadId3v24(t *testing.T) {
	asrt := assert.New(t)
	id3, err := id3golang.ReadFile("tests/meow_id2.4.mp3")
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

	date, ok := id3.GetDate()
	asrt.Equal(true, ok)
	asrt.Equal(time.Date(2008, time.September, 15, 15, 53, 0, 0, time.UTC), date)

	arranger, ok := id3.GetArranger()
	asrt.Equal(true, ok)
	asrt.Equal("CK", arranger)

	author, ok := id3.GetAuthor()
	asrt.Equal(true, ok)
	asrt.Equal("Kitten", author)

	bpm, ok := id3.GetBPM()
	asrt.Equal(true, ok)
	asrt.Equal(777, bpm)

	/* TODO: long keys TXXX:CATALOGNUMBER
	catalogNumber, ok := id3.GetCatalogNumber()
	asrt.Equal(true, ok)
	asrt.Equal("TITLE1234567890123456789012345", catalogNumber)*/

	compilation, ok := id3.GetCompilation()
	asrt.Equal(true, ok)
	asrt.Equal("catcomp", compilation)

	composer, ok := id3.GetComposer()
	asrt.Equal(true, ok)
	asrt.Equal("catcomposer", composer)

	conductor, ok := id3.GetConductor()
	asrt.Equal(true, ok)
	asrt.Equal("catconductor", conductor)

	copyright, ok := id3.GetCopyright()
	asrt.Equal(true, ok)
	asrt.Equal("2019", copyright)

	description, ok := id3.GetDescription()
	asrt.Equal(true, ok)
	asrt.Equal("subtitle", description)

	discNumber, ok := id3.GetDiscNumber()
	asrt.Equal(true, ok)
	asrt.Equal("1/7", discNumber)

	encodedBy, ok := id3.GetEncodedBy()
	asrt.Equal(true, ok)
	asrt.Equal("encodedbycat", encodedBy)

	trackNumber, ok := id3.GetTrackNumber()
	asrt.Equal(true, ok)
	asrt.Equal(12, trackNumber)
}
