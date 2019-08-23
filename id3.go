package id3golang

import (
	"strconv"
	"time"
)

type ID3 struct {
	version id3Version
	tags    TagReader
}

func (id3 *ID3) GetTitle() (string, bool) {
	return id3.getString(tagTitle)
}

func (id3 *ID3) GetArtist() (string, bool) {
	return id3.getString(tagArtist)
}

func (id3 *ID3) GetAlbum() (string, bool) {
	return id3.getString(tagAlbum)
}

func (id3 *ID3) GetYear() (int, bool) {
	year, ok := id3.getTimestamp(tagYear)
	if !ok {
		return 0, false
	}
	return year.Year(), true
}

func (id3 *ID3) GetComment() (string, bool) {
	return id3.getString(tagComment)
}

func (id3 *ID3) GetGenre() (string, bool) {
	return id3.getString(tagGenre)
}

func (id3 *ID3) GetAlbumArtist() (string, bool) {
	return id3.getString(tagAlbumArtist)
}

func (id3 *ID3) getString(name tagName) (string, bool) {
	data, ok := id3.getData(name)
	if !ok {
		return "", false
	}
	tagType := getTagType(name, id3.version)
	switch tagType {
	case TAG_TYPE_TEXT:
		switch id3.version {
		case TypeID3v1:
			return string(data), true
		case TypeID3v24:
			result, err := decodeString(data[1:], textEncoding(data))
			if err != nil {
				return "", false
			}
			return result, true
		}
	case TAG_TYPE_INT:
		switch id3.version {
		case TypeID3v1:
			// check for error
			number, err := strconv.Atoi(string(data))
			if err != nil {
				return "", false
			}
			return strconv.Itoa(number), true
		}
	}
	return "", false
}

func (id3 *ID3) getInt(name tagName) (int, bool) {
	data, ok := id3.getData(name)
	if !ok {
		return 0, false
	}
	tagType := getTagType(name, id3.version)
	switch tagType {
	case TAG_TYPE_TEXT:
		str, ok := id3.getString(name)
		if !ok {
			return 0, false
		}
		number, err := strconv.Atoi(str)
		if err != nil {
			return 0, false
		}
		return number, true
	case TAG_TYPE_INT:
		switch id3.version {
		case TypeID3v1:
			number, err := strconv.Atoi(string(data))
			if err != nil {
				return 0, false
			}
			return number, true
		}
	}
	return 0, false
}

func (id3 *ID3) getTimestamp(name tagName) (time.Time, bool) {
	str, ok := id3.getString(name)
	if !ok {
		return time.Now(), false
	}
	result, err := time.Parse("2006-01-02T15:04:05", str)
	if err != nil {
		return time.Now(), false
	}
	return result, true
}

func (id3 *ID3) getData(name tagName) ([]byte, bool) {
	realTagName := getTagName(name, id3.version)
	return id3.tags.GetTag(realTagName)
}