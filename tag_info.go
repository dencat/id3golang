package id3golang

type tagType int

const (
	TAG_TYPE_TEXT tagType = 1
	TAG_TYPE_INT  tagType = 2
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
	ID3v1       *tagInfo
	ID3v22Tag   *tagInfo
	ID3v23Tag   *tagInfo
	ID3v24Tag   *tagInfo
}

var tagsInfo = map[tagName]TagsInfo{
	tagTitle:       {ID3v1: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Title"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TIT2"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TIT2"}, Description: "Title/songname/content description"},
	tagArtist:      {ID3v1: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Artist"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TPE1"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TPE1"}, Description: "Lead performer(s)/Soloist(s)"},
	tagAlbum:       {ID3v1: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Album"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TALB"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TALB"}, Description: "Album/Movie/Show title"},
	tagYear:        {ID3v1: &tagInfo{Type: TAG_TYPE_INT, Name: "Year"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TYER"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TYER"}, Description: "Year"},
	tagComment:     {ID3v1: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Comment"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "COMM"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "COMM"}, Description: "Comments"},
	tagGenre:       {ID3v1: &tagInfo{Type: TAG_TYPE_TEXT, Name: "Genre"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TCON"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TCON"}, Description: "Content type"},
	tagAlbumArtist: {ID3v1: &tagInfo{Type: TAG_TYPE_TEXT, Name: "AlbumArtist"}, ID3v22Tag: nil, ID3v23Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TPE2"}, ID3v24Tag: &tagInfo{Type: TAG_TYPE_TEXT, Name: "TPE2"}, Description: "Band/orchestra/accompaniment"},
}
