package tests

import (
	"github.com/dencat/id3golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeHeaders(t *testing.T) {
	asrt := assert.New(t)
	id3, err := id3golang.ReadFile("tests/meow_id2.4.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	id3.SetTitle("gav gav")
	id3.SetAlbum("dog stars")
	id3.SetArtist("dog")

	err = id3golang.SaveFile(id3, "tests/meow_id2.4.change.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	id3, err = id3golang.ReadFile("tests/meow_id2.4.change.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	title, ok := id3.GetTitle()
	asrt.Equal(true, ok)
	asrt.Equal("gav gav", title)

	album, ok := id3.GetAlbum()
	asrt.Equal(true, ok)
	asrt.Equal("dog stars", album)

	artist, ok := id3.GetArtist()
	asrt.Equal(true, ok)
	asrt.Equal("dog", artist)
}
