package id3golang

/*
*	ID3 Version
*	Now supported id3v1, id3v2.2, id3v2.3, id3v2.4
 */
type Id3Version byte

const (
	TypeID3Undefined Id3Version = 0
	TypeID3v1        Id3Version = 1
	TypeID3v22       Id3Version = 2
	TypeID3v23       Id3Version = 3
	TypeID3v24       Id3Version = 4
)

var id3VersionMap = map[Id3Version]string{
	TypeID3Undefined: "",
	TypeID3v1:        "id3v1",
	TypeID3v22:       "id3v2.2",
	TypeID3v23:       "id3v2.3",
	TypeID3v24:       "id3v2.4",
}

func (v *Id3Version) String() string {
	return id3VersionMap[*v]
}

/*
*	Tags interface
 */
type TagReader interface {
	GetTag(key string) ([]byte, bool)
	SetTag(key string, data []byte) bool
	GetAll() []ID3Tag
}

type ID3Tag struct {
	Key   string
	Value []byte
}
