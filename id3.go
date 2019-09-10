package id3golang

import (
	"strconv"
	"time"
)

type ID3 struct {
	// id3 version
	version Id3Version
	// Tag information
	tags TagReader
	// All another information
	data []byte
}

func (id3 *ID3) GetVersion() Id3Version {
	return id3.version
}

func (id3 *ID3) SetVersion(version Id3Version) {
	id3.version = version
}

func (id3 *ID3) GetAllTagsName() []string {
	result := []string{}
	allTags := id3.tags.GetAll()
	for i := range allTags {
		result = append(result, string(realTagNameToName(allTags[i].Key, id3.version)))
	}
	return result
}

func (id3 *ID3) DeleteAllTags() {
	allTags := id3.tags.GetAll()
	for i := range allTags {
		id3.tags.DeleteTag(allTags[i].Key)
	}
}

func (id3 *ID3) GetAllTags() []ID3Tag {
	return id3.tags.GetAll()
}

func (id3 *ID3) GetTitle() (string, bool) {
	return id3.getString(tagTitle)
}

func (id3 *ID3) SetTitle(title string) bool {
	return id3.setString(tagTitle, title)
}

func (id3 *ID3) DeleteTitle() {
	id3.DeleteTag(tagTitle)
}

func (id3 *ID3) GetArtist() (string, bool) {
	return id3.getString(tagArtist)
}

func (id3 *ID3) SetArtist(artist string) bool {
	return id3.setString(tagArtist, artist)
}

func (id3 *ID3) DeleteArtist() {
	id3.DeleteTag(tagArtist)
}

func (id3 *ID3) GetAlbum() (string, bool) {
	return id3.getString(tagAlbum)
}

func (id3 *ID3) SetAlbum(album string) bool {
	return id3.setString(tagAlbum, album)
}

func (id3 *ID3) DeleteAlbum() {
	id3.DeleteTag(tagAlbum)
}

func (id3 *ID3) GetYear() (int, bool) {
	year, ok := id3.getTimestamp(tagYear)
	if !ok {
		return 0, false
	}
	return year.Year(), true
}

type Comment struct {
	Language string
	Text     string
}

func (id3 *ID3) GetComment() (*Comment, bool) {
	commentStr, ok := id3.getString(tagComment)
	if !ok {
		return nil, false
	}
	// Comment struct must be greater than 4
	// [lang \x00 text] - comment format
	// lang - 3 symbols
	// \x00 - const, delimeter
	// text - all after
	if len(commentStr) < 4 {
		return nil, false
	}
	return &Comment{
		Language: commentStr[0:3],
		Text:     commentStr[4:],
	}, true
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

func (id3 *ID3) setString(name tagName, data string) bool {
	// Set UTF-8
	result := []byte{0}
	realName := getTagName(name, id3.version)
	// append data
	result = append(result, []byte(data)...)
	return id3.tags.SetTag(realName, result)
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

func (id3 *ID3) DeleteTag(name tagName) {
	realTagName := getTagName(name, id3.version)
	id3.tags.DeleteTag(realTagName)
}
