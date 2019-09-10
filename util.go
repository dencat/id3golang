package id3golang

import (
	"bytes"
	"errors"
	"unicode/utf16"
	"unicode/utf8"
)

const (
	encodingUTF8    string = "UTF-8"
	encodingUTF16   string = "UTF-16"
	encodingUTF16BE string = "UTF-16BE"
)

/*
*	Text Encoding for text frame header
*	First byte determinate text encoding. If ISO-8859-1 is used this byte should be $00, if Unicode is used it should be $01
*	Return text encoding. E.g. "utf8", "utf16", etc.
 */
func textEncoding(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	if b[0] == 0 || b[0] == 3 {
		return encodingUTF8
	}

	if b[0] == 1 {
		return encodingUTF16
	}

	if b[0] == 2 {
		return encodingUTF16BE
	}

	return ""
}

func DecodeString(b []byte, encoding string) (string, error) {
	switch encoding {
	case encodingUTF8:
		return string(b), nil
	case encodingUTF16:
		value, err := DecodeUTF16(b)
		if err != nil {
			return "", err
		}
		return value, nil
	case encodingUTF16BE:
		return DecodeUTF16BE(b)
	}

	return "", errors.New("unknown encoding format")
}

// Decode UTF-16 Little Endian to UTF-8
func DecodeUTF16(b []byte) (string, error) {
	if len(b)%2 != 0 {
		return "", errors.New("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}

// Decode UTF-16 Big Endian To UTF-8
func DecodeUTF16BE(b []byte) (string, error) {
	if len(b)%2 != 0 {
		return "", errors.New("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i+1]) + (uint16(b[i]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}

// Convert byte to int
// In some parts of the tag it is inconvenient to use the
// unsychronisation scheme because the size of unsynchronised data is
// not known in advance, which is particularly problematic with size
// descriptors. The solution in ID3v2 is to use synchsafe integers, in
// which there can never be any false synchs. Synchsafe integers are
// integers that keep its highest bit (bit 7) zeroed, making seven bits
//out of eight available. Thus a 32 bit synchsafe integer can store 28
// bits of information.
func ByteToIntSynchsafe(data []byte) int {
	return int(data[3]) | int(data[2])<<7 | int(data[1])<<14 | int(data[0])<<21
}

func IntToByteSynchsafe(data int) []byte {
	// 7F = 0111 1111
	return []byte{
		byte(data>>23) & 0x7F,
		byte(data>>15) & 0x7F,
		byte(data>>7) & 0x7F,
		byte(data) & 0x7F,
	}
}

// Convert byte to int
func ByteToInt(data []byte) int {
	return int(data[3]) | int(data[2])<<8 | int(data[1])<<16 | int(data[0])<<24
}
