package id3golang

import (
	"bytes"
	"image"
	"strconv"
	"strings"
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
	var comment Comment

	commentStr, ok := id3.getString(tagComment)
	if !ok {
		return nil, false
	}

	switch id3.version {
	case TypeID3v1:
		comment.Language = ""
		comment.Text = commentStr
	default:
		// id3v2
		// Comment struct must be greater than 4
		// [lang \x00 text] - comment format
		// lang - 3 symbols
		// \x00 - const, delimeter
		// text - all after
		if len(commentStr) < 4 {
			return nil, false
		}

		comment.Language = commentStr[0:3]
		comment.Text = commentStr[4:]
	}

	return &comment, true
}

func (id3 *ID3) GetGenre() (string, bool) {
	switch id3.version {
	case TypeID3v1:
		data, ok := id3.getData(tagGenre)
		if !ok {
			return "", false
		}

		genreCode := int(data[0])
		genre, ok := genres[genreCode]

		if genreCode == 255 || !ok {
			return "", false
		}

		return genre, true
	default:
		return id3.getString(tagGenre)
	}

}

func (id3 *ID3) GetAlbumArtist() (string, bool) {
	return id3.getString(tagAlbumArtist)
}

func (id3 *ID3) SetAlbumArtist(albumArtist string) bool {
	return id3.setString(tagAlbumArtist, albumArtist)
}

func (id3 *ID3) DeleteAlbumArtist() {
	id3.DeleteTag(tagAlbumArtist)
}

func (id3 *ID3) GetDate() (time.Time, bool) {
	return id3.getTimestamp(tagDate)
}

func (id3 *ID3) SetDate(date string) bool {
	return id3.setString(tagDate, date)
}

func (id3 *ID3) DeleteDate() {
	id3.DeleteTag(tagDate)
}

func (id3 *ID3) GetTrackNumber() (int, bool) {
	return id3.getInt(tagTrackNumber)
}

func (id3 *ID3) SetTrackNumber(trackNumber int) bool {
	return id3.setString(tagTrackNumber, strconv.Itoa(trackNumber))
}

func (id3 *ID3) DeleteTrackNumber() {
	id3.DeleteTag(tagTrackNumber)
}

func (id3 *ID3) GetArranger() (string, bool) {
	return id3.getString(tagArranger)
}

func (id3 *ID3) SetArranger(arranger string) bool {
	return id3.setString(tagArranger, arranger)
}

func (id3 *ID3) DeleteArranger() {
	id3.DeleteTag(tagArranger)
}

func (id3 *ID3) GetAuthor() (string, bool) {
	return id3.getString(tagAuthor)
}

func (id3 *ID3) SetAuthor(author string) bool {
	return id3.setString(tagAuthor, author)
}

func (id3 *ID3) DeleteAuthor() {
	id3.DeleteTag(tagAuthor)
}

func (id3 *ID3) GetBPM() (int, bool) {
	return id3.getInt(tagBPM)
}

func (id3 *ID3) SetBPM(bpm int) bool {
	return id3.setString(tagBPM, strconv.Itoa(bpm))
}

func (id3 *ID3) DeleteBPM() {
	id3.DeleteTag(tagBPM)
}

func (id3 *ID3) GetCatalogNumber() (string, bool) {
	return id3.getString(tagCatalogNumber)
}

func (id3 *ID3) SetCatalogNumber(catalogNumber string) bool {
	return id3.setString(tagCatalogNumber, catalogNumber)
}

func (id3 *ID3) DeleteCatalogNumber() {
	id3.DeleteTag(tagCatalogNumber)
}

func (id3 *ID3) GetCompilation() (string, bool) {
	return id3.getString(tagCompilation)
}

func (id3 *ID3) SetCompilation(compilation string) bool {
	return id3.setString(tagCompilation, compilation)
}

func (id3 *ID3) DeleteCompilation() {
	id3.DeleteTag(tagCompilation)
}

func (id3 *ID3) GetComposer() (string, bool) {
	return id3.getString(tagComposer)
}

func (id3 *ID3) SetComposer(composer string) bool {
	return id3.setString(tagComposer, composer)
}

func (id3 *ID3) DeleteComposer() {
	id3.DeleteTag(tagComposer)
}

func (id3 *ID3) GetConductor() (string, bool) {
	return id3.getString(tagConductor)
}

func (id3 *ID3) SetConductor(conductor string) bool {
	return id3.setString(tagConductor, conductor)
}

func (id3 *ID3) DeleteConductor() {
	id3.DeleteTag(tagConductor)
}

func (id3 *ID3) GetCopyright() (string, bool) {
	return id3.getString(tagCopyright)
}

func (id3 *ID3) SetCopyright(copyright string) bool {
	return id3.setString(tagCopyright, copyright)
}

func (id3 *ID3) DeleteCopyright() {
	id3.DeleteTag(tagCopyright)
}

func (id3 *ID3) GetDescription() (string, bool) {
	return id3.getString(tagDescription)
}

func (id3 *ID3) SetDescription(description string) bool {
	return id3.setString(tagDescription, description)
}

func (id3 *ID3) DeleteDescription() {
	id3.DeleteTag(tagDescription)
}

func (id3 *ID3) GetDiscNumber() (string, bool) {
	return id3.getString(tagDiscNumber)
}

func (id3 *ID3) SetDiscNumber(discNumber string) bool {
	return id3.setString(tagDiscNumber, discNumber)
}

func (id3 *ID3) DeleteDiscNumber() {
	id3.DeleteTag(tagDiscNumber)
}

func (id3 *ID3) GetEncodedBy() (string, bool) {
	return id3.getString(tagEncodedBy)
}

func (id3 *ID3) SetEncodedBy(encodedBy string) bool {
	return id3.setString(tagEncodedBy, encodedBy)
}

func (id3 *ID3) DeleteEncodedBy() {
	id3.DeleteTag(tagEncodedBy)
}

type Picture struct {
	Mime        string
	Description string
	PictureType PictureType
	Image       image.Image
}

type PictureType int

const (
	PictureTypeOther              PictureType = 0
	PictureType32x32FileIcon      PictureType = 1
	PictureTypeOtherFileIcon      PictureType = 2
	PictureTypeCoverFront         PictureType = 3
	PictureTypeCoverBack          PictureType = 4
	PictureTypeLeafletPage        PictureType = 5
	PictureTypeMedia              PictureType = 6
	PictureTypeLeadArtist         PictureType = 7
	PictureTypeArtist             PictureType = 8
	PictureTypeConductor          PictureType = 9
	PictureTypeBand               PictureType = 10
	PictureTypeComposer           PictureType = 11
	PictureTypeLyricist           PictureType = 12
	PictureTypeRecordingLocation  PictureType = 13
	PictureTypeDuringRecording    PictureType = 14
	PictureTypeDuringPerformance  PictureType = 15
	PictureTypeScreenCapture      PictureType = 16
	PictureTypeBrightColouredFish PictureType = 17
	PictureTypeIllustration       PictureType = 18
	PictureTypeBandLogotype       PictureType = 19
	PictureTypePublisherLogotype  PictureType = 20
)

func (p *PictureType) String() string {
	switch *p {
	case PictureTypeOther:
		return "Other"
	case PictureType32x32FileIcon:
		return "32x32 pixels 'file icon' (PNG only)"
	case PictureTypeOtherFileIcon:
		return "Other file icon"
	case PictureTypeCoverFront:
		return "Cover (front)"
	case PictureTypeCoverBack:
		return "Cover (back)"
	case PictureTypeLeafletPage:
		return "Leaflet page"
	case PictureTypeMedia:
		return "Media (e.g. lable side of CD)"
	case PictureTypeLeadArtist:
		return "Lead artist/lead performer/soloist"
	case PictureTypeArtist:
		return "Artist/performer"
	case PictureTypeConductor:
		return "Conductor"
	case PictureTypeBand:
		return "Band/Orchestra"
	case PictureTypeComposer:
		return "Composer"
	case PictureTypeLyricist:
		return "Lyricist/text writer"
	case PictureTypeRecordingLocation:
		return "Recording Location"
	case PictureTypeDuringRecording:
		return "During recording"
	case PictureTypeDuringPerformance:
		return "During performance"
	case PictureTypeScreenCapture:
		return "Movie/video screen capture"
	case PictureTypeBrightColouredFish:
		return "A bright coloured fish"
	case PictureTypeIllustration:
		return "Illustration"
	case PictureTypeBandLogotype:
		return "Band/artist logotype"
	case PictureTypePublisherLogotype:
		return "Publisher/Studio logotype"
	}
	return ""
}

func (id3 *ID3) GetPicture() (*Picture, bool) {
	var picture Picture
	str, ok := id3.getString(tagPicture)
	if !ok {
		return nil, false
	}

	data := strings.SplitN(str, "\x00", 3)
	if len(data) != 3 {
		return nil, false
	}

	// Mime
	picture.Mime = data[0]
	if len(data[1]) == 0 {
		return nil, false
	}

	picture.PictureType = PictureType(data[1][0])
	picture.Description = data[1][1:]

	switch picture.Mime {
	case "image/jpeg", "image/png":
		img, _, err := image.Decode(bytes.NewReader([]byte(data[2])))
		if err != nil {
			return nil, false
		}
		picture.Image = img
	default:
		return nil, false
	}

	return &picture, true
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
		case TypeID3v22, TypeID3v23, TypeID3v24:
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
