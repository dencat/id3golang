package tests

import (
	"github.com/dencat/id3golang"
	"github.com/stretchr/testify/assert"
	"image/jpeg"
	"os"
	"testing"
)

func Test_Picture(t *testing.T) {
	asrt := assert.New(t)

	id3, err := id3golang.ReadFile("tests/mozart.mp3")
	if !asrt.NoError(err) {
		return
	}

	picture, ok := id3.GetPicture()
	if !asrt.True(ok) {
		return
	}
	asrt.Equal("image/jpeg", picture.Mime)
	asrt.Equal(id3golang.PictureTypeCoverFront, picture.PictureType)

	img := picture.Image

	out, err := os.Create("tests/mozart1.jpeg")
	if err != nil {
		t.Error(err)
		return
	}
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)
	if err != nil {
		t.Error(err)
		return
	}
}
