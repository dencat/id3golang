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
| Year              | Year    |         | TYER    | TYER    |
| Comment           | Comment |         | COMM    | COMM    |
| Genre             | Genre   |         | TCON    | TCON    |
| Album Artist      | -       |         | TPE2    | TPE2    |        

# How to use

```go
id3, err := id3golang.ReadFile("song.mp3")
if err != nil {
	return err
}
fmt.Println(id3.GetTitle())
```