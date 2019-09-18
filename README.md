# id3golang

# Install

```go 
go get github.com/dencat/id3golang
```

# Supported tags

| Name              | ID3v1   | ID3v2.2 | ID3v2.3 | ID3v2.4 |
|-------------------|---------|---------|---------|---------|
| Title             | Title   |         | TIT2    | TIT2    |
| Artist            | Artist  |         | TPE1    | TPE1    |
| Album             | Album   |         | TALB    | TALB    |
| Year              | Year    |         | TYER    | TDOR    |
| Comment           | Comment |         | COMM    | COMM    |
| Genre             | Genre   |         | TCON    | TCON    |
| Album Artist      | -       |         | TPE2    | TPE2    | 
| Date              | -       |         | TYER    | TDRC    |
| Arranger          | -       |         | IPLS    | TIPL    |
| Author            | -       |         | TOLY    | TOLY    |
| BPM               | -       |         | TBPM    | TBPM    |
| Catalog Number    | -       |         | TXXX    | TXXX    |
| Compilation       | -       |         | TCMP    | TCMP    |
| Composer          | -       |         | TCOM    | TCOM    |
| Conductor         | -       |         | TPE3    | TPE3    |
| Copyright         | -       |         | TCOP    | TCOP    |
| Description       | -       |         | TIT2    | TIT2    |
| Disc Number       | -       |         | TPOS    | TPOS    |
| Encoded by        | -       |         | TENC    | TENC    |
| Track Number      | -       |         | TRCK    | TRCK    |  
       

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