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

func (id3 *ID3) GetAllTags() []ID3Tag {
	result := []ID3Tag{}
	allTags := id3.tags.GetAll()
	for i := range allTags {
		key := realTagNameToName(allTags[i].Key, id3.version)
		result = append(result, ID3Tag{string(key), allTags[i].Value})
	}
	return result
}

func (id3 *ID3) DeleteAllTags() {
	allTags := id3.tags.GetAll()
	for i := range allTags {
		id3.tags.DeleteTag(allTags[i].Key)
	}
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
	switch id3.version {
	case TypeID3v1:
		return id3.getInt(tagYear)
	default:
		year, ok := id3.getTimestamp(tagYear)
		if !ok {
			return 0, false
		}
		return year.Year(), true
	}

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
	switch id3.version {
	case TypeID3v1:
		data, ok := id3.getData(tagGenre)
		if !ok {
			return "", false
		}
		genreCode := int(data[0])
		return genres[genreCode], true
	default:
		return id3.getString(tagGenre)
	}

}

func (id3 *ID3) GetAlbumArtist() (string, bool) {
	return id3.getString(tagAlbumArtist)
}

func (id3 *ID3) GetDate() (string, bool) {
	return id3.getString(tagDate)
}

func (id3 *ID3) SetDate(title string) bool {
	return id3.setString(tagDate, title)
}

func (id3 *ID3) DeleteDate() {
	id3.DeleteTag(tagDate)
}

func (id3 *ID3) GetTrackNumber() (string, bool) {
	return id3.getString(tagTrackNumber)
}

func (id3 *ID3) SetTrackNumber(title string) bool {
	return id3.setString(tagTrackNumber, title)
}

func (id3 *ID3) DeleteTrackNumber() {
	id3.DeleteTag(tagTrackNumber)
}

func (id3 *ID3) GetArranger() (string, bool) {
	return id3.getString(tagArranger)
}

func (id3 *ID3) SetArranger(title string) bool {
	return id3.setString(tagArranger, title)
}

func (id3 *ID3) DeleteArranger() {
	id3.DeleteTag(tagArranger)
}

func (id3 *ID3) GetAuthor() (string, bool) {
	return id3.getString(tagAuthor)
}

func (id3 *ID3) SetAuthor(title string) bool {
	return id3.setString(tagAuthor, title)
}

func (id3 *ID3) DeleteAuthor() {
	id3.DeleteTag(tagAuthor)
}

func (id3 *ID3) GetBPM() (string, bool) {
	return id3.getString(tagBPM)
}

func (id3 *ID3) SetBPM(title string) bool {
	return id3.setString(tagBPM, title)
}

func (id3 *ID3) DeleteBPM() {
	id3.DeleteTag(tagBPM)
}

func (id3 *ID3) GetCatalogNumber() (string, bool) {
	return id3.getString(tagCatalogNumber)
}

func (id3 *ID3) SetCatalogNumber(title string) bool {
	return id3.setString(tagCatalogNumber, title)
}

func (id3 *ID3) DeleteCatalogNumber() {
	id3.DeleteTag(tagCatalogNumber)
}

func (id3 *ID3) GetCompilation() (string, bool) {
	return id3.getString(tagCompilation)
}

func (id3 *ID3) SetCompilation(title string) bool {
	return id3.setString(tagCompilation, title)
}

func (id3 *ID3) DeleteCompilation() {
	id3.DeleteTag(tagCompilation)
}

func (id3 *ID3) GetComposer() (string, bool) {
	return id3.getString(tagComposer)
}

func (id3 *ID3) SetComposer(title string) bool {
	return id3.setString(tagComposer, title)
}

func (id3 *ID3) DeleteComposer() {
	id3.DeleteTag(tagComposer)
}

func (id3 *ID3) GetConductor() (string, bool) {
	return id3.getString(tagConductor)
}

func (id3 *ID3) SetConductor(title string) bool {
	return id3.setString(tagConductor, title)
}

func (id3 *ID3) DeleteConductor() {
	id3.DeleteTag(tagConductor)
}

func (id3 *ID3) GetCopyright() (string, bool) {
	return id3.getString(tagCopyright)
}

func (id3 *ID3) SetCopyright(title string) bool {
	return id3.setString(tagCopyright, title)
}

func (id3 *ID3) DeleteCopyright() {
	id3.DeleteTag(tagCopyright)
}

func (id3 *ID3) GetDescription() (string, bool) {
	return id3.getString(tagDescription)
}

func (id3 *ID3) SetDescription(title string) bool {
	return id3.setString(tagDescription, title)
}

func (id3 *ID3) DeleteDescription() {
	id3.DeleteTag(tagDescription)
}

func (id3 *ID3) GetDiscNumber() (string, bool) {
	return id3.getString(tagDiscNumber)
}

func (id3 *ID3) SetDiscNumber(title string) bool {
	return id3.setString(tagDiscNumber, title)
}

func (id3 *ID3) DeleteDiscNumber() {
	id3.DeleteTag(tagDiscNumber)
}

func (id3 *ID3) GetEncodedBy() (string, bool) {
	return id3.getString(tagEncodedBy)
}

func (id3 *ID3) SetEncodedBy(title string) bool {
	return id3.setString(tagEncodedBy, title)
}

func (id3 *ID3) DeleteEncodedBy() {
	id3.DeleteTag(tagEncodedBy)
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
			result, err := DecodeString(data[1:], TextEncoding(data))
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

// Not recommended
func (id3 *ID3) SetTag(realTagName string, data []byte) bool {
	return id3.tags.SetTag(realTagName, data)
}
