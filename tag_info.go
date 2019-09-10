package id3golang

type tagType int

const (
	TAG_TYPE_UNDEFINED tagType = 0
	TAG_TYPE_TEXT      tagType = 1
	TAG_TYPE_INT       tagType = 2
)

type tagName string

const (
	tagTitle       tagName = "Title"
	tagArtist      tagName = "Artist"
	tagAlbum       tagName = "Album"
	tagYear        tagName = "Year"
	tagComment     tagName = "Comment"
	tagGenre       tagName = "Genre"
	tagAlbumArtist tagName = "AlbumArtist"
)

type tagInfo struct {
	Type tagType
	Name string
}

type TagsInfo struct {
	Description string
	ID3v1Tag    *tagInfo
	ID3v22Tag   *tagInfo
	ID3v23Tag   *tagInfo
	ID3v24Tag   *tagInfo
}

var tagsInfo = map[tagName]TagsInfo{
	tagTitle:       {ID3v1Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Title"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TIT2"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TIT2"}, Description: "Title/songname/content description"},
	tagArtist:      {ID3v1Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Artist"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TPE1"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TPE1"}, Description: "Lead performer(s)/Soloist(s)"},
	tagAlbum:       {ID3v1Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Album"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TALB"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TALB"}, Description: "Album/Movie/Show title"},
	tagYear:        {ID3v1Tag: &tagInfo{Type: TAG_TYPE_INT, Name: "Year"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TYER"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TDOR"}, Description: "Year"},
	tagComment:     {ID3v1Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Comment"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "COMM"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "COMM"}, Description: "Comments"},
	tagGenre:       {ID3v1Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Genre"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TCON"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TCON"}, Description: "Content type"},
	tagAlbumArtist: {ID3v1Tag: nil, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TPE2"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TPE2"}, Description: "Band/orchestra/accompaniment"},
}

func getTagInfo(tagName tagName, version Id3Version) *tagInfo {
	switch version {
	case TypeID3v1:
		return tagsInfo[tagName].ID3v1Tag
	case TypeID3v22:
		return tagsInfo[tagName].ID3v22Tag
	case TypeID3v23:
		return tagsInfo[tagName].ID3v23Tag
	case TypeID3v24:
		return tagsInfo[tagName].ID3v24Tag
	}
	return nil
}

func getTagName(name tagName, version Id3Version) string {
	result := getTagInfo(name, version)
	if result == nil {
		return ""
	}
	return result.Name
}

func getTagType(name tagName, version Id3Version) tagType {
	result := getTagInfo(name, version)
	if result == nil {
		return TAG_TYPE_UNDEFINED
	}
	return result.Type
}

func GetAllSupportedTags(version Id3Version) []string {
	result := []string{}
	for key, _ := range tagsInfo {
		info := getTagInfo(key, version)
		if info != nil {
			result = append(result, string(key))
		}
	}
	return result
}

func realTagNameToName(realName string, version Id3Version) tagName {
	for key, _ := range tagsInfo {
		info := getTagInfo(key, version)
		if realName == info.Name {
			return key
		}
	}
	return tagName("")
}
