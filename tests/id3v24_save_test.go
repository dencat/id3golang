package tests

import (
	"bytes"
	"github.com/dencat/id3golang"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSave(t *testing.T) {
	asrt := assert.New(t)
	id3, err := id3golang.ReadFile("tests/meow_id2.4.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	err = id3golang.SaveFile(id3, "tests/meow_id2.4.save.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	cmp := compareFiles("tests/meow_id2.4.mp3", "tests/meow_id2.4.save.mp3")
	asrt.Equal(true, cmp)
}

func compareFiles(path1, path2 string) bool {
	data1, err := ioutil.ReadFile(path1)
	if err != nil {
		return false
	}

	data2, err := ioutil.ReadFile(path2)
	if err != nil {
		return false
	}

	result := bytes.Compare(data1, data2)
	return result == 0
}
