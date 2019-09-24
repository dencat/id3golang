# id3golang

# Install

```go 
go get github.com/dencat/id3golang
```

# Supported tags

| Name              | ID3v1       | ID3v2.2 | ID3v2.3 | ID3v2.4 |
|-------------------|-------------|---------|---------|---------|
| Title             | Title       | TT2     | TIT2    | TIT2    |
| Artist            | Artist      | TP1     | TPE1    | TPE1    |
| Album             | Album       | TOT     | TALB    | TALB    |
| Year              | Year        | TYE     | TYER    | TDOR    |
| Comment           | Comment     | COM     | COMM    | COMM    |
| Genre             | Genre       | -       | TCON    | TCON    |
| Album Artist      | -           | TOA     | TPE2    | TPE2    | 
| Date              | -           | TIM     | TYER    | TDRC    |
| Arranger          | -           | -       | IPLS    | TIPL    |
| Author            | -           | TOL     | TOLY    | TOLY    |
| BPM               | -           | BPM     | TBPM    | TBPM    |
| Catalog Number    | -           | -       | TXXX    | TXXX    |
| Compilation       | -           | -       | TCMP    | TCMP    |
| Composer          | -           | TCM     | TCOM    | TCOM    |
| Conductor         | -           | TP3     | TPE3    | TPE3    |
| Copyright         | -           | TCR     | TCOP    | TCOP    |
| Description       | -           | TXX     | TIT2    | TIT2    |
| Disc Number       | -           | -       | TPOS    | TPOS    |
| Encoded by        | -           | TEN     | TENC    | TENC    |
| Track Number      | TrackNumber | TRK     | TRCK    | TRCK    |  
| Picture           | -           | PIC     | APIC    | APIC    |
       

# Status 
In progress  
Future features:
*  Convert formats
*  Support all tags (id3 v1, v1.1, v2.2, v2.3, v2.4)
*  Fix errors in files (empty tags, incorrect size, tag size, tag parameters)
*  Command line arguments 

# How to use

```go
id3, err := id3golang.ReadFile("song.mp3")
if err != nil {
	return err
}
fmt.Println(id3.GetTitle())
```

# Examples 

You can read, modify and save tags in your id3 file
```go
// Read file
id3, err := id3golang.ReadFile("my_best_song.mp3")
if err != nil {
	return err
}
// Print all tags
fmt.Println(id3)

// Set tags 
res := id3.SetTitle("Best music ever")
if !res {
	fmt.Println("error set title")
} 

res = id3.SetAlbum("album")
if !res {
	fmt.Println("error set album")
} 

res = id3.SetArtist("My group")
if !res {
	fmt.Println("error set artist")
} 

// Save file
err = id3golang.SaveFile(id3, "my_best_song_with_tags.mp3")
if err != nil {
	fmt.Println(err)
}
```